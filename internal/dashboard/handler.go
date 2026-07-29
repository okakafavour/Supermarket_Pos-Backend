package dashboard

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

/*
========================================
Dashboard Summary
========================================
*/

func (h *Handler) GetSummary(c *gin.Context) {

	summary, err := h.service.GetSummary()
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

/*
========================================
Revenue Chart
========================================
*/

func (h *Handler) GetRevenueChart(c *gin.Context) {

	data, err := h.service.GetRevenueChart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

/*
========================================
Sales Chart
========================================
*/

func (h *Handler) GetSalesChart(c *gin.Context) {

	data, err := h.service.GetSalesChart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

/*
========================================
Recent Sales
========================================
*/

func (h *Handler) GetRecentSales(c *gin.Context) {

	data, err := h.service.GetRecentSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

/*
========================================
Top Selling Products
========================================
*/

func (h *Handler) GetTopProducts(c *gin.Context) {

	data, err := h.service.GetTopProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

/*
========================================
Low Stock Products
========================================
*/

func (h *Handler) GetLowStockProducts(c *gin.Context) {

	data, err := h.service.GetLowStockProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
