package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	category := router.Group("/categories")
	category.Use(middleware.AuthMiddleware())

	{
		// Create
		category.POST(
			"",
			middleware.RequireRole("admin"),
			handler.Create,
		)

		// Read
		category.GET("", handler.GetAll)
		category.GET("/deleted", middleware.RequireRole("admin"), handler.GetDeleted)
		category.GET("/:id", handler.GetByID)

		// Update
		category.PUT(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.Update,
		)

		// Soft Delete
		category.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		// Restore
		category.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		// Permanent Delete
		category.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)
	}
}
