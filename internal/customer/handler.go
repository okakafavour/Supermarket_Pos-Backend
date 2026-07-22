package customer

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

// Create Customer
func (h *Handler) Create(c *gin.Context) {

	var req CreateCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	customer, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    ToCustomerResponse(customer),
	})
}

// Get All Customers
func (h *Handler) GetAll(c *gin.Context) {

	customers, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	responses := make([]CustomerResponse, 0, len(customers))

	for i := range customers {
		responses = append(responses, ToCustomerResponse(&customers[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responses,
	})
}

// Get Customer By ID
func (h *Handler) GetByID(c *gin.Context) {

	customer, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "customer not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ToCustomerResponse(customer),
	})
}

// Update Customer
func (h *Handler) Update(c *gin.Context) {

	var req UpdateCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	customer, err := h.service.Update(c.Param("id"), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ToCustomerResponse(customer),
	})
}

// Delete Customer
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
		"message": "Customer deleted successfully",
	})
}

// Search Customer
func (h *Handler) Search(c *gin.Context) {

	query := c.Query("q")

	customers, err := h.service.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	responses := make([]CustomerResponse, 0, len(customers))

	for i := range customers {
		responses = append(responses, ToCustomerResponse(&customers[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responses,
	})
}

// Add Loyalty Points
func (h *Handler) AddLoyaltyPoints(c *gin.Context) {

	var req struct {
		Points int64 `json:"points" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	customer, err := h.service.AddLoyaltyPoints(
		c.Param("id"),
		req.Points,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ToCustomerResponse(customer),
	})
}

func (h *Handler) GetDeleted(c *gin.Context) {

	customers, err := h.service.GetDeleted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	responses := make([]CustomerResponse, 0, len(customers))

	for i := range customers {
		responses = append(responses, ToCustomerResponse(&customers[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responses,
	})
}

func (h *Handler) Restore(c *gin.Context) {

	err := h.service.Restore(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Customer restored successfully",
	})
}

func (h *Handler) PermanentDelete(c *gin.Context) {

	err := h.service.PermanentDelete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Customer permanently deleted successfully",
	})
}
