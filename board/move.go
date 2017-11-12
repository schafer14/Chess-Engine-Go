package board

import (
	"math"
	"strconv"
	"strings"
)

func (b board) Move(move move) board {
	newBoard := b

	promotionOrder := [4]string {"K", "B", "R", "Q"}
	pieces := [10]uint64 { newBoard.whitePawns, newBoard.blackPawns, newBoard.whiteKnights, newBoard.blackKnights, newBoard.whiteBishops, newBoard.blackBishops, newBoard.whiteRooks, newBoard.blackRooks, newBoard.whiteQueens, newBoard.blackQueens, }

	for i, pieceBB := range pieces {
		/*
			This is the piece being taken
		*/
		if pieceBB & move.to > 0 {
			pieces[i] = pieceBB & (^move.to)
		}

		/*
			This is the piece being moved
		*/
		if pieceBB & move.from > 0 && move.special == "" {
			pieces[i] = (pieceBB & (^move.from)) | move.to
		}

		/*
			This if statement handles promotion
		*/
		if pieceBB & move.from > 0 && move.special != "" {
			pieces[i] = pieceBB & (^move.from)

			for i1, p := range promotionOrder {
				if p == move.special {
					var indexAdder int = 1
					if b.Turn == "w" {
						indexAdder = 0
					}
					var pieceIndex = 2 * i1 + 2 + indexAdder
					pieces[pieceIndex] = pieces[pieceIndex] | move.to
				}
			}
		}
	}

	// Remove En Passent Piece
	if move.to == b.enPassentTarget && move.from & b.whitePawns > 0 {
		pieces[1] = b.blackPawns ^ b.enPassentTarget >> 8
	} else if  move.to == b.enPassentTarget && move.from & b.blackPawns > 0 {
		pieces[0] = b.whitePawns ^ b.enPassentTarget << 8
	}


	// Set En Passent Target
	if move.from & b.whitePawns > 0 && move.to & (move.from << 16) > 0 {
		newBoard.enPassentTarget = move.to >> 8
	} else if move.from & b.blackPawns > 0 && move.to & (move.from >> 16) > 0 {
		newBoard.enPassentTarget = move.to << 8
	} else {
		newBoard.enPassentTarget = 0
	}


	// White king movements
	if newBoard.whiteKings & move.from > 0 {
		// Castle right
		if move.from == 0x10 && move.to == 0x40 {
			pieces[6] = pieces[6] & ^uint64(0x80) | 0x20
		}

		if move.from == 0x10 && move.to == 0x4 {
			pieces[6] = pieces[6] & ^uint64(0x1) | 0x08
		}

		newBoard.whiteKings = newBoard.whiteKings & (^move.from) | move.to
		newBoard.availableCastling[0] = false
		newBoard.availableCastling[1] = false
	}

	// Black king movements
	if newBoard.blackKings & move.from > 0 {
		// Castle right
		if move.from == 0x1000000000000000 && move.to == 0x4000000000000000 {
			pieces[7] = pieces[7] & ^uint64(0x8000000000000000) | 0x2000000000000000
		}

		if move.from == 0x1000000000000000 && move.to == 0x400000000000000 {
			pieces[7] = pieces[7] & ^uint64(0x100000000000000) | 0x0800000000000000
		}

		newBoard.blackKings= newBoard.blackKings & (^move.from) | move.to
		newBoard.availableCastling[2] = false
		newBoard.availableCastling[3] = false
	}

	// Update castling options
	if move.from & 0x1 > 0 || move.to & 0x1 > 0 {
		newBoard.availableCastling[1] = false
	}
	if move.from & 0x80 > 0 || move.to & 0x80 > 0 {
		newBoard.availableCastling[0] = false
	}
	if move.from & 0x8000000000000000 > 0 || move.to & 0x8000000000000000 > 0 {
		newBoard.availableCastling[2] = false
	}
	if move.from & 0x0100000000000000 > 0 || move.to & 0x0100000000000000 > 0 {
		newBoard.availableCastling[3] = false
	}


	newBoard.whitePawns = pieces[0]
	newBoard.blackPawns = pieces[1]
	newBoard.whiteKnights = pieces[2]
	newBoard.blackKnights = pieces[3]
	newBoard.whiteBishops = pieces[4]
	newBoard.blackBishops = pieces[5]
	newBoard.whiteRooks = pieces[6]
	newBoard.blackRooks = pieces[7]
	newBoard.whiteQueens = pieces[8]
	newBoard.blackQueens = pieces[9]

	if b.Turn == "w" {
		newBoard.Turn = "b"
	} else {
		newBoard.Turn = "w"
	}

	return newBoard
}

func (b board) HumanFriendlyMove(move string) board {
	m := moveFromString(move)

	return b.Move(m)
}

/*
	Creates a algebraic notation representation of a move
*/
func (m move) toString () string {
	str := ""

	squareNumber := math.Log2(float64(m.from))
	row := int(squareNumber / 8)
	colNumber := int(squareNumber) % 8

	str += columns[colNumber] + strconv.Itoa(row + 1)

	squareNumber2 := math.Log2(float64(m.to))
	row2 := int(squareNumber2 / 8)
	colNumber2 := int(squareNumber2) % 8

	str += columns[colNumber2] + strconv.Itoa(row2 + 1)

	if m.special != "" && m.special != "e" {
		str += m.special
	}

	return str
}

/*
	Creates a move struct from a algebraic notation of a move
*/
func moveFromString(m string) move {
	fromChar := string(m[0])
	toChar := string(m[2])
	var fromFile int
	var toFile int

	for i, e := range columns {
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
		return move{ from: 1 << fromSquare, to: 1 << toSquare, special: string(m[4]) }
	} else {
		return move{ from: 1 << fromSquare, to: 1 << toSquare, special: strings.TrimSpace("") }
	}
}

