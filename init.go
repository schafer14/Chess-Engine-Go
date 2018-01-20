package maurice

import (
	"math/bits"
)

func init() {
	initMagic()
}

func initMagic() {
	for i := 0; i < 64; i++ {
		buildBitBoardSquare(i, "rook")
		buildBitBoardSquare(i, "bishop")
	}
}

func buildBitBoardSquare(square int, piece string) {
	var mask Bitboard
	var magic Bitboard
	var squareBB Bitboard = Bitboard(1 << uint(square))
	var fn func(Bitboard, Bitboard) Bitboard
	var bits uint

	switch piece {
	case "rook":
		mask = rookMagic[square].mask
		magic = rookMagic[square].magic
		fn = straightBB
		bits = 52
	case "bishop":
		mask = bishopMagic[square].mask
		magic = bishopMagic[square].magic
		fn = diagBB
		bits = 55
	}

	permutations := getPermutations(0, mask)

	for _, blocker := range permutations {
		index := (blocker * magic) >> bits
		if piece == "rook" {
			rookMagicMoves[square][index] = fn(blocker, squareBB)
		}
		if piece == "bishop" {
			bishopMagicMoves[square][index] = fn(blocker, squareBB)
		}
	}
}

func getPermutations(set Bitboard, mutable Bitboard) []Bitboard {
	if mutable.Count() == 0 {
		return []Bitboard{set}
	}

	bit := mutable & -mutable
	mutable ^= bit

	withBitSet := getPermutations(set|bit, mutable)
	withoutBitSet := getPermutations(set, mutable)

	return append(withBitSet, withoutBitSet...)
}

func straightBB(occ Bitboard, square Bitboard) Bitboard {
	squareNum := bits.TrailingZeros64(uint64(square))

	forward := slideAttacks(occ, square, columns[squareNum%8])
	right := slideAttacks(occ, square, ranks[squareNum/8])
	backwards := reversSlideAttacks(occ, square, columns[squareNum%8])
	left := reversSlideAttacks(occ, square, ranks[squareNum/8])

	return forward | right | backwards | left
}

/*
	Generates a bitboard containing all the legal straight moves.
*/
func diagBB(occ Bitboard, square Bitboard) Bitboard {
	squareNum := bits.TrailingZeros64(uint64(square))

	mask := diag[((squareNum/8)-(squareNum%8))&15]
	antiMask := antiDiag[7^((squareNum/8)+(squareNum%8))]

	northEast := slideAttacks(occ, square, mask)
	northWest := slideAttacks(occ, square, antiMask)
	southWest := reversSlideAttacks(occ, square, mask)
	southEast := reversSlideAttacks(occ, square, antiMask)

	return northEast | southWest | northWest | southEast
}

/*
	Generates move bitboard for sliding pieces using positive rays
*/
func slideAttacks(occ Bitboard, square Bitboard, mask Bitboard) Bitboard {
	potentialBlockers := occ & mask

	diff := potentialBlockers - 2*square
	changed := diff ^ occ

	return changed & mask
}

/*
	Generates move bitboard for sliding pieces using negitive rays
*/
func reversSlideAttacks(occ Bitboard, square Bitboard, maskB Bitboard) Bitboard {
	o := bits.Reverse64(uint64(occ))
	s := bits.Reverse64(uint64(square))
	mask := bits.Reverse64(uint64(maskB))

	potentialBlockers := o & mask

	diff := potentialBlockers - 2*uint64(s)
	changed := diff ^ o

	return Bitboard(bits.Reverse64(changed & mask))
}
