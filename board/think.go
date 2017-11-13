package board

import (
	"time"
	"fmt"
)

const MAX = 999999999
const MIN = -999999999

func (b board) Think() string {
	start := time.Now()
	depth := 0
	plan := make([]move, 0)
	humanPlan := make([]string, 0)
	score := 0
	var nodes float64 = 0

	for time.Now().Sub(start).Seconds() < 1 {
		depth = depth + 1
		var n float64 = 0

		plan, score, n = negamax(b, depth, MIN, MAX)
		nodes += n
	}

	for _, move := range plan {
		humanPlan = append(humanPlan, move.toString())
	}

	fmt.Println(humanPlan, score)
	//fmt.Println(nodes / time.Now().Sub(start).Seconds(), "nodes per second")
	return humanPlan[len(humanPlan) - 1]
}

func negamax(b board, depth int, alpha int, beta int) ([]move, int, float64) {
	var nodes float64 = 1

	if depth == 0 || b.isTerminal() {
		return make([]move, 0), b.utility(), nodes
	}

	best := MIN
	bestPlan := make([]move, 0)
	for _, move := range b.PseudoMoves() {
		child := b.Move(move)

		plan, value, n := negamax(child, depth - 1, -beta, -alpha)
		nodes += n
		value = -value

		if value > best {
			plan = append(plan, move)
			best = value
			bestPlan = plan
		}

		if value > alpha {
			alpha = value
		}

		if alpha > beta {
			break
		}
	}

	return bestPlan, best, nodes
}

func (b board) isTerminal() bool {
	if b.whiteKings == 0 || b.blackKings == 0 {
		return true
	} else {
		return false
	}
}

func (b board) utility() int {
	whiteScore := 20000 * (count(b.whiteKings) - count(b.blackKings)) +
					900 * (count(b.whiteQueens) - count(b.blackQueens)) +
					500 * (count(b.whiteRooks) - count(b.blackRooks)) +
					330 * (count(b.whiteBishops) - count(b.blackBishops)) +
					320 * (count(b.whiteKnights) - count(b.blackKnights)) +
					100 * (count(b.whitePawns) - count(b.blackPawns))

	if b.Turn == "w" {
		return whiteScore
	} else {
		return -whiteScore
	}
}