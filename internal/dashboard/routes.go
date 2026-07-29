package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	dashboard := router.Group("/dashboard")
	dashboard.Use(middleware.AuthMiddleware())

	{
		// Dashboard Summary
		dashboard.GET(
			"/summary",
			middleware.RequireRole("admin", "manager"),
			handler.GetSummary,
		)

		// Revenue Chart
		dashboard.GET(
			"/revenue-chart",
			middleware.RequireRole("admin", "manager"),
			handler.GetRevenueChart,
		)

		// Sales Chart
		dashboard.GET(
			"/sales-chart",
			middleware.RequireRole("admin", "manager"),
			handler.GetSalesChart,
		)

		// Recent Sales
		dashboard.GET(
			"/recent-sales",
			middleware.RequireRole("admin", "manager"),
			handler.GetRecentSales,
		)

		// Top Selling Products
		dashboard.GET(
			"/top-products",
			middleware.RequireRole("admin", "manager"),
			handler.GetTopProducts,
		)

		// Low Stock Products
		dashboard.GET(
			"/low-stock",
			middleware.RequireRole("admin", "manager"),
			handler.GetLowStockProducts,
		)
	}
}
