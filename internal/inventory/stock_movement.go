package inventory

import "github.com/google/uuid"

type StockMovementRequest struct {
	ProductID uuid.UUID

	Quantity int

	MovementType MovementType

	Reason InventoryReason

	Reference string

	CreatedBy string
}
