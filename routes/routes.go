package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/okakafavour/supermarket-pos-backend/internal/auth"
	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/dashboard"
	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/reports"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "API is running",
		})
	})

	// Auth
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.GET("/profile", middleware.AuthMiddleware(), authHandler.Profile)
	}

	// Category
	category.RegisterRoutes(api, db)
	supplier.RegisterRoutes(api, db)
	product.RegisterRoutes(api, db)
	inventory.RegisterRoutes(api, db)
	sale.RegisterRoutes(api, db)
	payment.RegisterRoutes(api, db)
	dashboard.RegisterRoutes(api, db)
	reports.RegisterRoutes(api, db)
}
