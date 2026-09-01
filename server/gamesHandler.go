package server

import (
	"encoding/json"
	"net/http"
)


type GameSummary struct {
	IdRoom int `json:"idroom"`
	IsPlaying bool `json:"isplaying"`
	Turn bool `json:"turn"`
	IsGameOver bool `json:"isgameover"`
}

func GamesHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var summaries []GameSummary
	gameRoomsMu.Lock()
	defer gameRoomsMu.Unlock()

	for id, room := range gameRooms{

		room.mu.Lock()
		response := GameSummary{
			IdRoom: id,
			IsPlaying: room.Player2 != nil,
			Turn: room.Turn,
			IsGameOver: room.GameOver,
		}

		summaries = append(summaries, response)

		room.mu.Unlock()
	}

	if err := json.NewEncoder(w).Encode(summaries); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
