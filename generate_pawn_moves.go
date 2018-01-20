package maurice

// TODO: Rewrite ... seriously!
func (p Position) pawnMoves() []Move {
	moves := make([]Move, 0)
	empty := p.empty()
	var direction bool
	var bb = p.PieceBitboards[Pawn+p.color]
	var opponents = p.defenders() | (1 << p.enPassent)
	var finalRank Bitboard
	var startRank Bitboard
	var cantCapLeft Bitboard
	var cantCapRight Bitboard

	if p.color == 0 {
		direction = true
		finalRank = ranks[7]
		startRank = ranks[1]
		cantCapLeft = columns[0]
		cantCapRight = columns[7]
	} else {
		direction = false
		finalRank = ranks[0]
		startRank = ranks[6]
		cantCapLeft = columns[7]
		cantCapRight = columns[0]
	}

	/*
		Will return a function that shifts an uint64 by n
		in the direction specified
	*/
	backN := func(n uint) func(bitboard Bitboard) Bitboard {
		return func(sq Bitboard) Bitboard {
			if direction {
				return sq >> n
			} else {
				return sq << n
			}
		}
	}

	forwardN := func(bb Bitboard, n uint) Bitboard {
		if direction {
			return bb << n
		} else {
			return bb >> n
		}
	}

	// Moving forward one square
	forward1 := forwardN(bb, 8) & empty
	moves = append(moves, p.movesFromBitboard(forward1, backN(8))...)

	// Moving forward two squares
	forward2 := forwardN(bb, 16) & empty & forwardN(empty, 8) & forwardN(startRank, 16)
	moves = append(moves, p.movesFromBitboard(forward2, backN(16))...)

	// Capturing Left (for white)
	capLeft := forwardN(bb, 9) & ^cantCapLeft & opponents
	moves = append(moves, p.movesFromBitboard(capLeft, backN(9))...)

	// Capturing Right (for white)
	capRight := forwardN(bb, 7) & ^cantCapRight & opponents
	moves = append(moves, p.movesFromBitboard(capRight, backN(7))...)

	// Pawn promotion stuff. TODO: REWRITE
	if (forward1|forward2|capLeft|capRight)&finalRank > 0 {
		oldMoves := moves
		moves = make([]Move, 0)

		for _, m := range oldMoves {
			_, to, _, _ := m.split()

			if 1<<to&finalRank > 0 {
				moves = append(moves, m.Promote(p.color)...)
			} else {
				moves = append(moves, m)
			}
		}
	}

	return moves
}

func (p Position) pawnAttacks(color int) Bitboard {
	var direction bool
	var bb Bitboard = p.PieceBitboards[Pawn+color]
	var cantCapLeft Bitboard
	var cantCapRight Bitboard

	if color == 0 {
		direction = true
		cantCapLeft = columns[0]
		cantCapRight = columns[7]
	} else {
		direction = false
		cantCapLeft = columns[7]
		cantCapRight = columns[0]
	}

	forwardN := func(bb Bitboard, n uint) Bitboard {
		if direction {
			return bb << n
		} else {
			return bb >> n
		}
	}

	// Capturing Left (for white)
	capLeft := forwardN(bb, 9) & ^cantCapLeft

	// Capturing Right (for white)
	capRight := forwardN(bb, 7) & ^cantCapRight

	return capLeft | capRight
}
