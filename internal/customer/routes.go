package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	customers := router.Group("/customers")
	customers.Use(middleware.AuthMiddleware())

	{
		customers.POST(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.Create,
		)

		customers.GET(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.GetAll,
		)

		customers.GET(
			"/search",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.Search,
		)

		// IMPORTANT: static routes BEFORE :id
		customers.GET(
			"/deleted",
			middleware.RequireRole("admin", "manager"),
			handler.GetDeleted,
		)

		customers.GET(
			"/:id",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.GetByID,
		)

		customers.PUT(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.Update,
		)

		customers.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		customers.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		customers.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)

		customers.POST(
			"/:id/loyalty",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.AddLoyaltyPoints,
		)
	}
}
