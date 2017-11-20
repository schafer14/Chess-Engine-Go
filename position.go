package maurice

import (
	"strings"
	"strconv"
	"fmt"
)

var (
	// Maps from a unicode character to the
	// index in  the pieceBitboard
	pieceMap = map[int]int {
		0x50: 3,
		0x70: 4,
		0x4E: 5,
		0x6E: 6,
		0x42: 7,
		0x62: 8,
		0x52: 9,
		0x72: 10,
		0x51: 11,
		0x71: 12,
		0x4B: 13,
		0x6B: 14,
	}
)

type Position struct {
	positionHash 	uint64 // Position Hash
	pawnHash	 	uint64 // Pawn Hash
	pieceBitboards  [15]Bitboard // Bitboards for each pice
	positionScore	int
	materialScore	int
	score			Score
	color 			int // White is 0 black is 1
	enPassent 		Bitboard // A bitboard containing available en passent moves
	castlingRights	[4]bool
	count50			uint8
}

func (p Position) Help() {
	fmt.Println(p)
}

func PositionFromFEN(fen string) Position {
	position := Position{}
	// Get list of each field in Fen String
	fields := strings.Fields(fen)

	// Set Pieces
	for i, row := range strings.Split(fields[0], "/") {
		square := bbFromInts(7-i, 0)
		for _, c:= range row {
			switch {
			// Matches unicode for characters between 0-9
			case c >= 0x30 && c <= 0x39:
				k := uint(c - 0x30)
				square = square << k
			default:
				bitboardIndex := pieceMap[int(c)]
				position.pieceBitboards[0].union(square)
				// Matches unicode characters for upper case characters
				if c >= 0x41 && c <= 0x5A {
					position.pieceBitboards[1].union(square)
				} else {
					position.pieceBitboards[2].union(square)
				}
				position.pieceBitboards[bitboardIndex].union(square)
				square.shift(1)
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
		position.enPassent = bbFromString(fields[3])
	}
	//
	//// Set Half Move Count
	halfMoves, _ := strconv.ParseInt(fields[4], 10, 32)

	position.count50 = uint8(halfMoves)


	return position
}

func (p Position) toFen() string {
	str := ""

	// Add Piece Position
	occupied := p.occupied()

	for row := 7; row >= 0; row -- {
		count := 0

		for col := 0; col < 8; col ++ {
			square := bbFromInts(row, col)

			if count > 0 && occupied.isOn(square) {
				str += strconv.Itoa(count)
				count = 0
			}

			switch {
			case p.pawns(0).isOn(square):
				str += "P"
			case p.pawns(1).isOn(square):
				str += "p"
			case p.knights(0).isOn(square):
				str += "N"
			case p.knights(1).isOn(square):
				str += "n"
			case p.bishops(0).isOn(square):
				str += "B"
			case p.bishops(1).isOn(square):
				str += "p"
			case p.rooks(0).isOn(square):
				str += "R"
			case p.rooks(1).isOn(square):
				str += "r"
			case p.queens(0).isOn(square):
				str += "Q"
			case p.queens(1).isOn(square):
				str += "q"
			case p.kings(0).isOn(square):
				str += "K"
			case p.kings(1).isOn(square):
				str += "k"
			default:
				count += 1
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
		str += " " + p.enPassent.toString()
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
			square := bbFromInts(row, col)

			if p.pawns(1).isOn(square) {
				fmt.Print("♟ ")
			} else if p.pawns(0).isOn(square) {
				fmt.Print("♙ ")
			} else if p.knights(1).isOn(square) {
				fmt.Print("♞ ")
			} else if p.knights(0).isOn(square) {
				fmt.Print("♘ ")
			} else if p.bishops(1).isOn(square) {
				fmt.Print("♝ ")
			} else if p.bishops(0).isOn(square) {
				fmt.Print("♗ ")
			} else if p.rooks(1).isOn(square) {
				fmt.Print("♜ ")
			} else if p.rooks(0).isOn(square) {
				fmt.Print("♖ ")
			} else if p.queens(1).isOn(square) {
				fmt.Print("♛ ")
			} else if p.queens(0).isOn(square) {
				fmt.Print("♕ ")
			} else if p.kings(1).isOn(square) {
				fmt.Print("♚ ")
			} else if p.kings(0).isOn(square) {
				fmt.Print("♔ ")
			} else {
				fmt.Print("• ")
			}
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

func (p Position) pawns(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[3]
	} else {
		return p.pieceBitboards[4]
	}
}

func (p Position) knights(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[5]
	} else {
		return p.pieceBitboards[6]
	}
}

func (p Position) bishops(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[7]
	} else {
		return p.pieceBitboards[8]
	}
}

func (p Position) rooks(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[9]
	} else {
		return p.pieceBitboards[10]
	}
}

func (p Position) queens(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[11]
	} else {
		return p.pieceBitboards[12]
	}
}

func (p Position) kings(color int) Bitboard {
	if color == 0 {
		return p.pieceBitboards[13]
	} else {
		return p.pieceBitboards[14]
	}
}

func (p Position) attackers() Bitboard {
	if p.color == 0 {
		return p.pieceBitboards[1]
	} else {
		return p.pieceBitboards[2]
	}
}

func (p Position) defenders() Bitboard {
	if p.color == 1 {
		return p.pieceBitboards[1]
	} else {
		return p.pieceBitboards[2]
	}
}

func (p Position) occupied() Bitboard {
	return p.pieceBitboards[0]
}

func (p Position) empty() Bitboard {
	return ^ p.occupied()
}

func (p Position) isInCheck() bool {
	var king Bitboard

	if p.color == 0 {
		king = p.kings(1)
	} else {
		king = p.kings(0)
	}

	if p.attacks(p.color) & king > 0 {
		return true
	} else {
		return false
	}
}
