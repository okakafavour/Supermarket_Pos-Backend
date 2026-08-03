package sale

import (
	"net/http"
	"strconv"

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

// Create Sale
func (h *Handler) Create(c *gin.Context) {

	var req CreateSaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")

	sale, warning, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"success": true,
		"data":    sale,
	}

	if warning != "" {
		response["warning"] = warning
	}

	c.JSON(http.StatusCreated, response)
}

// Get All Sales
func (h *Handler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filter := SaleFilter{
		Page:          page,
		Limit:         limit,
		Search:        c.Query("search"),
		PaymentMethod: c.Query("payment"),
		Status:        c.Query("status"),
		From:          c.Query("from"),
		To:            c.Query("to"),
		SortBy:        c.DefaultQuery("sortBy", "latest"),
		Order:         c.DefaultQuery("order", "desc"),
	}

	sales, pagination, err := h.service.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       sales,
		"pagination": pagination,
	})
}

// Get Sale By ID
func (h *Handler) GetByID(c *gin.Context) {

	id := c.Param("id")

	sale, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "sale not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    sale,
	})
}

// Soft Delete
func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sale deleted successfully",
	})
}

// Restore
func (h *Handler) Restore(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Restore(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sale restored successfully",
	})
}

// Permanent Delete
func (h *Handler) PermanentDelete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.PermanentDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sale permanently deleted",
	})
}

// Get Deleted Sales
func (h *Handler) GetDeleted(c *gin.Context) {

	sales, err := h.service.GetDeleted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    sales,
	})
}

// Dashboard
func (h *Handler) Dashboard(c *gin.Context) {

	data, err := h.service.GetDashboard()
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
