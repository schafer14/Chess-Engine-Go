package board

import (
	"math/bits"
)

type move struct {
	from, to uint64
	special string
}

func (b board) HumanFriendlyMoves() []string {
	allMoves := b.LegalMoves()

	strMoves := make([]string, 0)

	for _, m := range allMoves {
		strMoves = append(strMoves, m.toString())
	}

	return strMoves
}

/*
	Returns a list of all legal moves for the boards position
*/
func (b board) LegalMoves() []move {
	allMoves := b.PseudoMoves()
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

		if nb.Attacks(nb.Turn) & kings > 0 {
			isLegal = false
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
func (b board) PseudoMoves() []move {
	allMoves := make([]move, 0)
	allMoves = append(allMoves, pawnMoves(b, b.Turn)...)
	allMoves = append(allMoves, rookMoves(b, b.Turn)...)
	allMoves = append(allMoves, bishopMoves(b, b.Turn)...)
	allMoves = append(allMoves, queenMoves(b, b.Turn)...)
	allMoves = append(allMoves, knightMoves(b, b.Turn)...)
	allMoves = append(allMoves, kingMoves(b, b.Turn)...)

	allMoves = append(allMoves, castle(b, b.Turn)...)

	return allMoves
}

/*
	Returns a list of Attacks a player can make, this is not a complete move set
	as it does not include special moves or prevent moving into check.
*/
func (b board) Attacks(color string) uint64 {
	var attacks uint64 = 0

	attacks |= pawnAttacks(b, color)
	attacks |= rookAttackBB(b, color)
	attacks |= bishopAttackBB(b, color)
	attacks |= queenAttackBB(b, color)
	attacks |= knightAttackBB(b, color)
	attacks |= kingAttackBB(b, color)

	return attacks
}

/*
	Given a bitboard and a function to make a move given the end position
	will create a list of moves based on each bit in the resulting bitboard
	being a valid destination
*/
func bbToMoves(bb uint64, fn func(uint64) uint64) []move {
	moves := make([]move, 0)

	for bb > 0 {
		square := bb & -bb
		bb &= bb-1

		moves = append(moves, move{fn(square), square, ""})
	}

	return moves
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func straightBB(occ uint64, square uint64) uint64 {
	squareNum := magic(square)

	forward := slideAttacks(occ, square, cols[squareNum % 8])
	right := slideAttacks(occ, square, ranks[squareNum / 8])
	backwards := reversSlideAttacks(occ, square, cols[squareNum % 8])
	left := reversSlideAttacks(occ, square, ranks[squareNum / 8])

	return forward | right | backwards | left
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func diagBB(occ uint64, square uint64) uint64 {
	squareNum := magic(square)

	mask := diag[((squareNum / 8) - (squareNum % 8)) & 15]
	antiMask := antiDiag[7 ^ ((squareNum / 8) + (squareNum % 8))]

	northEast := slideAttacks(occ, square, mask)
	northWest := slideAttacks(occ, square, antiMask)
	southWest := reversSlideAttacks(occ, square, mask)
	southEast := reversSlideAttacks(occ, square, antiMask)

	return northEast | southWest | northWest | southEast
}

/*
	Generates move bitboard for sliding pieces using positive rays
*/
func slideAttacks (occ uint64, square uint64, mask uint64) uint64 {
	potentialBlockers := occ & mask

	diff := potentialBlockers - 2 * square
	changed := diff ^ occ

	return changed & mask
}

/*
	Generates move bitboard for sliding pieces using negitive rays
*/
func reversSlideAttacks(occ uint64, square uint64, mask uint64) uint64 {
	o := bits.Reverse64(occ)
	s := bits.Reverse64(square)
	mask = bits.Reverse64(mask)

	potentialBlockers := o & mask

	diff := potentialBlockers - 2*uint64(s)
	changed := diff ^ o

	return bits.Reverse64(changed & mask)
}