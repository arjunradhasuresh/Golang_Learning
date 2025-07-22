package app

import (
	"log"
	"net/http"
)

func StartAPIServer() {
	log.Println("Starting API server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server running on port: 8080")
}
