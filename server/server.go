package server

import (
	"fmt"
	"net/http"
)


type Response struct{
	Status string `json:"status"`}





func ConnectServer(){
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/echo", echo)
	fmt.Println("Server starting on 8080...")
	http.ListenAndServe(":8080", nil)
}
