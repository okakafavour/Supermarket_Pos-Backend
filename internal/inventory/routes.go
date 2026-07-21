package inventory

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	inventory := router.Group("/inventory")
	inventory.Use(middleware.AuthMiddleware())

	{
		// Admin & Manager can restock
		inventory.POST(
			"/restock",
			middleware.RequireRole("admin", "manager"),
			handler.Restock,
		)

		// Admin & Manager can adjust stock
		inventory.POST(
			"/adjust",
			middleware.RequireRole("admin", "manager"),
			handler.Adjust,
		)

		// All authenticated users can view logs
		inventory.GET(
			"/logs",
			handler.GetLogs,
		)

		inventory.GET(
			"/product/:id",
			handler.GetProductLogs,
		)
	}
}
