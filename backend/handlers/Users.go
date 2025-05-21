package handlers

import (
	"encoding/json"
	"net/http"
	Models "social-network/backend/models"
	"time"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create response data
	response := Models.AboutResponse{
		Message:   "Welcome to our Social Network",
		Version:   "1.0.0",
		Status:    "active",
		Timestamp: time.Now().Unix(),
	}

	// Send JSON response
	json.NewEncoder(w).Encode(response)
}
