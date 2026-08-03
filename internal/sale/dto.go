package sale

import (
	"time"

	"github.com/google/uuid"
)

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

type SaleFilter struct {
	Page  int
	Limit int

	Search        string
	Status        string
	PaymentMethod string

	From string
	To   string

	SortBy string
	Order  string
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type SaleListResponse struct {
	ID            uuid.UUID     `json:"id"`
	InvoiceNumber string        `json:"invoice_number"`
	CustomerName  string        `json:"customer_name"`
	TotalAmount   float64       `json:"total_amount"`
	Status        SaleStatus    `json:"status"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	SoldBy        string        `json:"sold_by"`
	CreatedAt     time.Time     `json:"created_at"`
}

type DashboardResponse struct {
	TotalSales    int64   `json:"total_sales"`
	TotalRevenue  float64 `json:"total_revenue"`
	AverageSale   float64 `json:"average_sale"`
	PendingSales  int64   `json:"pending_sales"`
	PaidSales     int64   `json:"paid_sales"`
	TodaysRevenue float64 `json:"todays_revenue"`
	TodaysSales   int64   `json:"todays_sales"`
}
