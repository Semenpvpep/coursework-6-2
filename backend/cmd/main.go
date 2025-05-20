package main

import (
	"backend/internal/delivery/http"
	"backend/internal/entity"
	"backend/pkg/database"
	"log"
	"os"
)

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Принудительная миграция
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Appointment{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	router := http.NewRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
