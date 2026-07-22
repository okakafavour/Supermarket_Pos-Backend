package purchase

import "github.com/google/uuid"

type CreatePurchaseRequest struct {
	SupplierID uuid.UUID               `json:"supplier_id" binding:"required"`
	Items      []CreatePurchaseItemDTO `json:"items" binding:"required"`
}

type CreatePurchaseItemDTO struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	UnitCost  float64   `json:"unit_cost" binding:"required"`
}
