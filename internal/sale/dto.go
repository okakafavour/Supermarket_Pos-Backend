package sale

import "github.com/google/uuid"

type CreateSaleRequest struct {
	CustomerName  string                  `json:"customer_name"`
	Discount      float64                 `json:"discount"`
	Tax           float64                 `json:"tax"`
	PaymentMethod PaymentMethod           `json:"payment_method" binding:"required"`
	Items         []CreateSaleItemRequest `json:"items" binding:"required,min=1"`
}

type CreateSaleItemRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
}

type UpdatePaymentRequest struct {
	PaymentMethod PaymentMethod `json:"payment_method" binding:"required"`
}

type SaleResponse struct {
	ID            uuid.UUID          `json:"id"`
	InvoiceNumber string             `json:"invoice_number"`
	CustomerName  string             `json:"customer_name"`
	TotalAmount   float64            `json:"total_amount"`
	Discount      float64            `json:"discount"`
	Tax           float64            `json:"tax"`
	PaymentMethod PaymentMethod      `json:"payment_method"`
	SoldBy        string             `json:"sold_by"`
	Items         []SaleItemResponse `json:"items"`
}

type SaleItemResponse struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unit_price"`
	TotalPrice  float64   `json:"total_price"`
}
