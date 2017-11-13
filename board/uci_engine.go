package board

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var startpos string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func Uci() error {
	var b board = board{}

	reader := bufio.NewReader(os.Stdin)

		for {
			text, _ := reader.ReadString('\n');
			args := strings.Split(text, " ")

			switch cmd := strings.TrimSpace(args[0]); cmd {
			case "uci":
				fmt.Println("id name Maurice");
				fmt.Println("id author Banner B. Schafer");
				fmt.Println("uciok");
			case "setoption":
			case "isready":
				fmt.Println("readyok")
			case "position":
				if strings.TrimSpace(args[1]) == "startpos" {
					b = FromFEN(startpos)
					for i := 3; i < len(args); i++ {
						m := strings.TrimSpace(args[i])
						b = b.HumanFriendlyMove(m)
					}
				} else {
					fmt.Println("TODO")
					b = FromFEN(startpos)
				}
			case "go":
				fmt.Println("bestmove", b.Think())
			case "quit":
				break
			}
		}
}
