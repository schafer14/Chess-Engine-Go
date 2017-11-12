package board


func queenMoves(board board, color string) []move {
	var friendly uint64
	var bb uint64
	var occ uint64 = occupied(board)
	var moves []move = make([]move, 0)

	if (color == "w") {
		bb = board.whiteQueens
		friendly = whitePieces(board)
	} else {
		bb = board.blackQueens
		friendly = blackPieces(board)
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		moveBb := diagBB(occ, square) | straightBB(occ, square)
		legalMovesBb := moveBb & (^friendly)

		newMoves := bbToMoves(legalMovesBb, func(_ uint64) uint64 {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func queenAttackBB(board board, color string) uint64 {
	var bb uint64
	var occ uint64 = occupied(board)
	var attackBB uint64 = 0

	if color == "w" {
		bb = board.whiteQueens
	} else {
		bb = board.blackQueens
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		attackBB |= diagBB(occ, square) | straightBB(occ, square)

	}

	return attackBB
}