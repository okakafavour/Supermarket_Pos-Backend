package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	reports := router.Group("/reports")
	reports.Use(middleware.AuthMiddleware())

	{
		reports.GET(
			"/sales",
			middleware.RequireRole("admin", "manager"),
			handler.GetSalesSummary,
		)

		reports.GET(
			"/sales/daily",
			middleware.RequireRole("admin", "manager"),
			handler.GetDailySalesReport,
		)
	}
}
