package maurice

func (p Position) knightMoves() []Move {
	var friendly = p.attackers()
	var bb = p.PieceBitboards[Knight+p.color]
	var moves []Move

	for bb != 0 {
		square := bb & -bb
		bb &= bb - 1

		squareNum := square.FirstSquare()

		moveBb := knightAttacks[squareNum]
		legalMovesBb := moveBb & (^friendly)

		newMoves := p.movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func (p Position) knightAttacks(color int) Bitboard {
	var bb Bitboard = p.PieceBitboards[Knight+color]
	var attacks Bitboard = 0

	for bb > 0 {
		square := bb & -bb
		bb &= bb - 1

		squareNum := square.FirstSquare()

		attacks |= knightAttacks[squareNum]
	}

	return attacks
}
