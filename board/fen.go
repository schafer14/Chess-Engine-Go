package board

import (
	"strings"
	"regexp"
	"strconv"
	"math"
)

/*
	Fen Strings consist of six space separated fields:
		1. Piece Placement
		2. Active color
		3. Castling available
		4. En Passent available
		5. Half moves since last pawn was captured
		6. Full Moves
	The full spec can be found here https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation

	This function takes a description of a board in FEN notation and returns an internal representation
	of the board
*/
func FromFEN(fenString string) board {
	board := board{}
	// Get list of each field in Fen String
	fields := strings.Fields(fenString)

	// Set Pieces
	numRe := regexp.MustCompile(`\d`)
	for i, row := range strings.Split(fields[0], "/") {
		square := uint64((7 - i) * 8)
		for _, c:= range row {
			char := string(c)
			switch {
			case numRe.MatchString(char):
				k, _ := strconv.ParseUint(char, 10, 8)
				square += k
			case char == "p":
				board.blackPawns += 1 << square
				square += 1
			case char == "P":
				board.whitePawns += 1 << square
				square += 1
			case char == "n":
				board.blackKnights += 1 << square
				square += 1
			case char == "N":
				board.whiteKnights += 1 << square
				square += 1
			case char == "b":
				board.blackBishops += 1 << square
				square += 1
			case char == "B":
				board.whiteBishops += 1 << square
				square += 1
			case char == "r":
				board.blackRooks += 1 << square
				square += 1
			case char == "R":
				board.whiteRooks += 1 << square
				square += 1
			case char == "q":
				board.blackQueens += 1 << square
				square += 1
			case char == "Q":
				board.whiteQueens += 1 << square
				square += 1
			case char == "k":
				board.blackKings += 1 << square
				square += 1
			case char == "K":
				board.whiteKings += 1 << square
				square += 1
			}
		}
	}



	// Set Active Color
	board.Turn = string(fields[1])

	// Set Castling Options
	board.availableCastling = [4]bool{false, false, false, false}
	if strings.Contains(fields[2], "K") {
		board.availableCastling[0] = true
	}
	if strings.Contains(fields[2], "Q") {
		board.availableCastling[1] = true
	}
	if strings.Contains(fields[2], "k") {
		board.availableCastling[2] = true
	}
	if strings.Contains(fields[2], "q") {
		board.availableCastling[3] = true
	}


	// Set En Passent target
	if fields[3] == "-" {
		board.enPassentTarget = 0
	} else {
		char := string(fields[3][0])
		var file int = 0

		for i, e := range columns {
			if e == char {
				file = i
			}
		}

		rank, _ := strconv.Atoi(string(fields[3][1]))
		enPassentSquare := uint(8 * (rank - 1) + file)

		board.enPassentTarget = 1 << enPassentSquare

	}

	// Set Half Move Count
	halfMoves, _ := strconv.ParseInt(fields[4], 10, 32)

	board.halfMoveCount = int(halfMoves)

	// Set Full Move Count
	fullMoves, _ := strconv.ParseInt(fields[5], 10, 32)

	board.fullMoveCount = int(fullMoves)

	return board
}

/*
	Creates a FEN string from a board
*/
func ToFEN(b board) string {
	str := ""

	// Add Piece Position
	occupied := occupied(b)

	for row := 7; row >= 0; row -- {
		count := 0

		for col := 0; col < 8; col ++ {
			square := float64(row * 8 + col)
			digit := (uint64)(math.Pow(2, square))

			if (count > 0 && occupied & digit > 0) {
				str += strconv.Itoa(count)
				count = 0
			}

			switch {
			case b.whitePawns & digit > 0:
				str += "P"
			case b.blackPawns & digit > 0:
				str += "p"
			case b.whiteKnights & digit > 0:
				str += "N"
			case b.blackKnights & digit > 0:
				str += "n"
			case b.whiteBishops & digit > 0:
				str += "B"
			case b.blackBishops & digit > 0:
				str += "b"
			case b.whiteRooks & digit > 0:
				str += "R"
			case b.blackRooks & digit > 0:
				str += "r"
			case b.whiteQueens & digit > 0:
				str += "Q"
			case b.blackQueens & digit > 0:
				str += "q"
			case b.whiteKings & digit > 0:
				str += "K"
			case b.blackKings & digit > 0:
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
	str += " " + b.Turn

	// Add Castling Avaliability
	str += " "
	if b.availableCastling[0] {
		str += "K"
	}
	if b.availableCastling[1] {
		str += "Q"
	}
	if b.availableCastling[2] {
		str += "k"
	}
	if b.availableCastling[3] {
		str += "q"
	}
	if !b.availableCastling[0] && !b.availableCastling[1] && !b.availableCastling[2] && !b.availableCastling[3] {
		str += "-"
	}

	// Add En Passent Target
	if b.enPassentTarget == 0 {
		str += " -"
	} else {
		str += " " + bbToString(b.enPassentTarget)
	}

	// Add half moves since pawn moved
	str += " " + strconv.Itoa(b.halfMoveCount)

	// Add move count
	str += " " + strconv.Itoa(b.fullMoveCount)

	return str
}