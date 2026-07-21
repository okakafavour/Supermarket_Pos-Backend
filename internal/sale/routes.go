package sale

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	// Repositories
	repo := NewRepository(db)
	productRepo := product.NewRepository(db)
	inventoryRepo := inventory.NewRepository(db)

	// Service
	service := NewService(
		repo,
		productRepo,
		inventoryRepo,
	)

	handler := NewHandler(service)

	sales := router.Group("/sales")
	sales.Use(middleware.AuthMiddleware())

	{
		// Create Sale
		sales.POST(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.Create,
		)

		// Get All Sales
		sales.GET(
			"",
			middleware.RequireRole("admin", "manager"),
			handler.GetAll,
		)

		// Get Deleted Sales
		// MUST come before "/:id"
		sales.GET(
			"/deleted",
			middleware.RequireRole("admin"),
			handler.GetDeleted,
		)

		// Get Sale By ID
		sales.GET(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.GetByID,
		)

		// Restore Sale
		sales.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		// Soft Delete
		sales.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		// Permanent Delete
		sales.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)
	}
}
