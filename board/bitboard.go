package board

import (
	"fmt"
	"math"
	"strconv"
)

/*
	Draws a visual representation of a bitboard.
*/
func drawBitboard(b uint64) error {
	for row := 0; row < 8; row++ {

		fmt.Print(8 - row, " ")

		for col := 0; col < 8; col ++ {
			square := (float64)((7 - row) * 8 + col)
			digit := (uint64)(math.Pow(2, square))

			if (digit & (uint64)(b) > 0) {
				fmt.Print("• ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println("  a b c d e f g h")

	return nil
}

func bbToString(b uint64) string {
	squareNumber := math.Log2(float64(b))
	row := int(squareNumber / 8)
	colNumber := int(squareNumber) % 8

	return columns[colNumber] + strconv.Itoa(row + 1)
}