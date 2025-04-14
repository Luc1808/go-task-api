package main

import (
	"log"

	"github.com/Luc1808/go-task-api/internal/db"
	"github.com/Luc1808/go-task-api/internal/handlers"
	"github.com/Luc1808/go-task-api/internal/models"
	"github.com/Luc1808/go-task-api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := db.InitDB()

	// Auto-migrate models
	db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.ProjectMember{},
		&models.Task{},
		&models.Comment{},
		&models.RefreshToken{},
	)

	authHandler := handlers.NewAuthHandler(db)

	// Clean up expired tokens on start up
	if err := authHandler.ClearExpiredTokens(); err != nil {
		log.Printf("Error clearing expired tokens on startup: %v", err)
	}

	r := gin.Default()

	routes.AuthRoutes(r, db)

	r.Run(":8080")
}
