package maurice

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPositionFromFEN(t *testing.T) {
	assert := assert.New(t)

	p := PositionFromFEN(initialFEN)



	assert.Equal(Bitboard(0xFFFF00000000FFFF), p.pieceBitboards[0])
	assert.Equal(Bitboard(0xFFFF), p.pieceBitboards[1])
	assert.Equal(Bitboard(0xFFFF000000000000), p.pieceBitboards[2])
	assert.Equal(Bitboard(0xFF00), p.pieceBitboards[3])
	assert.Equal(Bitboard(0x00FF000000000000), p.pieceBitboards[4])
	assert.Equal(Bitboard(0x42), p.pieceBitboards[5])
	assert.Equal(Bitboard(0x4200000000000000), p.pieceBitboards[6])
	assert.Equal(Bitboard(0x24), p.pieceBitboards[7])
	assert.Equal(Bitboard(0x2400000000000000), p.pieceBitboards[8])
	assert.Equal(Bitboard(0x81), p.pieceBitboards[9])
	assert.Equal(Bitboard(0x8100000000000000), p.pieceBitboards[10])
	assert.Equal(Bitboard(0x08), p.pieceBitboards[11])
	assert.Equal(Bitboard(0x0800000000000000), p.pieceBitboards[12])
	assert.Equal(Bitboard(0x10), p.pieceBitboards[13])
	assert.Equal(Bitboard(0x1000000000000000), p.pieceBitboards[14])

	assert.Equal(0, p.color)

	assert.Equal([4]bool{ true, true, true, true }, p.castlingRights)
	assert.Equal(Bitboard(0), p.enPassent)
	assert.Equal(uint8(0), p.count50)

	p = PositionFromFEN("k7/8/8/8/8/8/8/K7 b - d4 56 1")

	assert.Equal(1, p.color)

	assert.Equal([4]bool{ false, false, false, false}, p.castlingRights)
	assert.Equal(Bitboard(0x08000000), p.enPassent)
	assert.Equal(uint8(56), p.count50)


	p = PositionFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1")
	p.Draw()
	for _, m := range p.pseudoMoves() {
		fmt.Print(m.toString(), " ")
	}
}
