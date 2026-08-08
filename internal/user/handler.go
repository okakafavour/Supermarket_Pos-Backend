package user

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

func (h *Handler) GetUsers(c *gin.Context) {

	users, err := h.service.GetUsers()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func (h *Handler) GetUser(c *gin.Context) {

	user, err := h.service.GetUserByID(c.Param("id"))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {

	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	if err := h.service.CreateUser(req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {

	var req UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	if err := h.service.UpdateUser(c.Param("id"), req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
	})
}

func (h *Handler) UpdateStatus(c *gin.Context) {

	currentUserID := c.GetString("user_id")
	targetUserID := c.Param("id")

	if currentUserID == targetUserID {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "You cannot disable your own account",
		})
		return
	}

	var req UpdateUserStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := h.service.UpdateStatus(targetUserID, req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User status updated successfully",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {

	currentUser := c.GetString("user_id")

	if currentUser == c.Param("id") {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "You cannot delete your own account",
		})

		return
	}

	if err := h.service.DeleteUser(c.Param("id")); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})

		return
	}

	var req UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	user, err := h.service.UpdateProfile(
		userID,
		req,
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
		"message": "Profile updated successfully",
		"data":    user,
	})
}
