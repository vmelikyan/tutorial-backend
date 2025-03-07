package main

import (
	"log"
	"net/http"
	"taskapi/internal/api"
	"taskapi/internal/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	handler := api.NewHandler(database)
	router := api.SetupRoutes(handler)

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
