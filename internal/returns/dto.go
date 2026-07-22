package returns

import "github.com/google/uuid"

type ReturnItemRequest struct {
	SaleItemID uuid.UUID `json:"sale_item_id" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required,min=1"`
}

type CreateReturnRequest struct {
	SaleID uuid.UUID `json:"sale_id" binding:"required"`

	Reason string `json:"reason"`

	Items []ReturnItemRequest `json:"items" binding:"required,min=1"`
}
