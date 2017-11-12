package board

type BitBoard = uint64

func castle(b board, color string) []move {
	var moves []move = make([]move, 0)
	var occ uint64 = occupied(b)

	// Abstract this for any castling position
	castleMoves := func(oppColor string, schedule int) []move {
		var moves []move = make([]move, 0)
		possibleCastle := true
		attacks := attacks(b, oppColor)

		for _, attack := range attacks {
			if oppColor == "w" && (b.whitePawns << 7 | b.whitePawns << 9) & castlingAttackSquares[schedule] > 1 {
				possibleCastle = false
			}

			if oppColor == "b" && (b.blackPawns >> 7 | b.blackPawns >> 9) & castlingAttackSquares[schedule] > 1 {
				possibleCastle = false
			}

			if attack.to & castlingAttackSquares[schedule] > 0 {
				possibleCastle = false
			}
		}

		if possibleCastle {
			moves = append(moves, move{castlingFromSquare[schedule], castlingToSquare[schedule], ""})
		}

		return moves
	}

	if color == "w" && b.availableCastling[0] && castlingFreeSquares[0] & occ == 0 {
		moves = append(moves, castleMoves("b", 0)...)
	}

	if color == "w" && b.availableCastling[1] && castlingFreeSquares[1] & occ == 0 {
		moves = append(moves, castleMoves("b", 1)...)
	}

	if color == "b" && b.availableCastling[2] && castlingFreeSquares[2] & occ == 0 {
		moves = append(moves, castleMoves("w", 2)...)
	}

	if color == "b" && b.availableCastling[3] && castlingFreeSquares[3] & occ == 0 {
		moves = append(moves, castleMoves("w", 3)...)
	}

	return moves
}
