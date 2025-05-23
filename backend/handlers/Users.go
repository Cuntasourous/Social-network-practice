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

// func RegisterUser(w http.ResponseWriter, r *http.Request) {
// 	models.EnableCors(&w)
// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	log.Println("Received registration request")

// 	// Get the form data
// 	email := r.FormValue("email")
// 	password := r.FormValue("password")
// 	firstName := r.FormValue("firstName")
// 	lastName := r.FormValue("lastName")
// 	dateOfBirth := r.FormValue("dateOfBirth")
// 	nickname := r.FormValue("nickname")
// 	about := r.FormValue("about")
// 	isPrivate := r.FormValue("isPrivate") == "Private"

// 	// Begin transaction
// 	tx, err := db.DB.Begin()
// 	if err != nil {
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		log.Println("Error starting transaction:", err)
// 		return
// 	}

// 	// Prepare statement
// 	stmt, err := tx.Prepare(`
// 		INSERT INTO users (
// 			id, email, password, first_name, last_name,
// 			date_of_birth, avatar, nickname, about, is_private
// 		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`)
// 	if err != nil {
// 		tx.Rollback()
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		log.Println("Error preparing statement:", err)
// 		return
// 	}
// 	defer stmt.Close()

// 	// Send the response
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message":       "User registered successfully",
// 		"userId":        userID,
// 		"session_token": sessionToken,
// 	})

// }
