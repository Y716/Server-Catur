package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/Y716/Server-Catur/board"
	"github.com/Y716/Server-Catur/server"
	"github.com/gorilla/websocket"
)
var color string
func main(){
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/broadcast"}
	log.Printf("Connecting to %s...\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)	
	if err != nil{
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, jsonMsg, err := conn.ReadMessage()
			if err != nil{
				log.Printf("Read error: %v", err)
			}

			var msg server.Message
			err = json.Unmarshal(jsonMsg, &msg)
			if err != nil {
				fmt.Println("unmarshall:", err)
				continue
			}

			switch msg.Type{
			case "Board":
				var boardMsg server.BoardMessage
				err = json.Unmarshal(jsonMsg, &boardMsg)
				if err != nil {
					log.Println("unmarshall:", err)
					continue
				}
				var isBlack bool
				if color == "Black"{
					isBlack = true
				}else{
					isBlack = false
				}
				fmt.Printf("\n%s\n", renderBoardForPlayer(boardMsg.Board, isBlack))

			case "Turn":
				var turnMsg server.TurnMessage
				err = json.Unmarshal(jsonMsg, &turnMsg)
				if err != nil {
					log.Println("unmarshall:", err)
					continue
				}
				fmt.Print(turnMsg.Message)	
			case "GameOver":
				var turnMsg server.TurnMessage
				err = json.Unmarshal(jsonMsg, &turnMsg)
				if err != nil {
					log.Println("unmarshall:", err)
					continue
				}
				fmt.Print(turnMsg.Message)	

			case "Color":
				var colorMsg server.ColorMessage
				err = json.Unmarshal(jsonMsg, &colorMsg)
				if err != nil {
					log.Println("unmarshall:", err)
					continue
				}
				color = strings.Split(colorMsg.Message, " ")[3]
				
				fmt.Print(colorMsg.Message)	

			case "Warning":
				var warnMsg server.WarningMessage
				err = json.Unmarshal(jsonMsg, &warnMsg)
				if err != nil {
					log.Println("unmarshall:", err)
					continue
				}
				fmt.Print(warnMsg.Message)	

			}
		} 
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := scanner.Text()
			if text == "Resign"{

				moveStruct := server.Message{
					Type: "Resign",
				}

				moveJson, err := json.Marshal(moveStruct) 
				if err != nil {
					log.Printf("Error Marshal message: %v", err)
					continue	
				}

				
				err = conn.WriteMessage(websocket.TextMessage, moveJson)
				if err != nil {
					log.Printf("Error write message: %v", err)
					continue
				}


			} else { 
				textSlice := strings.Split(text, " ")
				if len( textSlice ) != 2{
					fmt.Println("Hanya menerima format [square1] [square2]")
					fmt.Println("ex: e2 e3")
				}else{ 
					moveStruct := server.MoveMessage{
						Type: "Move",
						From: textSlice[0],
						To: textSlice[1],

					}

					moveJson, err := json.Marshal(moveStruct) 
					if err != nil {
						log.Printf("Error Marshal message: %v", err)
						continue	
					}
					conn.WriteMessage(websocket.TextMessage, moveJson)
				} 
			}

			if err := scanner.Err(); err != nil{
				fmt.Printf("Scanner error: %v", err)
			}
		}

	}
}

func flipBoard(b [8][8]board.Piece) [8][8]board.Piece{
	var flipped [8][8]board.Piece
	
	for i := range b{
		for j := range b{
			flipped[i][j] = b[7-i][7-j]
		}
		

	}
	return flipped
}

func renderBoardForPlayer(b [8][8]board.Piece, isBlack bool) string{

	var boardState strings.Builder
	if isBlack{

		b = flipBoard(b)
		boardState.WriteString("  ")
		for ch := 'H'; ch >= 'A'; ch--{
			fmt.Fprintf(&boardState, "%c ", ch)
		}
		boardState.WriteString("\n")

		for i := range b{
			for j := range b{
				if j == 0{
					fmt.Fprintf(&boardState, "%d ", i+1)
				}
				uniCodePiece := board.GetUniCodePiece(b[i][j])
				fmt.Fprintf(&boardState, "%s ", uniCodePiece)
				if j == 7{
					fmt.Fprintf(&boardState, "%d ", i+1)
				}
			}

			fmt.Fprintln(&boardState, "")

		}

		boardState.WriteString("  ")
		for ch := 'H'; ch >= 'A'; ch--{
			fmt.Fprintf(&boardState, "%c ", ch)
		}
		boardState.WriteString("\n")
	}else{
		boardState.WriteString("  ")
		for ch := 'A'; ch <= 'H'; ch++{
			fmt.Fprintf(&boardState, "%c ", ch)
		}
		boardState.WriteString("\n")

		for i := range b{
			for j := range b{
				if j == 0{
					fmt.Fprintf(&boardState, "%d ", (i*-1)+8)
				}
				uniCodePiece := board.GetUniCodePiece(b[i][j])
				fmt.Fprintf(&boardState, "%s ", uniCodePiece)
				if j == 7{
					fmt.Fprintf(&boardState, "%d ", (i*-1)+8)
				}
			}

			fmt.Fprintln(&boardState, "")

		}

		boardState.WriteString("  ")
		for ch := 'A'; ch <= 'H'; ch++{
			fmt.Fprintf(&boardState, "%c ", ch)
		}
		boardState.WriteString("\n")
	}
	return boardState.String()
}
