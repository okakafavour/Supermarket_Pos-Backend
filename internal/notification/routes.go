package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"github.com/okakafavour/supermarket-pos-backend/internal/user"
	"gorm.io/gorm"
)

func RegisterRoutes(
	router *gin.RouterGroup,
	db *gorm.DB,
) {
	// Repositories
	repo := NewRepository(db)
	userRepo := user.NewRepository(db)

	// Service
	service := NewService(
		repo,
		userRepo,
	)

	// Handler
	handler := NewHandler(service)

	// Notification routes
	notifications := router.Group("/notifications")

	notifications.Use(
		middleware.AuthMiddleware(),
	)

	// GET /notifications
	notifications.GET(
		"",
		handler.GetNotifications,
	)

	// GET /notifications/unread-count
	notifications.GET(
		"/unread-count",
		handler.GetUnreadCount,
	)

	// PATCH /notifications/read-all
	notifications.PATCH(
		"/read-all",
		handler.MarkAllAsRead,
	)

	// PATCH /notifications/:id/read
	notifications.PATCH(
		"/:id/read",
		handler.MarkAsRead,
	)

	// DELETE /notifications
	notifications.DELETE(
		"",
		handler.DeleteAll,
	)

	// DELETE /notifications/:id
	notifications.DELETE(
		"/:id",
		handler.Delete,
	)
}
