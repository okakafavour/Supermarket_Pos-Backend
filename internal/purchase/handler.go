package purchase

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

// Create Purchase
func (h *Handler) Create(c *gin.Context) {

	var req CreatePurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")

	purchase, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    purchase,
	})
}

// Get All Purchases
func (h *Handler) GetAll(c *gin.Context) {

	purchases, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    purchases,
	})
}

// Get Purchase By ID
func (h *Handler) GetByID(c *gin.Context) {

	id := c.Param("id")

	purchase, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "purchase not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    purchase,
	})
}

// Receive Purchase
func (h *Handler) Receive(c *gin.Context) {

	id := c.Param("id")
	userID := c.GetString("userID")

	if err := h.service.Receive(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Purchase received successfully",
	})
}

// Delete Purchase
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
		"message": "Purchase deleted successfully",
	})
}
