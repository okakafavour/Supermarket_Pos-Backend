package sale

import (
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
)

type PaymentMethod string

const (
	Cash        PaymentMethod = "cash"
	Card        PaymentMethod = "card"
	Transfer    PaymentMethod = "transfer"
	MobileMoney PaymentMethod = "mobile_money"
)

type Sale struct {
	common.BaseModel

	InvoiceNumber string `gorm:"size:50;uniqueIndex;not null"`

	CustomerName string `gorm:"size:150"`

	TotalAmount float64 `gorm:"not null"`
	Discount    float64 `gorm:"default:0"`
	Tax         float64 `gorm:"default:0"`

	PaymentMethod PaymentMethod `gorm:"type:varchar(30);not null"`

	SoldBy string `gorm:"size:100;not null"`

	Items []SaleItem `gorm:"foreignKey:SaleID"`
}
