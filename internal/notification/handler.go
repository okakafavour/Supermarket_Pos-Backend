package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// GET /notifications
func (h *Handler) GetNotifications(c *gin.Context) {

	userID := c.GetString("user_id")

	notifications, err := h.service.GetNotifications(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    notifications,
	})
}

// GET /notifications/unread-count
func (h *Handler) GetUnreadCount(c *gin.Context) {

	userID := c.GetString("user_id")

	count, err := h.service.GetUnreadCount(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"count": count,
		},
	})
}

// PATCH /notifications/:id/read
func (h *Handler) MarkAsRead(c *gin.Context) {

	userID := c.GetString("user_id")
	id := c.Param("id")

	if err := h.service.MarkAsRead(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Notification marked as read",
	})
}

// PATCH /notifications/read-all
func (h *Handler) MarkAllAsRead(c *gin.Context) {

	userID := c.GetString("user_id")

	if err := h.service.MarkAllAsRead(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "All notifications marked as read",
	})
}

// DELETE /notifications/:id
func (h *Handler) Delete(c *gin.Context) {

	userID := c.GetString("user_id")
	id := c.Param("id")

	if err := h.service.Delete(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Notification deleted",
	})
}

// DELETE /notifications
func (h *Handler) DeleteAll(c *gin.Context) {

	userID := c.GetString("user_id")

	if err := h.service.DeleteAll(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "All notifications deleted",
	})
}

// POST /notifications/test
func (h *Handler) CreateTestNotification(c *gin.Context) {
	userID := c.GetString("user_id")

	err := h.service.Create(
		userID,
		LowStockNotification,
		"Low Stock Alert",
		"Test product is below the minimum stock level. Only 3 left.",
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Test notification created",
	})
}
