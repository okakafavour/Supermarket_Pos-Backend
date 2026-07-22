package returns

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

// Create Return
func (h *Handler) Create(c *gin.Context) {

	var req CreateReturnRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userID := c.GetString("user_id")

	ret, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    ret,
	})
}

// Get All Returns
func (h *Handler) GetAll(c *gin.Context) {

	returns, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    returns,
	})
}

// Get Return By ID
func (h *Handler) GetByID(c *gin.Context) {

	ret, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "return not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ret,
	})
}

// Get Deleted Returns
func (h *Handler) GetDeleted(c *gin.Context) {

	returns, err := h.service.GetDeleted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    returns,
	})
}

// Restore Return
func (h *Handler) Restore(c *gin.Context) {

	if err := h.service.Restore(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Return restored successfully",
	})
}

// Delete Return
func (h *Handler) Delete(c *gin.Context) {

	if err := h.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Return deleted successfully",
	})
}

// Permanent Delete
func (h *Handler) PermanentDelete(c *gin.Context) {

	if err := h.service.PermanentDelete(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Return permanently deleted",
	})
}
