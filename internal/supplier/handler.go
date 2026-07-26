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

// ==========================================
// CREATE SUPPLIER
// ==========================================

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

// ==========================================
// GET ALL SUPPLIERS
// ==========================================

func (h *Handler) GetAll(c *gin.Context) {

	var filter SupplierFilter

	if err := c.ShouldBindQuery(&filter); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	result, err := h.service.GetAll(filter)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// ==========================================
// GET SUPPLIER BY ID
// ==========================================

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

// ==========================================
// UPDATE SUPPLIER
// ==========================================

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

// ==========================================
// DELETE SUPPLIER
// ==========================================

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
		"message": "Supplier deleted successfully",
	})
}

// ==========================================
// RESTORE SUPPLIER
// ==========================================

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
		"message": "Supplier restored successfully",
	})
}

// ==========================================
// PERMANENT DELETE
// ==========================================

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
		"message": "Supplier permanently deleted",
	})
}

// ==========================================
// GET DELETED SUPPLIERS
// ==========================================

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
