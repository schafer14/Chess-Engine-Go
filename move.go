package maurice

import (
	"math"
	"strconv"
	"strings"
)

type Move uint32
type Moves []Move

const (
	isCapture   = 0x00F00000
	isPromo     = 0x0F000000
	isCastle    = 0x10000000
	isEnpassant = 0x20000000
)

// Taken from Donna chess engine.
// Bits 00:00:00:FF => Source square (0 .. 63).
// Bits 00:00:FF:00 => Destination square (0 .. 63).
// Bits 00:0F:00:00 => Piece making the move.
// Bits 00:F0:00:00 => Captured piece if any.
// Bits 0F:00:00:00 => Promoted piece if any.
// Bits F0:00:00:00 => Castle and en-passant flags.

func (s Moves) Len() int {
	return len(s)
}
func (s Moves) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Moves) Less(i, j int) bool {
	_, _, _, c1 := s[i].split()
	_, _, _, c2 := s[j].split()
	return c1 < c2
}

func (m Move) ToString() string {
	from, to, _, _ := m.split()
	promo := ""

	if m.isPromotion() {
		promo = m.promotionPiece().symbol()
	}
	return numtoString(int(from)) + numtoString(int(to)) + promo
}

func NewMove(p Position, from int, to int) Move {
	piece := p.pieces[from]
	captured := p.pieces[to]
	enPassent := 0
	castle := 0

	if piece == pawn(p.color) && int(p.enPassent) == to && p.enPassent > 0 {
		captured = pawn(p.oppColor())
		enPassent = isEnpassant
	}

	if piece == king(p.color) && math.Abs(float64(from-to)) == 2 {
		castle = isCastle
	}

	return Move(from | (to << 8) | (int(piece) << 16) | (int(captured) << 20) | enPassent | castle)
}

func NewMovePromotion(p Position, from int, to int, promo string) Move {
	move := NewMove(p, from, to)

	if promo == "Q" || promo == "q" {
		move |= Queen << 24
	}
	if promo == "R" || promo == "r" {
		move |= Rook << 24
	}
	if promo == "B" || promo == "b" {
		move |= Bishop << 24
	}
	if promo == "N" || promo == "n" {
		move |= Knight << 24
	}

	return move
}

func (m Move) Promote(color int) []Move {
	return []Move{
		m | Move(Queen<<24),
		m | Move(Rook<<24),
		m | Move(Bishop<<24),
		m | Move(Knight<<24),
	}
}

