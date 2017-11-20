package main

import (
	"github.com/pkg/profile"
	"github.com/schafer14/maurice"
	"fmt"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	position := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	p := maurice.PositionFromFEN(position)


	x := p.Perft(5)
	fmt.Println(x)
}

