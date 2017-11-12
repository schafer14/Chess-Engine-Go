package board

import (
)

func pawnMoves(board board, color string) []move {
	moves := make([]move, 0)
	empty := ^occupied(board)
	var direction bool
	var bb uint64
	var opponents uint64
	var finalRank uint64
	var startRank uint64
	var cantCapLeft uint64
	var cantCapRight uint64

	if color == "w" {
		direction = true
		bb = board.whitePawns
		opponents = blackPieces(board)
		finalRank = ranks[7]
		startRank = ranks[1]
		cantCapLeft = cols[0]
		cantCapRight = cols[7]
	} else {
		direction = false
		bb = board.blackPawns
		opponents = whitePieces(board)
		finalRank = ranks[0]
		startRank = ranks[6]
		cantCapLeft = cols[7]
		cantCapRight = cols[0]
	}

	opponents |= board.enPassentTarget

	/*
		Will return a function that shifts an uint64 by n
		in the direction specified
	*/
	backN := func(n uint) func(uint64) uint64 {
		return func(sq uint64) uint64 {
			if direction {
				return sq >> n
			} else {
				return sq << n
			}
		}
	}

	forwardN := func(bb uint64, n uint) uint64 {
		if direction {
			return bb << n
		} else {
			return bb >> n
		}
	}

	// Moving forward one square
	forward1 := forwardN(bb, 8) & empty
	moves = append(moves, generateMoves(forward1, backN(8 ))...)

	// Moving forward two squares
	forward2 := forwardN(bb, 16) & empty & forwardN(empty, 8) & forwardN(startRank, 16)
	moves = append(moves, generateMoves(forward2, backN(16 ))...)

	// Capturing Left (for white)
	capLeft := forwardN(bb, 9) & ^cantCapLeft & opponents
	moves = append(moves, generateMoves(capLeft, backN(9 ))...)

	// Capturing Right (for white)
	capRight := forwardN(bb, 7) & ^cantCapRight & opponents
	moves = append(moves, generateMoves(capRight, backN(7 ))...)

	// Pawn promotion stuff.
	if (forward1 | forward2 | capLeft | capRight) & finalRank > 0 {
		oldMoves := moves
		moves = make([]move, 0)


		for _, m := range oldMoves {
			if m.to & finalRank > 0 {
				moves = append(moves, move{m.from, m.to, "Q"})
				moves = append(moves, move{m.from, m.to, "K"})
				moves = append(moves, move{m.from, m.to, "B"})
				moves = append(moves, move{m.from, m.to, "R"})
			} else {
				moves = append(moves, m)
			}
		}
	}

	return moves
}
