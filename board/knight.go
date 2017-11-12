package board

func knightMoves(board board, color string) []move {
	var friendly uint64
	var bb uint64
	var moves []move = make([]move, 0)

	if color == "w" {
		bb = board.whiteKnights
		friendly = whitePieces(board)
	} else {
		bb = board.blackKnights
		friendly = blackPieces(board)
	}


	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := magic(square)

		moveBb := knightAttacks[squareNum]
		legalMovesBb := moveBb & (^friendly)

		newMoves := bbToMoves(legalMovesBb, func(_ uint64) uint64 {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func knightAttackBB(board board, color string) uint64 {
	var bb uint64
	var attacks uint64 = 0

	if color == "w" {
		bb = board.whiteKnights
	} else {
		bb = board.blackKnights
	}

	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := magic(square)

		attacks |= knightAttacks[squareNum]
	}

	return attacks
}

