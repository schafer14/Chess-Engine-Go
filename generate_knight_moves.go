package maurice


func (p Position)knightMoves() []Move {
	var friendly = p.attackers()
	var bb = p.knights(p.color)
	var moves []Move


	for bb != 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := square.firstSquare()

		moveBb := knightAttacks[squareNum]
		legalMovesBb := moveBb & (^friendly)

		newMoves := movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard{
			return square
		})

		moves = append(moves, newMoves...)
	}


	return moves
}

func (p Position) knightAttacks(color int) Bitboard {
	var bb Bitboard = p.knights(color)
	var attacks Bitboard = 0

	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := square.firstSquare()

		attacks |= knightAttacks[squareNum]
	}

	return attacks
}