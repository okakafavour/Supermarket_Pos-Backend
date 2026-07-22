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
		dashboard.GET(
			"/summary",
			middleware.RequireRole("admin", "manager"),
			handler.GetSummary,
		)
	}
}
