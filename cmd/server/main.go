package main

import (
	"log"

	"github.com/Luc1808/go-task-api/internal/db"
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
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.AuthRoutes(r, db)

	r.Run(":8080")
}
