package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Y716/Server-Catur/board"
	"github.com/Y716/Server-Catur/game"
)

func main() {
	Board := board.NewBoard()

	for i := 0; i < 3; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Your move: ")
		if scanner.Scan() {

			input := scanner.Text()
			notations := strings.Split(input, " ")
			game.MovePiece(notations[0], notations[1], &Board)
		}
	}

}
