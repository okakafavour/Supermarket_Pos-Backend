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

	filter := InventoryLogFilter{
		Page:      1,
		Limit:     20,
		Search:    c.Query("search"),
		Movement:  c.Query("movement"),
		Reason:    c.Query("reason"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	}

	if page := c.Query("page"); page != "" {
		fmt.Sscanf(page, "%d", &filter.Page)
	}

	if limit := c.Query("limit"); limit != "" {
		fmt.Sscanf(limit, "%d", &filter.Limit)
	}

	logs, err := h.service.GetLogs(filter)
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

// GET /inventory/export/csv
func (h *Handler) ExportCSV(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, gin.H{
		"success": false,
		"message": "CSV export not implemented yet",
	})
}

// GET /inventory/export/pdf
func (h *Handler) ExportPDF(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, gin.H{
		"success": false,
		"message": "PDF export not implemented yet",
	})
}
