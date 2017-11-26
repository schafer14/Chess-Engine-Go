package maurice

import (
	"fmt"
	"strings"
	"strconv"
)

func PositionFromFEN(fen string) Position {
	position := Position{}
	// Get list of each field in Fen String
	fields := strings.Fields(fen)

	// Set Pieces
	for i, row := range strings.Split(fields[0], "/") {
		square := bbFromInts(7-i, 0)
		squareNum := (7 - i) * 8

		for _, c:= range row {
			switch {
			// Matches unicode for characters between 0-9
			case c >= 0x30 && c <= 0x39:
				k := uint(c - 0x30)
				square = square << k
				squareNum += int(k)
			default:
				bitboardIndex := pieceMap[int(c)]
				// Matches unicode characters for upper case characters
				if c >= 0x41 && c <= 0x5A {
					position.pieceBitboards[White].union(square)
				} else {
					position.pieceBitboards[Black].union(square)
				}
				position.pieceBitboards[bitboardIndex].union(square)
				position.pieces[squareNum] = Piece(bitboardIndex)
				square.shift(1)
				squareNum += 1
			}
		}
	}

	// Set Active Color
	position.color = 0
	if fields[1][0] == 0x62 {
		position.color = 1
	}
	//
	//// Set Castling Options
	position.castlingRights = [4]bool{false, false, false, false}
	if strings.Contains(fields[2], "K") {
		position.castlingRights[0] = true
	}
	if strings.Contains(fields[2], "Q") {
		position.castlingRights[1] = true
	}
	if strings.Contains(fields[2], "k") {
		position.castlingRights[2] = true
	}
	if strings.Contains(fields[2], "q") {
		position.castlingRights[3] = true
	}
	//
	//
	//// Set En Passent target
	if fields[3] == "-" {
		position.enPassent = 0
	} else {
		position.enPassent = uint8(squareFromString(fields[3]))
	}
	//
	//// Set Half Move Count
	halfMoves, _ := strconv.ParseInt(fields[4], 10, 32)

	position.count50 = uint8(halfMoves)

	return position
}

func (p Position) ToFen() string {
	str := ""

	for row := 7; row >= 0; row -- {
		count := 0

		for col := 0; col < 8; col ++ {
			piece := p.pieces[row * 8 + col]

			if count > 0 && piece > 0 {
				str += strconv.Itoa(count)
				count = 0
			}

			if piece == 0 {
				count += 1
			} else {
				str += piece.char()
			}
		}

		if count > 0 {
			str += strconv.Itoa(count)
		}
		if row > 0 {
			str += "/"
		}
	}

	// Add Turn
	if p.color == 0 {
		str += " w"
	} else {
		str += " b"
	}

	// Add Castling Avaliability
	str += " "
	if p.castlingRights[0] {
		str += "K"
	}
	if p.castlingRights[1] {
		str += "Q"
	}
	if p.castlingRights[2] {
		str += "k"
	}
	if p.castlingRights[3] {
		str += "q"
	}
	if p.castlingRights == [4]bool{false, false, false, false} {
		str += "-"
	}

	// Add En Passent Target
	if p.enPassent == 0 {
		str += " -"
	} else {
		str += " " + numtoString(int(p.enPassent))
	}

	// Add half moves since pawn moved
	str += " " + strconv.Itoa(int(p.count50))

	// Add move count
	str += " 0"

	return str
}

func (p Position) Draw() {
	for row := 7; row >= 0; row-- {

		fmt.Print(row + 1, " ")

		for col := 0; col < 8; col ++ {
			piece := p.pieces[row * 8 + col]
			fmt.Print(piece.symbol() + " ")
		}
		fmt.Println()
	}
	if p.color == 0 {
		fmt.Print("o")
	} else {
		fmt.Print("•")
	}
	fmt.Println(" a b c d e f g h")
}