package Server

import (
	"fmt"
	"net/http"
	// "social-network/backend/handlers"
)

func Server() {
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
