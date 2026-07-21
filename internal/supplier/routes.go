package supplier

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	supplier := router.Group("/suppliers")
	supplier.Use(middleware.AuthMiddleware())

	{
		supplier.POST(
			"",
			middleware.RequireRole("admin"),
			handler.Create,
		)

		supplier.GET("", handler.GetAll)

		supplier.GET(
			"/deleted",
			middleware.RequireRole("admin"),
			handler.GetDeleted,
		)

		supplier.GET("/:id", handler.GetByID)

		supplier.PUT(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.Update,
		)

		supplier.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		supplier.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		supplier.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)
	}
}
