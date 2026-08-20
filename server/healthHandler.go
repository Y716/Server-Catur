package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct{
	Status string `json:"status"`}

func HealthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{Status: "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConnectServer(){
	http.HandleFunc("/health", HealthHandler)

	fmt.Println("Server starting on 8080...")
	http.ListenAndServe(":8080", nil)
}
