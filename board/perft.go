package board

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"time"
)

func Perft() {
	b, err := ioutil.ReadFile( "src/chess/board/perft.epd")

	if err != nil {
		fmt.Println(err)
	}

	positions := strings.Split(string(b), "\n")
	positive := true

	for p := 0; p < len(positions) && positive; p ++ {
		position := positions[p]

		parts := strings.Split(position, ";")
		fen := parts[0]

		fmt.Println(p, "Checking board", fen)

		board := FromFEN(fen)

		for i := 1; i < len(parts) && positive; i ++ {
			part := parts[i]
			expectedNumMoves , _ := strconv.Atoi(strings.Split(part, " ")[1])

			start := time.Now()
			actualNumMoves := perft(board, i)
			end := time.Now()

			if expectedNumMoves != actualNumMoves {
				positive = false
			}

			fmt.Println(expectedNumMoves == actualNumMoves, i, expectedNumMoves, actualNumMoves, end.Sub(start))
		}
		fmt.Println("----------------")
	}

}

func (b board) Divide (d int) error {
	sum := 0
	for _, m := range b.LegalMoves() {
		nb := b.Move(m)
		num := perft(nb, d - 1)
		sum += num
		fmt.Println(m.toString(), num)
	}
	fmt.Println("Moves:", sum)

	return nil
}

func perft(b board, d int) int {
	if d == 0 {
		return 1
	}
	nodes := 0

	for _, m := range b.PseudoMoves() {
		nb := b.Move(m)
		if !nb.isInCheck() {
			nodes += perft(nb, d - 1)
		}
	}

	return nodes
}