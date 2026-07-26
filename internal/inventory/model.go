package inventory

import (
	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
)

type MovementType string

const (
	StockIn    MovementType = "stock_in"
	Restock    MovementType = "restock"
	Sale       MovementType = "sale"
	Return     MovementType = "return"
	Adjustment MovementType = "adjustment"
)

type InventoryReason string

const (
	PurchaseReason   InventoryReason = "purchase"
	SaleReason       InventoryReason = "sale"
	ReturnReason     InventoryReason = "return"
	DamagedReason    InventoryReason = "damaged"
	ExpiredReason    InventoryReason = "expired"
	ManualReason     InventoryReason = "manual"
	CorrectionReason InventoryReason = "correction"
)

type InventoryLog struct {
	common.BaseModel

	ProductID uuid.UUID       `gorm:"type:uuid;index"`
	Product   product.Product `gorm:"foreignKey:ProductID"`

	MovementType MovementType `gorm:"type:varchar(20);index"`

	Quantity int

	PreviousStock int

	NewStock int

	Reason InventoryReason `gorm:"type:varchar(30);index"`

	// Invoice Number, Purchase Number, Sale Number etc.
	Reference string `gorm:"size:100"`

	// User that performed the action
	CreatedBy string `gorm:"size:100"`

	// Optional note
	Remarks string `gorm:"type:text"`
}

type InventorySummary struct {
	TotalProducts int `json:"total_products"`
	InStock       int `json:"in_stock"`
	LowStock      int `json:"low_stock"`
	OutOfStock    int `json:"out_of_stock"`

	TotalStockValue float64 `json:"total_stock_value"`
}

type PaginatedInventoryLogs struct {
	Data []InventoryLog `json:"data"`

	Page int `json:"page"`

	Limit int `json:"limit"`

	Total int64 `json:"total"`

	TotalPages int `json:"total_pages"`
}

type StockMovementChart struct {
	Day string `json:"day"`

	StockIn int `json:"stock_in"`

	StockOut int `json:"stock_out"`
}

type ProductMovement struct {
	Product string `json:"product"`

	Quantity int `json:"quantity"`
}

type InventoryAnalytics struct {
	TotalInventoryValue float64 `json:"total_inventory_value"`

	FastMoving []ProductMovement `json:"fast_moving"`

	SlowMoving []ProductMovement `json:"slow_moving"`

	LowStock []ProductMovement `json:"low_stock"`

	WeeklyMovement []StockMovementChart `json:"weekly_movement"`
}

type InventoryLogFilter struct {
	Page int

	Limit int

	Search string

	Movement string

	Reason string

	StartDate string

	EndDate string
}
