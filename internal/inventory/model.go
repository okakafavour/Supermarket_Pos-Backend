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

type InventoryLog struct {
	common.BaseModel

	ProductID uuid.UUID
	Product   product.Product `gorm:"foreignKey:ProductID"`

	MovementType MovementType `gorm:"type:varchar(20)"`

	Quantity int

	PreviousStock int

	NewStock int

	Reason string

	CreatedBy string
}
