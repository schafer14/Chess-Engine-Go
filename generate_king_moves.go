package maurice


func (p Position) kingMoves() []Move {
	var friendly Bitboard = p.attackers()
	var bb Bitboard = p.kings(p.color)
	var moves = make([]Move, 0)

	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		squareNum := square.firstSquare()

		moveBb := kingAttacks[squareNum]
		legalMovesBb := moveBb & (^friendly)

		newMoves := movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func (p Position)kingAttacks(color int) Bitboard {
	var bb Bitboard = p.kings(color)
	var attackBB Bitboard = 0

	for bb > 0 {
		square := bb & -bb
		squareNum := square.firstSquare()
		bb&= bb-1


		attackBB |= kingAttacks[squareNum]
	}

	return attackBB
}


func (p Position) castle() []Move {
	var moves = make([]Move, 0)
	var occ Bitboard = p.occupied()

// Abstract this for any castling position
	castleMoves := func(schedule int) []Move {
		var moves = make([]Move, 0)
		possibleCastle := true
		oppColor := (p.color + 1) % 2

		if p.attacks(oppColor) & castlingAttackSquares[schedule] > 0 {
			possibleCastle = false
		}

		if possibleCastle {
			moves = append(moves, Move{castlingFromSquare[schedule], castlingToSquare[schedule], ""})
		}

		return moves
	}

	if p.color == 0 && p.castlingRights[0] && castlingFreeSquares[0] & occ == 0 {
		moves = append(moves, castleMoves(0)...)
	}

	if p.color == 0 && p.castlingRights[1] && castlingFreeSquares[1] & occ == 0 {
		moves = append(moves, castleMoves(1)...)
	}

	if p.color == 1 && p.castlingRights[2] && castlingFreeSquares[2] & occ == 0 {
		moves = append(moves, castleMoves(2)...)
	}

	if p.color == 1 && p.castlingRights[3] && castlingFreeSquares[3] & occ == 0 {
		moves = append(moves, castleMoves(3)...)
	}

	return moves
}