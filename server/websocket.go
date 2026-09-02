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

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}

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
	GameOver bool
}

type Player struct {
	Conn    *websocket.Conn
	IsWhite bool
}


var gameRoomsMu sync.Mutex
var gameRooms = map[int]*Room{
}

func broadcast(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade:", err)
		return
	}
	defer conn.Close()

	var safeRoom *Room
	var gameID int

	gameRoomsMu.Lock()
	for id, room := range gameRooms{
		if room.Player2 == nil{

			room.Player2 = &Player{
				Conn:    conn,
				IsWhite: false,
			}
			gameID = id

			log.Printf("[ROOM%d] Player2 connected (Black)", gameID )
			safeRoom = room
			break

		}
	}

	if safeRoom == nil{
		gameID = len(gameRooms)
		gameRooms[gameID] = &Room{
			Turn:  true,
			Board: board.NewBoard(),
			GameOver: false,
			Player1: &Player{
				Conn: conn,
				IsWhite: true,
			},
		}
		safeRoom = gameRooms[gameID]

		log.Printf("[ROOM%d] Player1 connected (White)", gameID )
	}
	gameRoomsMu.Unlock()

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
			if safeRoom.Player1 != nil && conn == safeRoom.Player1.Conn {
				if safeRoom.Player2 != nil{

					msg := TurnMessage{
						Type:    "Turn",
						Message: "You win by forfeit",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						break
					}

					err = safeRoom.Player2.Conn.WriteMessage(websocket.TextMessage, messageJson)
					if err != nil {
						break
					}
					log.Printf("[ROOM%d] Player1 disconnected — Player2 wins by forfeit", gameID)
				}
			}else if safeRoom.Player2 != nil && conn == safeRoom.Player2.Conn {
				if safeRoom.Player1 != nil{

					msg := TurnMessage{
						Type:    "Turn",
						Message: "You win by forfeit",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						break
					}

					err = safeRoom.Player1.Conn.WriteMessage(websocket.TextMessage, messageJson)
					if err != nil {
						break
					}
					log.Printf("[ROOM%d] Player2 disconnected — Player1 wins by forfeit", gameID)
				}
			}
			safeRoom.mu.Lock()
			safeRoom.GameOver = true
			safeRoom.mu.Unlock()

			break
		}
		var msg MoveMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("unmarshall:", err)
			continue
		}

		if msg.Type == "Resign"{
			if safeRoom.Player1 != nil && conn == safeRoom.Player1.Conn {
				if safeRoom.Player2 != nil{

					msg := TurnMessage{
						Type:    "Turn",
						Message: "Your opponent has resign. You win by forfeit",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						break
					}

					err = safeRoom.Player2.Conn.WriteMessage(websocket.TextMessage, messageJson)
					if err != nil {
						break
					}
					log.Printf("[ROOM%d] Player1 resigned — Player2 wins", gameID)
				} else {

					msg := TurnMessage{
						Type:    "Turn",
						Message: "Can't resign, must wait for opponent...",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						break
					}

					err = safeRoom.Player1.Conn.WriteMessage(websocket.TextMessage, messageJson)
					if err != nil {
						break
					}
					continue
				}
			}else if safeRoom.Player2 != nil && conn == safeRoom.Player2.Conn {
				if safeRoom.Player1 != nil{

					msg := TurnMessage{
						Type:    "Turn",
						Message: "Your opponent has resign. You win by forfeit",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						break
					}

					err = safeRoom.Player1.Conn.WriteMessage(websocket.TextMessage, messageJson)
					if err != nil {
						break
					}
					log.Printf("[ROOM%d] Player2 resigned — Player1 wins", gameID)
				}
			}
			safeRoom.mu.Lock()
			safeRoom.GameOver = true
			safeRoom.mu.Unlock()

			break
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

				if safeRoom.GameOver{
					msg := TurnMessage{
						Type:    "Turn",
						Message: "Game Over!",
					}

					messageJson, err := json.Marshal(msg)
					if err != nil {
						log.Printf("Error encoding JSON: %s", err)
						continue
					}


					err = safeRoom.BroadcastMessage(websocket.TextMessage, messageJson)
					if err != nil {
						log.Println(err)
					}
					continue
				}else if msg.Type == "Move" {
					safeRoom.mu.Lock()
					Valid := game.MovePiece(safeRoom.Board, msg.From, msg.To, safeRoom.Turn)
					if Valid {

						safeRoom.Turn = !safeRoom.Turn
						log.Printf("[ROOM%d] recv: %s", gameID, msg)

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

						if game.IsCheckmate(*safeRoom.Board, safeRoom.Turn){
							winMessage := ""
							if safeRoom.Turn{
								winMessage = "Black wins!"
							}else{
								winMessage = "White wins!"

							}
							msg := TurnMessage{
								Type:    "Turn",
								Message: winMessage,
							}

							messageJson, err := json.Marshal(msg)
							if err != nil {
								log.Printf("Error encoding JSON: %s", err)
								continue
							}


							err = safeRoom.BroadcastMessage(websocket.TextMessage, messageJson)
							if err != nil {
								log.Println(err)
							}
							safeRoom.GameOver = true
							log.Printf("[ROOM%d] Game over: %s", gameID, msg.Message )

						}else if game.IsStalemate(*safeRoom.Board, safeRoom.Turn){

							msg := TurnMessage{
								Type:    "Turn",
								Message: "Draw!",
							}

							messageJson, err := json.Marshal(msg)
							if err != nil {
								log.Printf("Error encoding JSON: %s", err)
								continue
							}


							err = safeRoom.BroadcastMessage(websocket.TextMessage, messageJson)
							if err != nil {
								log.Println(err)
							}
							safeRoom.GameOver = true
							log.Printf("[ROOM%d] Game over: %s", gameID, msg.Message )
						}
					}else {
						err := SendWarningMessage(conn, "Move Invalid!\n", mt)
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
