package sale

import (
	"github.com/google/uuid"

	"github.com/okakafavour/supermarket-pos-backend/internal/common"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
)

type SaleItem struct {
	common.BaseModel

	SaleID uuid.UUID

	ProductID uuid.UUID

	Product product.Product

	Quantity int

	UnitPrice float64

	Subtotal float64
}
