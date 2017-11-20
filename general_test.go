package maurice

import (
	"testing"
	"github.com/pkg/profile"
)


func TestStuff(t *testing.T) {
	perft()
}

func BenchmarkPawnMoves(b *testing.B) {
	defer profile.Start().Stop()

	position := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	p := PositionFromFEN(position)


	p.Perft(5)
}