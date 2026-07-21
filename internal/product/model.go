package product

import (
	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
)

type Product struct {
	common.BaseModel

	Name        string `gorm:"size:150;not null"`
	Description string

	SKU     string `gorm:"uniqueIndex;not null"`
	Barcode string `gorm:"uniqueIndex"`

	CategoryID uuid.UUID
	Category   category.Category `gorm:"foreignKey:CategoryID"`

	SupplierID uuid.UUID
	Supplier   supplier.Supplier `gorm:"foreignKey:SupplierID"`

	CostPrice    float64
	SellingPrice float64

	Quantity int

	MinimumStock int `gorm:"default:5"`

	ImageURL string

	IsActive bool `gorm:"default:true"`
}
