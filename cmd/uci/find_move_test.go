package main

import (
	"fmt"
	"github.com/pkg/profile"
	"github.com/schafer14/maurice"
	"testing"
)

func Test_FindMove(t *testing.T) {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	position := maurice.PositionFromFEN("rnbqkbnr/pp2pppp/3p4/8/3NP3/8/PPP2PPP/RNBQKB1R w KQkq - 0 5")
	ch := make(chan string)
	go FindMove(position, ch)
	mo := <-ch
	fmt.Println("bestmove", mo)
}
