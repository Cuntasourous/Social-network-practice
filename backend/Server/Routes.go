package Server

import (
	"net/http"
	"social-network/backend/handlers"
)

func Routes() {
	http.HandleFunc("/", SpaHandler)
	http.HandleFunc("/api/about", handlers.AboutHandler)
}
