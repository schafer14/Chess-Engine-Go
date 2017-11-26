package maurice

func (p Position)rookMoves() []Move {
	var friendly = p.attackers()
	var bb = p.pieceBitboards[Rook + p.color]
	var occ = p.occupied()
	var moves = make([]Move, 0)

	for bb > 0 {
		square := bb & -bb
		squareNum := square.firstSquare()
		bb&= bb-1


		blocker := occ & rookMagic[squareNum].mask
		index := (blocker * rookMagic[squareNum].magic) >> 52
		moveBb := rookMagicMoves[squareNum][index]

		legalMovesBb := moveBb & (^friendly)

		newMoves := p.movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard{
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func (p Position)rookAttacks(color int) Bitboard {
	var bb Bitboard = p.pieceBitboards[Rook + color]
	var occ Bitboard = p.occupied()
	var attackBB Bitboard = 0


	for bb > 0 {
		square := bb & -bb
		squareNum := square.firstSquare()
		bb&= bb-1

		blocker := occ & rookMagic[squareNum].mask
		index := (blocker * rookMagic[squareNum].magic) >> 52
		moveBb := rookMagicMoves[squareNum][index]

		attackBB |= moveBb
	}

	return attackBB
}