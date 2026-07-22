package customer

import (
	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
)

type Customer struct {
	common.BaseModel

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`

	Email string `gorm:"size:150;uniqueIndex"`
	Phone string `gorm:"size:20;uniqueIndex;not null"`

	Address string `gorm:"type:text"`

	LoyaltyPoints int64   `gorm:"default:0"`
	TotalSpent    float64 `gorm:"default:0"`
	TotalOrders   int64   `gorm:"default:0"`

	IsActive bool `gorm:"default:true"`
}

type CustomerResponse struct {
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	LoyaltyPoints int64     `json:"loyalty_points"`
	TotalSpent    float64   `json:"total_spent"`
	TotalOrders   int64     `json:"total_orders"`
	IsActive      bool      `json:"is_active"`
}

// Add this at the bottom
func ToCustomerResponse(c *Customer) CustomerResponse {
	return CustomerResponse{
		ID:            c.ID,
		FirstName:     c.FirstName,
		LastName:      c.LastName,
		FullName:      c.FirstName + " " + c.LastName,
		Email:         c.Email,
		Phone:         c.Phone,
		Address:       c.Address,
		LoyaltyPoints: c.LoyaltyPoints,
		TotalSpent:    c.TotalSpent,
		TotalOrders:   c.TotalOrders,
		IsActive:      c.IsActive,
	}
}
