package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", handler.Profile)
}
