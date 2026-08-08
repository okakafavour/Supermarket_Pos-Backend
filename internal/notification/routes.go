package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/okakafavour/supermarket-pos-backend/internal/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(
	router *gin.RouterGroup,
	db *gorm.DB,
) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	notifications := router.Group("/notifications")

	notifications.Use(
		middleware.AuthMiddleware(),
	)

	notifications.GET(
		"",
		handler.GetNotifications,
	)

	notifications.GET(
		"/unread-count",
		handler.GetUnreadCount,
	)

	notifications.PATCH(
		"/read-all",
		handler.MarkAllAsRead,
	)

	notifications.PATCH(
		"/:id/read",
		handler.MarkAsRead,
	)

	notifications.DELETE(
		"",
		handler.DeleteAll,
	)

	notifications.DELETE(
		"/:id",
		handler.Delete,
	)
}
