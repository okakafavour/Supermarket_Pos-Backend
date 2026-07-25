package user

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	users := router.Group("/users")

	users.Use(
		middleware.AuthMiddleware(),
		middleware.RequireRole("admin"),
	)

	users.GET("", handler.GetUsers)
	users.GET("/:id", handler.GetUser)

	users.POST("", handler.CreateUser)

	users.PUT("/:id", handler.UpdateUser)

	users.PATCH("/:id/status", handler.UpdateStatus)

	users.DELETE("/:id", handler.DeleteUser)
}
