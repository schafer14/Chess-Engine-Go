package board

import (
	"fmt"
	"math"
)

type board struct {
	// One bitboard for each piece type
	whitePawns, blackPawns 		uint64
	whiteKnights, blackKnights 	uint64
	whiteBishops, blackBishops 	uint64
	whiteRooks, blackRooks 		uint64
	whiteQueens, blackQueens 	uint64
	whiteKings, blackKings 		uint64

	Turn string // White: true; Black: false
	availableCastling [4]bool // [White King side, White Queen side, Black King side, White Queen side]
	enPassentTarget uint64 // Target Square of en passent if available
	halfMoveCount int // Half moves since last pawn captured
	fullMoveCount int // Number of full moves
}

/*
	Draws a board to the console (this will look backwards on a black terminal)
*/
func (b board) Draw() error {
	for row := 0; row < 8; row++ {

		fmt.Print(8 - row, " ")

		for col := 0; col < 8; col ++ {
			square := (float64)((7 - row) * 8 + col)
			digit := (uint64)(math.Pow(2, square))

			if digit & b.blackPawns > 0 {
				fmt.Print("♟ ")
			} else if digit & b.whitePawns > 0 {
				fmt.Print("♙ ")
			} else if digit & b.blackKnights > 0 {
				fmt.Print("♞ ")
			} else if digit & b.whiteKnights > 0 {
				fmt.Print("♘ ")
			} else if digit & b.blackBishops > 0 {
				fmt.Print("♝ ")
			} else if digit & b.whiteBishops > 0 {
				fmt.Print("♗ ")
			} else if digit & b.blackRooks > 0 {
				fmt.Print("♜ ")
			} else if digit & b.whiteRooks > 0 {
				fmt.Print("♖ ")
			} else if digit & b.blackQueens > 0 {
				fmt.Print("♛ ")
			} else if digit & b.whiteQueens > 0 {
				fmt.Print("♕ ")
			} else if digit & b.blackKings > 0 {
				fmt.Print("♚ ")
			} else if digit & b.whiteKings > 0 {
				fmt.Print("♔ ")
			} else {
				fmt.Print("• ")
			}
		}
		fmt.Println()
	}
	if b.Turn == "w" {
		fmt.Print("o")
	} else {
		fmt.Print("•")
	}
	fmt.Println(" a b c d e f g h")

	return nil
}

