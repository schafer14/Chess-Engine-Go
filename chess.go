package main


import (
	b "chess/board"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var (
	position = "4k2r/8/8/8/8/8/8/4K3 w k - 0 1 "
)

func main() {

	if os.Args[1] == "perft" {
		fmt.Println("Running Perft Test\n")
		b.Perft()
	}

	if os.Args[1] == "explore" {
		board := b.FromFEN(position)
		reader := bufio.NewReader(os.Stdin)

		REPL:
			for {
				text, _ := reader.ReadString('\n');
				args := strings.Split(text, " ")

				switch cmd := strings.TrimSpace(args[0]); cmd {
				case "q":
					break REPL
				case "m":
					fmt.Println(board.HumanFriendlyMoves())
				case "d":
					board.Draw()
				case "r":
					board = b.FromFEN(position)
				case "divide":
					i, _ := strconv.Atoi(strings.TrimSpace(args[1]))
					board.Divide(i)
				case "s":
					fmt.Println(b.ToFEN(board))
				default:
					fmt.Println("Moving")
					board = board.HumanFriendlyMove(text)
				}
			}
	}
}