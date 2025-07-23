package app

import (
	"crud-app/internal/app/router"
	"log"
	"net/http"
)

func StartAPIServer() {
	log.Println("Starting API server...")
	err := http.ListenAndServe(":3000", router.InitRouter())
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server running on port: 3000")
}
