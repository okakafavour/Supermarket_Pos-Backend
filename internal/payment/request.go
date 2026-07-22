package payment

import (
	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	SaleID uuid.UUID `json:"sale_id" binding:"required"`

	Amount float64 `json:"amount" binding:"required,gt=0"`

	Method PaymentMethod `json:"method" binding:"required"`

	Reference string `json:"reference"`
}
