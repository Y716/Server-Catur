package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "net/http"

	"github.com/Y716/Server-Catur/board"
	"github.com/Y716/Server-Catur/game"
	"github.com/Y716/Server-Catur/server"
)

func main() {
	go server.ConnectServer()
	Board := board.NewBoard()
	colorFlag := true //true equals white, false equals black
	for {
		board.PrintBoard(Board)
		scanner := bufio.NewScanner(os.Stdin)
		if colorFlag == true {
			fmt.Print("White's move: ")
		} else {
			fmt.Print("Black's Move: ")
		}

		if scanner.Scan() {

			input := scanner.Text()
			notations := strings.Split(input, " ")
			isValidMove := game.MovePiece(&Board, notations[0], notations[1], colorFlag)

			if isValidMove {
				colorFlag = !colorFlag
			}
		}
	}

}
