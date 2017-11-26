package main

import (
	"github.com/pkg/profile"
	"github.com/schafer14/maurice"
	"fmt"
	"time"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	position := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	p := maurice.PositionFromFEN(position)

	t := time.Now()
	x := p.Perft(6)
	e := time.Since(t)
	fmt.Println(x, e)
}

