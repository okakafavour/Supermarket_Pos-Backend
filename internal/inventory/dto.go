package inventory

import "github.com/google/uuid"

type RestockRequest struct {
	ProductID uuid.UUID       `json:"product_id" binding:"required"`
	Quantity  int             `json:"quantity" binding:"required"`
	Reason    InventoryReason `json:"reason"`
	Reference string          `json:"reference"`
}

type AdjustmentRequest struct {
	ProductID uuid.UUID       `json:"product_id" binding:"required"`
	Quantity  int             `json:"quantity" binding:"required"`
	Reason    InventoryReason `json:"reason"`
	Reference string          `json:"reference"`
}
