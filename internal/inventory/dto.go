package inventory

import "github.com/google/uuid"

type RestockRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Reason    string    `json:"reason"`
}

type AdjustmentRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Reason    string    `json:"reason"`
}
