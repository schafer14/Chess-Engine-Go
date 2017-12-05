package maurice

type Piece uint8

const (
	White = iota
	Black
	Pawn
	BlackPawn
	Knight
	BlackKnight
	Bishop
	BlackBishop
	Rook
	BlackRook
	Queen
	BlackQueen
	King
	BlackKing
)

var charMap = [14]string{" ", " ", "P", "p", "N", "n", "B", "b", "R", "r", "Q", "q", "K", "k"}
var symbolMap = [14]string{" ", " ", "♙", "♟", "♘", "♞", "♗", "♝", "♖", "♜", "♕", "♛", "♔", "♚"}

func pawn(color int) Piece {
	return Piece(Pawn + color)
}

func knight(color int) Piece {
	return Piece(Knight + color)
}

func bishop(color int) Piece {
	return Piece(Bishop + color)
}

func rook(color int) Piece {
	return Piece(Rook + color)
}

func queen(color int) Piece {
	return Piece(Queen + color)
}

func king(color int) Piece {
	return Piece(King + color)
}

func (p Piece) char() string {
	return charMap[int(p)]
}

func (p Piece) symbol() string {
	return symbolMap[int(p)]
}
