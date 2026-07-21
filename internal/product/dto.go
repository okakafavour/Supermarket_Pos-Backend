package product

import "github.com/google/uuid"

type CreateProductRequest struct {
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description"`
	SKU          string    `json:"sku" binding:"required"`
	Barcode      string    `json:"barcode"`
	CategoryID   uuid.UUID `json:"category_id" binding:"required"`
	SupplierID   uuid.UUID `json:"supplier_id" binding:"required"`
	CostPrice    float64   `json:"cost_price" binding:"required"`
	SellingPrice float64   `json:"selling_price" binding:"required"`
	Quantity     int       `json:"quantity"`
	MinimumStock int       `json:"minimum_stock"`
	ImageURL     string    `json:"image_url"`
}

type UpdateProductRequest struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SKU          string    `json:"sku"`
	Barcode      string    `json:"barcode"`
	CategoryID   uuid.UUID `json:"category_id"`
	SupplierID   uuid.UUID `json:"supplier_id"`
	CostPrice    float64   `json:"cost_price"`
	SellingPrice float64   `json:"selling_price"`
	Quantity     int       `json:"quantity"`
	MinimumStock int       `json:"minimum_stock"`
	ImageURL     string    `json:"image_url"`
	IsActive     *bool     `json:"is_active"`
}
