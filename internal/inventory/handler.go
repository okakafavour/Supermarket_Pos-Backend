package inventory

import (
	"fmt"
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

// POST /inventory/restock
func (h *Handler) Restock(c *gin.Context) {

	var req RestockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	createdBy := c.GetString("user_id")

	log, err := h.service.Restock(req, createdBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    log,
	})
}

// POST /inventory/adjust
func (h *Handler) Adjust(c *gin.Context) {

	var req AdjustmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	createdBy := c.GetString("user_id")

	log, err := h.service.Adjust(req, createdBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    log,
	})
}

// GET /inventory/logs
func (h *Handler) GetLogs(c *gin.Context) {

	page := 1
	limit := 20

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	search := c.Query("search")
	movement := c.Query("movement")
	reason := c.Query("reason")

	logs, err := h.service.GetLogs(
		page,
		limit,
		search,
		movement,
		reason,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    logs,
	})
}

// GET /inventory/product/:id
func (h *Handler) GetProductLogs(c *gin.Context) {

	productID := c.Param("id")

	logs, err := h.service.GetProductLogs(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    logs,
	})
}

func (h *Handler) Summary(c *gin.Context) {

	summary, err := h.service.GetInventorySummary()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    summary,
	})
}

// GET /inventory/analytics
func (h *Handler) GetInventoryAnalytics(c *gin.Context) {

	analytics, err := h.service.GetInventoryAnalytics()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    analytics,
	})
}
