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

type Room struct{
	Player1 *Player 
	Player2 *Player 
	mu sync.Mutex
	Board *[8][8]board.Piece
	Turn bool
}

type Player struct{
	Conn *websocket.Conn
	IsWhite bool
}

type Message struct{
	Type string `json:"type"`
	From string `json:"from"`
	To string `json:"to"`
}

var safeRoom = &Room{
	Turn: true,
	Board: board.NewBoard(),
}

func broadcast(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w,r, nil)
	if err != nil{
		log.Print("Upgrade:", err)
		return
	}
	defer conn.Close()
	safeRoom.mu.Lock()
	if safeRoom.Player1 != nil && safeRoom.Player2 != nil{
		conn.WriteMessage(websocket.TextMessage, []byte( "Room is full. Please come again later!" ))	
		safeRoom.mu.Unlock()
		return
	}else{
		if safeRoom.Player1 == nil{
			safeRoom.Player1 = &Player{
				Conn: conn,
				IsWhite: true,
			}
		}else{
			safeRoom.Player2 = &Player{
				Conn: conn,
				IsWhite: false,
			}
		}
	}
	safeRoom.mu.Unlock()
	

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil{
			log.Println("read:", err)
			break
		}
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil{
			log.Println("unmarshall:", err)
		}
		
		if (safeRoom.Player1.Conn == conn && safeRoom.Turn == safeRoom.Player1.IsWhite || safeRoom.Player2.Conn == conn && safeRoom.Turn == safeRoom.Player2.IsWhite){
			if (safeRoom.Player1== nil || safeRoom.Player2 == nil){
				if safeRoom.Player1 != nil{
					err := safeRoom.Player1.Conn.WriteMessage(mt, []byte("Need one more player!"))
					if err != nil {
						log.Println("write:", err)
						continue
					}
				}else if safeRoom.Player2 != nil{
					err := safeRoom.Player2.Conn.WriteMessage(mt, []byte("Need one more player!"))
					if err != nil {
						log.Println("write:", err)
						continue

					}
				}
			}else{
				if msg.Type == "move" {
					safeRoom.mu.Lock()
					Valid := game.MovePiece(safeRoom.Board, msg.From, msg.To, safeRoom.Turn)
					if Valid {
						safeRoom.mu.Unlock()

						safeRoom.Turn = !safeRoom.Turn
						log.Printf("recv: %s", msg)
						boardState := board.PrintBoard(*safeRoom.Board)
						log.Printf("%s", boardState)

						// err = safeRoom.BroadcastMessage(mt, []byte(boardState)) 
						// if err != nil{
						// 	log.Println("write:", err)
						// 	return
						// }				
						err = safeRoom.BroadcastMessage(mt, message) 
						if err != nil{
							log.Println("write:", err)
						}
					}else{
						conn.WriteMessage(websocket.TextMessage, []byte("Move invalid!"))
						safeRoom.mu.Unlock()
					}
				}

		}
		}else {
			conn.WriteMessage(websocket.TextMessage, []byte("Not your turn!"))	
		}


	}

}

func (room *Room) BroadcastMessage(mt int, message []byte) error{
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

