package supplier

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

func (h *Handler) Create(c *gin.Context) {

	var req CreateSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	supplier, err := h.service.Create(req)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    supplier,
	})
}

func (h *Handler) GetAll(c *gin.Context) {

	suppliers, err := h.service.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    suppliers,
	})
}

func (h *Handler) GetByID(c *gin.Context) {

	id := c.Param("id")

	supplier, err := h.service.GetByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Supplier not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    supplier,
	})
}

func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")

	var req UpdateSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	supplier, err := h.service.Update(id, req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    supplier,
	})
}

func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := h.service.Delete(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Supplier deleted successfully",
	})
}

func (h *Handler) Restore(c *gin.Context) {

	id := c.Param("id")

	err := h.service.Restore(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Supplier restored successfully",
	})
}

func (h *Handler) PermanentDelete(c *gin.Context) {

	id := c.Param("id")

	err := h.service.PermanentDelete(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Supplier permanently deleted",
	})
}

func (h *Handler) GetDeleted(c *gin.Context) {

	suppliers, err := h.service.GetDeleted()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    suppliers,
	})
}
