package payment

import (
	"github.com/google/uuid"

	"github.com/okakafavour/supermarket-pos-backend/internal/common"
)

type Status string

const (
	Successful Status = "successful"
	Failed     Status = "failed"
	Pending    Status = "pending"
)

type Payment struct {
	common.BaseModel

	SaleID uuid.UUID

	Method string

	Amount int64

	Reference string

	Status Status `gorm:"type:varchar(20)"`
}
