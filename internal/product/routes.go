package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)

	categoryRepo := category.NewRepository(db)
	supplierRepo := supplier.NewRepository(db)

	service := NewService(
		repo,
		categoryRepo,
		supplierRepo,
	)

	handler := NewHandler(service)

	product := router.Group("/products")
	product.Use(middleware.AuthMiddleware())

	{
		// Create Product
		product.POST(
			"",
			middleware.RequireRole("admin", "manager"),
			handler.Create,
		)

		// Read Products
		product.GET("", handler.GetAll)
		product.GET("/:id", handler.GetByID)
		product.GET(
			"/deleted",
			middleware.RequireRole("admin"),
			handler.GetDeleted,
		)

		// Update Product
		product.PUT(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.Update,
		)

		// Soft Delete
		product.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		// Restore Product
		product.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		// Permanent Delete
		product.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)
	}
}
