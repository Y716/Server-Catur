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
				fmt.Printf("\n%s\n", board.PrintBoard(boardMsg.Board))
			case "Turn":
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
