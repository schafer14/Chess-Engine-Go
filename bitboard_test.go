package maurice

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestBitboard_IsOn(t *testing.T) {
	assert := assert.New(t)

	assert.True(Bitboard(0x01).isOn(0xFF))
	assert.False(Bitboard(0x0100).isOn(0xFF))
	assert.True(Bitboard(0x8000000000000000).isOn(0xFF00000000000000))
	assert.False(Bitboard(0x0800000000000000).isOn(0x0F000000000000000))
}

func TestBitboard_FirstSquare(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, Bitboard(0x01).firstSquare())
	assert.Equal(63, Bitboard(0x8000000000000000).firstSquare())
	assert.Equal(0, Bitboard(0xFF).firstSquare())
	assert.Equal(7, Bitboard(0x80).firstSquare())
	assert.Equal(7, Bitboard(0x0002040810204080).firstSquare())
	assert.Equal(1, Bitboard(0x48AF3D8BCA2C44B6).firstSquare())
	assert.Equal(10, Bitboard(0x48AF3D8BCA2C4400).firstSquare())
}

func TestBitboard_LastSquare(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, Bitboard(0x01).lastSquare())
	assert.Equal(63, Bitboard(0x8000000000000000).lastSquare())
	assert.Equal(7, Bitboard(0xFF).lastSquare())
	assert.Equal(7, Bitboard(0x80).lastSquare())
	assert.Equal(56, Bitboard(0x0102040810204080).lastSquare())
	assert.Equal(62, Bitboard(0x48AF3D8BCA2C44B6).lastSquare())
	assert.Equal(55, Bitboard(0x00AF3D8BCA2C4400).lastSquare())
}

func TestBitboard_Merge(t *testing.T) {
	assert := assert.New(t)

	bb := Bitboard(0x01)
	bb2 := Bitboard(0x03)
	assert.Equal(bb.union(0x02), &bb2)
	assert.Equal(bb, bb2)
	bb3 := Bitboard(0x07)
	assert.Equal(bb.union(0x04), &bb3)
	assert.Equal(bb, bb3)
}

func TestBitboard_ToString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("a1", Bitboard(0x01).toString())
	assert.Equal("h8", Bitboard(0x8000000000000000).toString())
	assert.Equal("d4", Bitboard(0x08000000).toString())
}

func TestBitboard_Count(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, Bitboard(0x01).count())
	assert.Equal(1, Bitboard(0x8000000000000000).count())
	assert.Equal(1, Bitboard(0x08000000).count())
	assert.Equal(2, Bitboard(0x03).count())
	assert.Equal(3, Bitboard(0x07).count())
	assert.Equal(24, Bitboard(0x00AF3D8BCA2C4400).count())
}

func TestBitboard_FromInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Bitboard(0x01), bbFromInts(0, 0))
	assert.Equal(Bitboard(0x8000000000000000), bbFromInts(7, 7))
	assert.Equal(Bitboard(0x08000000), bbFromInts(3, 3))
}

func TestBitboard_Shift(t *testing.T) {
	assert := assert.New(t)

	bb1 := Bitboard(0x08)
	bb2 := Bitboard(0x04)

	assert.Equal(&bb1, bb2.shift(1))
	assert.Equal(bb1, bb2)

	bb1 = Bitboard(0x04)
	bb2 = Bitboard(0x08)
	assert.Equal(&bb1, bb2.shift(-1))
	assert.Equal(bb1, bb2)

}

func TestBitboard_FromString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Bitboard(0x01), bbFromString("a1"))
	assert.Equal(Bitboard(0x80), bbFromString("a8"))
	assert.Equal(Bitboard(0x8000000000000000), bbFromString("h8"))
	assert.Equal(Bitboard(0x08000000), bbFromString("d4"))
}