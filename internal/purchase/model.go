package purchase

import (
	"time"

	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
)

type PurchaseStatus string

const (
	Pending   PurchaseStatus = "pending"
	Received  PurchaseStatus = "received"
	Cancelled PurchaseStatus = "cancelled"
)

type Purchase struct {
	common.BaseModel

	InvoiceNumber string `gorm:"size:100;uniqueIndex"`

	SupplierID uuid.UUID

	Supplier supplier.Supplier `gorm:"foreignKey:SupplierID"`

	Status PurchaseStatus `gorm:"type:varchar(20);default:'pending'"`

	TotalAmount float64

	ReceivedAt *time.Time

	CreatedBy string

	Items []PurchaseItem `gorm:"constraint:OnDelete:CASCADE"`
}

type PurchaseItem struct {
	common.BaseModel

	PurchaseID uuid.UUID

	ProductID uuid.UUID

	Product product.Product `gorm:"foreignKey:ProductID"`

	Quantity int

	UnitCost float64

	Subtotal float64
}
