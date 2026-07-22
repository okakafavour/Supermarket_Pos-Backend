package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)

	supplierRepo := supplier.NewRepository(db)
	productRepo := product.NewRepository(db)
	inventoryRepo := inventory.NewRepository(db)

	service := NewService(
		db,
		repo,
		supplierRepo,
		productRepo,
		inventoryRepo,
	)

	handler := NewHandler(service)

	purchases := router.Group("/purchases")
	purchases.Use(middleware.AuthMiddleware())

	{
		purchases.POST(
			"",
			middleware.RequireRole("admin", "manager"),
			handler.Create,
		)

		purchases.GET(
			"",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.GetAll,
		)

		purchases.GET(
			"/:id",
			middleware.RequireRole("admin", "manager", "cashier"),
			handler.GetByID,
		)

		purchases.POST(
			"/:id/receive",
			middleware.RequireRole("admin", "manager"),
			handler.Receive,
		)

		purchases.DELETE(
			"/:id",
			middleware.RequireRole("admin"),
			handler.Delete,
		)
	}
}
