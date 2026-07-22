package reports

import (
	"net/http"
	"time"

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

// Sales Summary
func (h *Handler) GetSalesSummary(c *gin.Context) {

	report, err := h.service.GetSalesSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}

// Daily Sales Report
func (h *Handler) GetDailySalesReport(c *gin.Context) {

	report, err := h.service.GetDailySalesReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}

// Monthly Sales Report
func (h *Handler) GetMonthlySalesReport(c *gin.Context) {

	report, err := h.service.GetMonthlySalesReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}

// Low Stock Report
func (h *Handler) GetLowStockReport(c *gin.Context) {

	report, err := h.service.GetLowStockReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}

// Top Selling Products
func (h *Handler) GetTopSellingProducts(c *gin.Context) {

	report, err := h.service.GetTopSellingProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}

// Date Range Report
func (h *Handler) GetDateRangeReport(c *gin.Context) {

	startStr := c.Query("start")
	endStr := c.Query("end")

	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "start and end query parameters are required (YYYY-MM-DD)",
		})
		return
	}

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid start date format",
		})
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid end date format",
		})
		return
	}

	// Include the entire end day
	end = end.Add(24 * time.Hour)

	report, err := h.service.GetDateRangeReport(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})
}
