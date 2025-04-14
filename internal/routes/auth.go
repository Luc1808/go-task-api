package routes

import (
	"github.com/Luc1808/go-task-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	handler := handlers.NewAuthHandler(db)

	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/refresh", handler.RefreshHandler)
	}
}
