package purchase

import (
	"time"

	"github.com/google/uuid"
)

type CreatePurchaseRequest struct {
	SupplierID uuid.UUID               `json:"supplier_id" binding:"required"`
	Items      []CreatePurchaseItemDTO `json:"items" binding:"required"`
}

type CreatePurchaseItemDTO struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gte=1"`
	UnitCost  float64   `json:"unit_cost" binding:"required,gt=0"`
}

// ==========================================
// Filters
// ==========================================

type PurchaseFilter struct {
	Page int `form:"page"`

	Limit int `form:"limit"`

	Search string `form:"search"`

	Status string `form:"status"`

	SupplierID string `form:"supplier_id"`

	Sort string `form:"sort"`

	From *time.Time `form:"from" time_format:"2006-01-02"`

	To *time.Time `form:"to" time_format:"2006-01-02"`
}

// ==========================================
// Paginated Response
// ==========================================

type PaginatedPurchases struct {
	Data []Purchase `json:"data"`

	Page int `json:"page"`

	Limit int `json:"limit"`

	Total int64 `json:"total"`

	TotalPages int `json:"total_pages"`
}

// ==========================================
// Statistics
// ==========================================

type PurchaseStats struct {
	TotalPurchases int64 `json:"total_purchases"`

	Pending int64 `json:"pending"`

	Received int64 `json:"received"`

	Cancelled int64 `json:"cancelled"`

	TotalAmount float64 `json:"total_amount"`
}
