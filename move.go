package maurice

import (
	"strconv"
	"strings"
)

type Move struct {
	from		Bitboard
	to	 		Bitboard
	promotion 	string
}

func (m Move) toString() string {
	return m.from.toString() + m.to.toString() + m.promotion
}

func (m Move) nil() bool {
	return m == Move{}
}


// TODO: REWRITE!!
func (p *Position) Move(move Move) *Position {
	newBoard := p

	promotionOrder := [4]string {"K", "B", "R", "Q"}
	newBoard.pieceBitboards = p.pieceBitboards


	for i, pieceBB := range newBoard.pieceBitboards {
		/*
			This is the piece being taken
		*/
		if pieceBB & move.to > 0 {
			newBoard.pieceBitboards[i] = pieceBB & (^move.to)
		}

		/*
			This is the piece being moved
		*/
		if pieceBB & move.from > 0 && move.promotion == "" {
			newBoard.pieceBitboards[i] = (pieceBB & (^move.from)) | move.to
		}

		/*
			This if statement handles promotion
		*/
		if pieceBB & move.from > 0 && move.promotion != "" {
			newBoard.pieceBitboards[i] = pieceBB & (^move.from)

			for i1, pro := range promotionOrder {
				if pro == move.promotion {
					var indexAdder int = 1
					if p.color == 0 {
						indexAdder = 0
					}
					var pieceIndex = 2 * i1 + 2 + indexAdder
					newBoard.pieceBitboards[pieceIndex] |= move.to
				}
			}
		}
	}

	// Remove En Passent Piece
	if move.to == p.enPassent && move.from & p.pawns(0) > 0 {
		newBoard.pieceBitboards[1] = p.pawns(1) ^ p.enPassent >> 8
	} else if  move.to == p.enPassent && move.from & p.pawns(1) > 0 {
		newBoard.pieceBitboards[0] = p.pawns(0) ^ p.enPassent << 8
	}


	// Set En Passent Target
	if move.from & p.pawns(0) > 0 && move.to & (move.from << 16) > 0 {
		newBoard.enPassent = move.to >> 8
	} else if move.from & p.pawns(1) > 0 && move.to & (move.from >> 16) > 0 {
		newBoard.enPassent = move.to << 8
	} else {
		newBoard.enPassent = 0
	}


	// White king movements
	if newBoard.kings(0) & move.from > 0 {
		// Castle right
		if move.from == 0x10 && move.to == 0x40 {
			newBoard.pieceBitboards[9] = newBoard.pieceBitboards[9] & ^Bitboard(0x80) | 0x20
		}

		if move.from == 0x10 && move.to == 0x4 {
			newBoard.pieceBitboards[9] = newBoard.pieceBitboards[9] & ^Bitboard(0x1) | 0x08
		}

		newBoard.pieceBitboards[13] = newBoard.kings(0) & (^move.from) | move.to
		newBoard.castlingRights[0] = false
		newBoard.castlingRights[1] = false
	}

	// Black king movements
	if newBoard.kings(1) & move.from > 0 {
		// Castle right
		if move.from == 0x1000000000000000 && move.to == 0x4000000000000000 {
			newBoard.pieceBitboards[10] = newBoard.pieceBitboards[10] & ^Bitboard(0x8000000000000000) | 0x2000000000000000
		}

		if move.from == 0x1000000000000000 && move.to == 0x400000000000000 {
			newBoard.pieceBitboards[10] = newBoard.pieceBitboards[10] & ^Bitboard(0x100000000000000) | 0x0800000000000000
		}

		newBoard.pieceBitboards[14] = newBoard.kings(1) & (^move.from) | move.to
		newBoard.castlingRights[2] = false
		newBoard.castlingRights[3] = false
	}

	// Update castling options
	if move.from & 0x1 > 0 || move.to & 0x1 > 0 {
		newBoard.castlingRights[1] = false
	}
	if move.from & 0x80 > 0 || move.to & 0x80 > 0 {
		newBoard.castlingRights[0] = false
	}
	if move.from & 0x8000000000000000 > 0 || move.to & 0x8000000000000000 > 0 {
		newBoard.castlingRights[2] = false
	}
	if move.from & 0x0100000000000000 > 0 || move.to & 0x0100000000000000 > 0 {
		newBoard.castlingRights[3] = false
	}

	newBoard.pieceBitboards[0] = newBoard.pieceBitboards[1] | newBoard.pieceBitboards[2]

	newBoard.color = (p.color + 1) % 2


	return newBoard
}


func (p Position) HumanFriendlyMove(move string) Position {
	m := moveFromString(move)

	return p.Move(m)
}

func moveFromString(m string) Move {
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

	fromSquare := uint(8 * (fromRank - 1) + fromFile)
	toSquare := uint(8 * (toRank - 1) + toFile)


	if len(strings.TrimSpace(m)) > 4 {
		return Move{ from: 1 << fromSquare, to: 1 << toSquare, promotion: string(m[4]) }
	} else {
		return Move{ from: 1 << fromSquare, to: 1 << toSquare, promotion: strings.TrimSpace("") }
	}
}