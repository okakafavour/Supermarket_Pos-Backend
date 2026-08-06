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

	InvoiceNumber string `gorm:"size:100;uniqueIndex" json:"invoice_number"`

	SupplierID uuid.UUID `json:"supplier_id"`

	Supplier supplier.Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`

	Status PurchaseStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`

	TotalAmount float64 `json:"total_amount"`

	ReceivedAt *time.Time `json:"received_at"`

	CreatedBy string `json:"created_by"`

	Items []PurchaseItem `gorm:"constraint:OnDelete:CASCADE" json:"items"`
}

type PurchaseItem struct {
	common.BaseModel

	PurchaseID uuid.UUID `json:"purchase_id"`

	ProductID uuid.UUID `json:"product_id"`

	Product product.Product `gorm:"foreignKey:ProductID" json:"product"`

	Quantity int `json:"quantity"`

	UnitCost float64 `json:"unit_cost"`

	Subtotal float64 `json:"subtotal"`
}
