package server

import (
	"fmt"
	"net/http"
	"os"
)


type Response struct{
	Status string `json:"status"`}

func ConnectServer(){
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/broadcast", broadcast)
	http.HandleFunc("/games", GamesHandler)
	fmt.Printf("Server starting on %s...\n", PORT)
	PORT = ":" + PORT
	http.ListenAndServe(PORT, nil)
}
