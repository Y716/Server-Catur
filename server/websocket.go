package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
// type SafeConnections struct{
// 	connections []*websocket.Conn
// 	mu sync.Mutex
// }

type Room struct{
	Player1 *websocket.Conn
	Player2 *websocket.Conn
	mu sync.Mutex
}

var safeRoom = &Room{}

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
			safeRoom.Player1 = conn
		}else{
			safeRoom.Player2 = conn
		}
	}
	safeRoom.mu.Unlock()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil{
			log.Println("read:", err)
			break
		}
		err = safeRoom.BroadcastMessage(mt, message)
		if err != nil{
			log.Println("write:", err)
		}
		log.Printf("recv: %s", message)
	}

}

func (room *Room) BroadcastMessage(mt int, message []byte) error{
	room.mu.Lock()
	defer room.mu.Unlock()

	err := room.Player1.WriteMessage(mt, message)
	if err != nil {
		return err
	}

	err = room.Player2.WriteMessage(mt, message)
	if err != nil {
		return err
	}

	return nil
}

