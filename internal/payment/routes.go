package payment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	saleRepo := sale.NewRepository(db)

	service := NewService(repo, saleRepo)

	handler := NewHandler(service)

	payments := router.Group("/payments")
	payments.Use(middleware.AuthMiddleware())

	{
		// Create Payment
		payments.POST(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.Create,
		)

		// Get All Payments
		payments.GET(
			"",
			middleware.RequireRole("admin", "manager"),
			handler.GetAll,
		)

		// Get Payment By ID
		payments.GET(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.GetByID,
		)

		// Soft Delete
		payments.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		// Restore
		payments.PUT(
			"/restore/:id",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		// Permanent Delete
		payments.DELETE(
			"/permanent/:id",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)

		// Get Deleted Payments
		payments.GET(
			"/deleted",
			middleware.RequireRole("admin"),
			handler.GetDeleted,
		)
	}
}
