package maurice

import (
	"fmt"
	"math"
	"strconv"
)

var (
	deBruijn = [64]int{
		0, 47,  1, 56, 48, 27,  2, 60,
		57, 49, 41, 37, 28, 16,  3, 61,
		54, 58, 35, 52, 50, 42, 21, 44,
		38, 32, 29, 23, 17, 11,  4, 62,
		46, 55, 26, 59, 40, 36, 15, 53,
		34, 51, 20, 43, 31, 22, 10, 45,
		25, 39, 14, 33, 19, 30,  9, 24,
		13, 18,  8, 12,  7,  6,  5, 63,
	}
)

type Bitboard uint64

func (b Bitboard) Draw() {
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
}

func (b Bitboard) isOn(bit Bitboard) bool {
	return b & bit > 0
}

func (b Bitboard) firstSquare() int {
	return deBruijn[((b ^ (b - 1)) * 0x03F79D71B4CB0A89) >> 58]
}

func (b Bitboard) lastSquare() int {
	b |= b >> 1;
	b |= b >> 2;
	b |= b >> 4;
	b |= b >> 8;
	b |= b >> 16;
	b |= b >> 32;
	return deBruijn[(b * 0x03F79D71B4CB0A89) >> 58];
}

func (b *Bitboard) union(bitBoard Bitboard) *Bitboard {
	*b |= bitBoard

	return b
}

func (b *Bitboard) intersect(bitBoard Bitboard) *Bitboard {
	*b &= bitBoard

	return b
}

func (b *Bitboard) shift(n int) *Bitboard {
	if n > 0 {
		*b = *b << uint(n)
	} else {
		*b = *b >> -uint(n)
	}

	return b
}

func (b Bitboard) toString() string {
	squareNumber := b.firstSquare()
	row := int(squareNumber / 8)
	colNumber := int(squareNumber) % 8

	return columnNames[colNumber] + strconv.Itoa(row + 1)
}


func (b Bitboard)count() int {
	b -= (b >> 1) & 0x5555555555555555
	b = ((b >> 2) & 0x3333333333333333) + (b & 0x3333333333333333)
	b = ((b >> 4) + b) & 0x0F0F0F0F0F0F0F0F

	return int((b * 0x0101010101010101) >> 56)
}

func bbFromInts(row int, col int) Bitboard {
	square := float64(row * 8 + col)
	return Bitboard(uint64(1 << uint(square)))
}

func bbFromString(str string) Bitboard {
	char := string(str[0])
	var row int = 0

	for i, e := range columnNames {
		if e == char {
			row = i
		}
	}


	col, _ := strconv.Atoi(string(str[1]))

	return bbFromInts(row, col - 1)
}