package server

import (
	"net/http"
	"encoding/json"
)
func HealthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{Status: "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
