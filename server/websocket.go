package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/Y716/Server-Catur/board"
	"github.com/Y716/Server-Catur/game"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// type SafeConnections struct{
// 	connections []*websocket.Conn
// 	mu sync.Mutex
// }

type Room struct {
	Player1 *Player
	Player2 *Player
	mu      sync.Mutex
	Board   *[8][8]board.Piece
	Turn    bool
}

type Player struct {
	Conn    *websocket.Conn
	IsWhite bool
}

var safeRoom = &Room{
	Turn:  true,
	Board: board.NewBoard(),
}

func broadcast(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade:", err)
		return
	}
	defer conn.Close()
	safeRoom.mu.Lock()
	if safeRoom.Player1 != nil && safeRoom.Player2 != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Room is full. Please come again later!"))
		safeRoom.mu.Unlock()
		return
	} else {
		if safeRoom.Player1 == nil {
			safeRoom.Player1 = &Player{
				Conn:    conn,
				IsWhite: true,
			}
		} else {
			safeRoom.Player2 = &Player{
				Conn:    conn,
				IsWhite: false,
			}
		}
	}
	safeRoom.mu.Unlock()
	boardMessage := BoardMessage{
		Type:  "Board",
		Board: *safeRoom.Board,
	}

	boardJSON, err := json.Marshal(boardMessage)
	if err != nil {
		log.Printf("Error encoding JSON: %s", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, boardJSON)
	if err != nil {
		log.Println(err)
	}

	if safeRoom.Player1 != nil && safeRoom.Player2 != nil {
		safeRoom.SendToPlayer()
	}

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var msg MoveMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("unmarshall:", err)
			continue
		}

		if safeRoom.Player1.Conn == conn && safeRoom.Turn == safeRoom.Player1.IsWhite || safeRoom.Player2.Conn == conn && safeRoom.Turn == safeRoom.Player2.IsWhite {
			if safeRoom.Player1 == nil || safeRoom.Player2 == nil {
				if safeRoom.Player1 != nil {
					err := SendWarningMessage(safeRoom.Player1.Conn, "Need one more player!", mt)
					if err != nil {
						log.Printf("Error Sending Warn Message: %v", err)
					}
				} else if safeRoom.Player2 != nil {
					err := SendWarningMessage(safeRoom.Player2.Conn, "Need one more player!", mt)
					if err != nil {
						log.Printf("Error Sending Warn Message: %v", err)
					}
				}
			} else {
				if msg.Type == "Move" {
					safeRoom.mu.Lock()
					Valid := game.MovePiece(safeRoom.Board, msg.From, msg.To, safeRoom.Turn)
					if Valid {

						safeRoom.Turn = !safeRoom.Turn
						log.Printf("recv: %s", msg)
						boardState := board.PrintBoard(*safeRoom.Board)
						log.Printf("%s", boardState)

						// err = safeRoom.BroadcastMessage(mt, []byte(boardState))
						// if err != nil{
						// 	log.Println("write:", err)
						// 	return
						// }
						boardMessage.Board = *safeRoom.Board
						boardJSON, err := json.Marshal(boardMessage)
						if err != nil {
							log.Printf("Error encoding JSON: %s", err)
							continue
						}

						safeRoom.mu.Unlock()
						err = safeRoom.BroadcastMessage(websocket.TextMessage, boardJSON)
						if err != nil {
							log.Println(err)
						}
						safeRoom.SendToPlayer()

					} else {
						err := SendWarningMessage(conn, "Move Invalid!\n", mt)
						if err != nil {
							log.Printf("Error Sending Warn Message: %v", err)
						}

						if err != nil {
							log.Printf("Error Sending Warn Message: %v", err)
						}
						safeRoom.mu.Unlock()
					}
				}

			}
		} else {
			err := SendWarningMessage(conn, "Not your turn!", mt)
			if err != nil {
				log.Printf("Error Sending Warn Message: %v", err)
			}
		}

	}

}

func (room *Room) SendToPlayer() error {
	players := []*Player{room.Player1, room.Player2}
	room.mu.Lock()
	defer room.mu.Unlock()
	for _, player := range players {
		if player == nil {
			return nil
		}
		if player.IsWhite == room.Turn {
			msg := TurnMessage{
				Type:    "Turn",
				Message: "Your Turn: ",
			}

			messageJson, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error encoding JSON: %s", err)
				return err
			}

			err = player.Conn.WriteMessage(websocket.TextMessage, messageJson)
			if err != nil {
				return err
			}
		} else {
			msg := TurnMessage{
				Type:    "Turn",
				Message: "Waiting Opponent's Turn...",
			}

			messageJson, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error encoding JSON: %s", err)
				return err
			}

			err = player.Conn.WriteMessage(websocket.TextMessage, messageJson)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (room *Room) BroadcastMessage(mt int, message []byte) error {
	room.mu.Lock()
	defer room.mu.Unlock()
	err := room.Player1.Conn.WriteMessage(mt, message)
	if err != nil {
		return err
	}

	err = room.Player2.Conn.WriteMessage(mt, message)
	if err != nil {
		return err
	}

	return nil
}

func SendWarningMessage(conn *websocket.Conn, msg string, mt int) error{
	msgStruct := WarningMessage{
		Type:    "Warning",
		Message: msg,
	}

	messageJson, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(mt, messageJson)
	if err != nil {
		return err
	}
	return nil
}
