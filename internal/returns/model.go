package returns

import (
	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
)

const (
	ReturnPending   = "pending"
	ReturnCompleted = "completed"
)

type Return struct {
	common.BaseModel

	SaleID uuid.UUID `gorm:"type:uuid;not null"`
	Sale   sale.Sale `gorm:"foreignKey:SaleID"`

	RefundAmount float64 `gorm:"not null"`

	Reason string `gorm:"type:text"`

	Status string `gorm:"default:'completed'"`

	ProcessedBy string

	Items []ReturnItem `gorm:"foreignKey:ReturnID"`
}

type ReturnItem struct {
	common.BaseModel

	ReturnID uuid.UUID `gorm:"type:uuid;not null"`

	ProductID uuid.UUID       `gorm:"type:uuid;not null"`
	Product   product.Product `gorm:"foreignKey:ProductID"`

	SaleItemID uuid.UUID `gorm:"type:uuid;not null"`

	Quantity int `gorm:"not null"`

	UnitPrice float64 `gorm:"not null"`

	Subtotal float64 `gorm:"not null"`
}