func (p Position) MakeMove(move Move) Position {
	from, to, piece, capture := move.split()

	if capture != 0 {
		p.PieceBitboards[capture] &= ^(1 << to)
		p.PieceBitboards[p.oppColor()] &= ^(1 << to)
		p.pieces[to] = 0
	}

	if move.isEnpassent() && p.enPassent > 0 {
		enPassentSquare := p.enPassent - 8*uint8(1-2*p.color)
		p.pieces[enPassentSquare] = 0
		p.PieceBitboards[p.oppColor()] &= ^(1 << enPassentSquare)
		p.PieceBitboards[Pawn+p.oppColor()] &= ^(1 << enPassentSquare)
	}

	p.PieceBitboards[piece] &= ^(1 << from)
	p.PieceBitboards[piece] |= 1 << to
	p.PieceBitboards[p.color] &= ^(1 << from)
	p.PieceBitboards[p.color] |= 1 << to

	p.pieces[to] = p.pieces[from]
	p.pieces[from] = 0

	if move.isPromotion() {
		promoPiece := move.promotionPiece() + Piece(p.color)
		p.pieces[to] = promoPiece
		p.PieceBitboards[promoPiece] |= 1 << to
		p.PieceBitboards[piece] &= ^(1 << to)
	}

	if move.isCastle() && p.color == White {
		p.castlingRights[0] = false
		p.castlingRights[1] = false
		if to == 6 {
			p.pieces[7] = 0
			p.pieces[5] = Rook
			p.PieceBitboards[Rook] &= ^Bitboard(1 << 7)
			p.PieceBitboards[Rook] |= 1 << 5
			p.PieceBitboards[White] &= ^Bitboard(1 << 7)
			p.PieceBitboards[White] |= 1 << 5
		} else {
			p.pieces[0] = 0
			p.pieces[3] = Rook
			p.PieceBitboards[Rook] &= ^Bitboard(1 << 0)
			p.PieceBitboards[Rook] |= 1 << 3
			p.PieceBitboards[White] &= ^Bitboard(1 << 0)
			p.PieceBitboards[White] |= 1 << 3
		}
	}
	if move.isCastle() && p.color == Black {
		p.castlingRights[2] = false
		p.castlingRights[3] = false
		if to == 62 {
			p.pieces[63] = 0
			p.pieces[61] = BlackRook
			p.PieceBitboards[BlackRook] |= 1 << 61
			p.PieceBitboards[BlackRook] &= ^Bitboard(1 << 63)
			p.PieceBitboards[Black] |= 1 << 61
			p.PieceBitboards[Black] &= ^Bitboard(1 << 63)
		} else {
			p.pieces[56] = 0
			p.pieces[59] = BlackRook
			p.PieceBitboards[BlackRook] |= 1 << 59
			p.PieceBitboards[BlackRook] &= ^Bitboard(1 << 56)
			p.PieceBitboards[Black] |= 1 << 59
			p.PieceBitboards[Black] &= ^Bitboard(1 << 56)
		}
	}

	if to == 7 || from == 7 {
		p.castlingRights[0] = false
	}
	if to == 0 || from == 0 {
		p.castlingRights[1] = false
	}
	if to == 63 || from == 63 {
		p.castlingRights[2] = false
	}
	if to == 56 || from == 56 {
		p.castlingRights[3] = false
	}
	if from == 4 {
		p.castlingRights[0] = false
		p.castlingRights[1] = false
	}
	if from == 60 {
		p.castlingRights[2] = false
		p.castlingRights[3] = false
	}

	// Setting enPassent target
	if (piece == Pawn || piece == BlackPawn) && math.Abs(float64(from)-float64(to)) == 16 {
		p.enPassent = uint8(from) + 8*uint8(1-2*p.color)
	} else {
		p.enPassent = 0
	}

	p.color = p.oppColor()

	return p
}

func (m Move) split() (uint, uint, Piece, Piece) {
	from := uint(0xFF & m)
	to := uint(0xFF & (m >> 8))
	movingPiece := 0xF & (m >> 16)
	capturedPiece := 0xF & (m >> 20)

	return from, to, Piece(movingPiece), Piece(capturedPiece)
}

func (m Move) isEnpassent() bool {
	return (m & isEnpassant) > 0
}

func (m Move) isCastle() bool {
	return (m & isCastle) > 0
}

func (m Move) isPromotion() bool {
	return (m & isPromo) > 0
}

func (m Move) promotionPiece() Piece {
	return Piece((m & isPromo) >> 24)
}

func (p Position) HumanFriendlyMove(move string) Position {
	m := p.moveFromString(move)

	return p.MakeMove(m)
}

func (p Position) moveFromString(m string) Move {
	fromChar := string(m[0])
	toChar := string(m[2])
	var fromFile int
	var toFile int

	for i, e := range columnNames {
		if e == fromChar {
			fromFile = i
		}
		if e == toChar {
			toFile = i
		}
	}

	fromRank, _ := strconv.Atoi(string(m[1]))
	toRank, _ := strconv.Atoi(string(m[3]))

	fromSquare := 8*(fromRank-1) + fromFile
	toSquare := 8*(toRank-1) + toFile

	if len(strings.TrimSpace(m)) > 4 {
		return NewMovePromotion(p, fromSquare, toSquare, string(m[4]))
	} else {
		return NewMove(p, fromSquare, toSquare)
	}
}
