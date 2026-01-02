package handler

import (
	"encoding/json"
	"net/http"

	"web-based/internal/service"
)

type HealthResponse struct {
	Status string `json:status`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	result := service.CheckHealth()

	response := HealthResponse{
		Status: result,
	}

	w.Header().Set("Conten-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}
