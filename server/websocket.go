package server

import (
	"log"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)
	
var upgrader = websocket.Upgrader{}
type SafeConnections struct{
	connections []*websocket.Conn
	mu sync.Mutex
}

var safeConnect = &SafeConnections{}

func broadcast(w http.ResponseWriter, r *http.Request){

	conn, err := upgrader.Upgrade(w,r, nil)
	if err != nil{
		log.Print("Upgrade:", err)
		return
	}
	safeConnect.mu.Lock()
	safeConnect.connections = append(safeConnect.connections, conn)
	safeConnect.mu.Unlock()
	defer conn.Close()
	
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil{
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		safeConnect.mu.Lock()
		for _, c := range safeConnect.connections {
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
		safeConnect.mu.Unlock()
	}

}
