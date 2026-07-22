package payment

import (
	"time"

	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
)

type PaymentMethod string

const (
	Cash        PaymentMethod = "cash"
	Card        PaymentMethod = "card"
	Transfer    PaymentMethod = "transfer"
	MobileMoney PaymentMethod = "mobile_money"
)

type PaymentStatus string

const (
	Pending  PaymentStatus = "pending"
	Paid     PaymentStatus = "paid"
	Failed   PaymentStatus = "failed"
	Refunded PaymentStatus = "refunded"
)

type Payment struct {
	common.BaseModel

	SaleID uuid.UUID `gorm:"not null"`

	Sale sale.Sale `gorm:"foreignKey:SaleID"`

	Amount float64 `gorm:"not null"`

	Method PaymentMethod `gorm:"type:varchar(30);not null"`

	Status PaymentStatus `gorm:"type:varchar(30);default:'paid'"`

	Reference string `gorm:"size:100"`

	PaidAt time.Time
}
