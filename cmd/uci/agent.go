package main

import (
	"github.com/schafer14/maurice"
	"math"
	"sort"
)

var (
	MIN = math.MinInt64
	MAX = math.MaxInt64
)

func FindMove(position maurice.Position, ch chan string) {
	var bestMove maurice.Move
	var bestScore int = MIN

	for _, move := range position.LegalMoves() {
		child := position.MakeMove(move)
		value := -negamax(child, 4, -(1 + -2*position.Turn()), MIN+1, MAX-1)
		if value > bestScore {
			bestScore = value
			bestMove = move
		}
	}
	ch <- bestMove.ToString()
}

func negamax(node maurice.Position, depth int, color int, alpha int, beta int) int {
	if depth == 0 {
		return evaluate(node) * color
	}

	best := MIN
	candidateMoves := node.PseudoMoves()

	sort.Sort(maurice.Moves(candidateMoves))

	for _, move := range candidateMoves {
		child := node.MakeMove(move)
		value := negamax(child, depth-1, -color, -beta, -alpha)
		value = -value

		if value > best {
			best = value
		}

		if value > alpha {
			alpha = value
		}

		if alpha > beta {
			break
		}
	}

	return best
}

func evaluate(node maurice.Position) int {
	whitePawns := node.PieceBitboards[maurice.Pawn]
	blackPawns := node.PieceBitboards[maurice.BlackPawn]

	pawnScore := evaluatePieceSet(whitePawns, blackPawns, 100, pawnPositionalEval, pawnPositionalEvalFlip)

	whiteKnights := node.PieceBitboards[maurice.Knight]
	blackKnight := node.PieceBitboards[maurice.BlackKnight]

	knightScore := evaluatePieceSet(whiteKnights, blackKnight, 320, knightPositionalEval, knightPositionalEvalFlip)

	whiteBishops := node.PieceBitboards[maurice.Bishop]
	blackBishops := node.PieceBitboards[maurice.BlackBishop]

	bishopScore := evaluatePieceSet(whiteBishops, blackBishops, 330, bishopPositionalEval, bishopPositionalEvalFlip)

	whiteRooks := node.PieceBitboards[maurice.Rook]
	blackRooks := node.PieceBitboards[maurice.BlackRook]

	rookScore := evaluatePieceSet(whiteRooks, blackRooks, 500, rookPositionalEval, rookPositionalEvalFlip)

	whiteQueens := node.PieceBitboards[maurice.Queen]
	blackQueens := node.PieceBitboards[maurice.BlackQueen]

	queenScore := evaluatePieceSet(whiteQueens, blackQueens, 900, queenPositionalEval, queenPositionalEvalFlip)

	whiteKing := node.PieceBitboards[maurice.King]
	blackKing := node.PieceBitboards[maurice.BlackKing]

	kingScore := evaluatePieceSet(whiteKing, blackKing, 20000, kingPositionalEval, kingPositionalEvalFlip)

	return pawnScore + knightScore + bishopScore + rookScore + queenScore + kingScore
}

func evaluatePieceSet(whitePiece maurice.Bitboard, blackPiece maurice.Bitboard, value int, positionalTable [64]int, positionalTableBlack [64]int) int {
	materialCount := (whitePiece.Count() - blackPiece.Count()) * value
	positionCount := 0

	for whitePiece > 0 {
		square := whitePiece & -whitePiece
		whitePiece &= whitePiece - 1
		num := square.FirstSquare()

		positionCount += positionalTable[num]
	}

	for blackPiece > 0 {
		square := blackPiece & -blackPiece
		blackPiece &= blackPiece - 1
		num := square.FirstSquare()

		positionCount -= positionalTableBlack[num]
	}

	return materialCount + positionCount
}
