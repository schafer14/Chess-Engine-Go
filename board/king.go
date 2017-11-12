package board

import "math"

func kingMoves(board board, color string) []move {
	var friendly uint64
	var bb uint64
	var moves []move = make([]move, 0)

	if color == "w" {
		bb = board.whiteKings
		friendly = whitePieces(board)
	} else {
		bb = board.blackKings
		friendly = blackPieces(board)
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := uint(math.Log2(float64(square)))

		moveBb := kingAttacks[squareNum]
		legalMovesBb := moveBb & (^friendly)

		newMoves := generateMoves(legalMovesBb, func(_ uint64) uint64 {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}