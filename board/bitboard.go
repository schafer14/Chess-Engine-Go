package board

import (
	"fmt"
	"math"
	"strconv"
)

var deBruijn = [64]int{
	0, 47,  1, 56, 48, 27,  2, 60,
	57, 49, 41, 37, 28, 16,  3, 61,
	54, 58, 35, 52, 50, 42, 21, 44,
	38, 32, 29, 23, 17, 11,  4, 62,
	46, 55, 26, 59, 40, 36, 15, 53,
	34, 51, 20, 43, 31, 22, 10, 45,
	25, 39, 14, 33, 19, 30,  9, 24,
	13, 18,  8, 12,  7,  6,  5, 63,
}

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
	squareNumber := magic(b)
	row := int(squareNumber / 8)
	colNumber := int(squareNumber) % 8

	return columns[colNumber] + strconv.Itoa(row + 1)
}

func magic(b uint64) int {
	return deBruijn[((b ^ (b - 1)) * 0x03F79D71B4CB0A89) >> 58]
}
