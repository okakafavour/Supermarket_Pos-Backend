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
		// ==========================
		// Inventory Summary
		// ==========================

		inventory.GET(
			"/summary",
			handler.Summary,
		)

		// ==========================
		// Stock Management
		// ==========================

		// Admin & Manager can restock inventory
		inventory.POST(
			"/restock",
			middleware.RequireRole("admin", "manager"),
			handler.Restock,
		)

		// Admin & Manager can adjust inventory
		inventory.POST(
			"/adjust",
			middleware.RequireRole("admin", "manager"),
			handler.Adjust,
		)

		// ==========================
		// Inventory Logs
		// ==========================

		// View all inventory logs
		inventory.GET(
			"/logs",
			handler.GetLogs,
		)

		// View logs for a specific product
		inventory.GET(
			"/product/:id",
			handler.GetProductLogs,
		)
	}
}
