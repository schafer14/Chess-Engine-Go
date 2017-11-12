package board

import (
	"math/bits"
)

type move struct {
	from, to uint64
	special string
}

func (b board) HumanFriendlyMoves() []string {
	allMoves := b.Moves()

	strMoves := make([]string, 0)

	for _, m := range allMoves {
		strMoves = append(strMoves, m.toString())
	}

	return strMoves
}

/*
	Returns a list of all legal moves for the boards position
*/
func (b board) Moves() []move {
	allMoves := CandidateMoves(b)
	legalMoves := make([]move, 0)

	for _, m := range allMoves {
		isLegal := true
		nb := b.Move(m)
		var kings uint64

		if nb.Turn == "w" {
			kings = nb.blackKings
		} else {
			kings = nb.whiteKings
		}

		for _, nm := range attacks(nb, nb.Turn) {
			if nm.to & kings > 0 {
				isLegal = false
				break
			}
		}

		if isLegal {
			legalMoves = append(legalMoves, m)
		}
	}

	return legalMoves
}

/*
	Returns a list of moves without validating that a player is moving into check.
*/
func CandidateMoves(b board) []move {
	allMoves := attacks(b, b.Turn)
	allMoves = append(allMoves, castle(b, b.Turn)...)

	return allMoves
}

/*
	Returns a list of attacks a player can make, this is not a complete move set
	as it does not include special moves or prevent moving into check.
*/
func attacks (b board, color string) []move {
	allMoves := make([]move, 0)

	allMoves = append(allMoves, pawnMoves(b, color)...)
	allMoves = append(allMoves, rookMoves(b, color)...)
	allMoves = append(allMoves, bishopMoves(b, color)...)
	allMoves = append(allMoves, queenMoves(b, color)...)
	allMoves = append(allMoves, knightMoves(b, color)...)
	allMoves = append(allMoves, kingMoves(b, color)...)

	return allMoves
}

/*
	Given a bitboard and a function to make a move given the end position
	will create a list of moves based on each bit in the resulting bitboard
	being a valid destination
*/
func generateMoves(bb uint64, fn func(uint64) uint64) []move {
	moves := make([]move, 0)

	for bb > 0 {
		square := bb & -bb
		bb&= bb-1

		moves = append(moves, move{fn(square), square, ""})
	}

	return moves
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func straightBB(occ uint64, square uint) uint64 {
	forward := slideAttacks(occ, square, cols[square % 8])
	right := slideAttacks(occ, square, ranks[square / 8])
	backwards := reversSlideAttacks(occ, square, cols[square % 8])
	left := reversSlideAttacks(occ, square, ranks[square / 8])

	return forward | right | backwards | left
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func diagBB(occ uint64, square uint) uint64 {
	mask := diag[((square / 8) - (square % 8)) & 15]
	antiMask := antiDiag[7 ^ ((square / 8) + (square % 8))]

	northEast := slideAttacks(occ, square, mask)
	northWest := slideAttacks(occ, square, antiMask)
	southWest := reversSlideAttacks(occ, square, mask)
	southEast := reversSlideAttacks(occ, square, antiMask)

	return northEast | southWest | northWest | southEast
}

/*
	Generates move bitboard for sliding pieces using positive rays
*/
func slideAttacks (occ uint64, square uint, mask uint64) uint64 {
	potentialBlockers := occ & mask

	diff := potentialBlockers - 2 * uint64(1 << square)
	changed := diff ^ occ

	return changed & mask
}

/*
	Generates move bitboard for sliding pieces using negitive rays
*/
func reversSlideAttacks(occ uint64, square uint, mask uint64) uint64 {
	o := bits.Reverse64(occ)
	s := bits.Reverse64(1 << square)
	mask = bits.Reverse64(mask)

	potentialBlockers := o & mask

	diff := potentialBlockers - 2*uint64(s)
	changed := diff ^ o

	return bits.Reverse64(changed & mask)
}