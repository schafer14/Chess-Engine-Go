package maurice

import "strconv"

type Position struct {
	positionHash   uint64       // Position Hash
	pawnHash       uint64       // Pawn Hash
	pieceBitboards [14]Bitboard // Bitboards for each pice
	pieces         [64]Piece
	positionScore  int
	materialScore  int
	score          Score
	color          int   // White is 0 black is 1
	enPassent      uint8 // A bitboard containing available en passent moves
	castlingRights [4]bool
	count50        uint8
}

func (p Position) attackers() Bitboard {
	return p.pieceBitboards[p.color]
}

func (p Position) defenders() Bitboard {
	return p.pieceBitboards[p.oppColor()]
}

func (p Position) oppColor() int {
	return (p.color + 1) % 2
}

func (p Position) occupied() Bitboard {
	return p.pieceBitboards[White] | p.pieceBitboards[Black]
}

func (p Position) empty() Bitboard {
	return ^p.occupied()
}

func (p Position) isInCheck() bool {
	var king Bitboard = p.pieceBitboards[king(p.oppColor())]

	if p.attacks(p.color)&king > 0 {
		return true
	} else {
		return false
	}
}

func numtoString(num int) string {
	row := int(num) / 8
	colNumber := int(num) % 8

	return columnNames[colNumber] + strconv.Itoa(row+1)
}
