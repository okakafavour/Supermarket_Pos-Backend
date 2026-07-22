package returns

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)

	saleRepo := sale.NewRepository(db)

	productRepo := product.NewRepository(db)

	inventoryRepo := inventory.NewRepository(db)

	service := NewService(
		db,
		repo,
		saleRepo,
		productRepo,
		inventoryRepo,
	)

	handler := NewHandler(service)

	returns := router.Group("/returns")
	returns.Use(middleware.AuthMiddleware())

	{
		returns.POST(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.Create,
		)

		returns.GET(
			"",
			middleware.RequireRole("admin", "manager"),
			handler.GetAll,
		)

		returns.GET(
			"/deleted",
			middleware.RequireRole("admin"),
			handler.GetDeleted,
		)

		returns.GET(
			"/:id",
			middleware.RequireRole("admin", "manager"),
			handler.GetByID,
		)

		returns.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)

		returns.PATCH(
			"/:id/restore",
			middleware.RequireRole("admin"),
			handler.Restore,
		)

		returns.DELETE(
			"/:id/permanent",
			middleware.RequireRole("admin"),
			handler.PermanentDelete,
		)
	}
}
