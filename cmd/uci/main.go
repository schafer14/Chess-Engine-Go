package main

import (
	"bufio"
	"fmt"
	"github.com/schafer14/maurice"
	"os"
	"strings"
)

func main() {
	position := maurice.InitialPosition()

	reader := bufio.NewReader(os.Stdin)

REPL:
	for {
		text, _ := reader.ReadString('\n')
		args := strings.Split(text, " ")

		for i, arg := range args {
			args[i] = strings.TrimSpace(arg)
		}

		switch cmd, args := args[0], args[1:]; cmd {
		case "quit":
			break REPL
		case "uci":
			fmt.Println("id name Maurice")
			fmt.Println("id author Banner B. Schafer")
			fmt.Println("uciok")
		case "isready":
			fmt.Println("readyok")
		case "position":
			if subcmd, args := args[0], args[1:]; subcmd == "startpos" {
				args = args[1:]
				position = maurice.InitialPosition()
				for _, moveStr := range args {
					position.Move(moveStr)
				}
			}
		case "go":
			ch := make(chan string)
			go FindMove(position, ch)
			mo := <-ch
			fmt.Println("bestmove", mo)
		case "d":
			position.Draw()
		case "m":
			fmt.Println(position.HumanFriendlyMoves())
		default:
			fmt.Println("Next")
		}
	}
}
