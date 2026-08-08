package settings

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	settings := router.Group("/settings")

	// Anyone who is authenticated can read store settings.
	settings.GET(
		"",
		middleware.AuthMiddleware(),
		handler.GetSettings,
	)

	// Only admin can update store settings.
	settings.PUT(
		"",
		middleware.AuthMiddleware(),
		middleware.RequireRole("admin"),
		handler.UpdateSettings,
	)
}
