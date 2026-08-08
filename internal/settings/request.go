package settings

type UpdateSettingsRequest struct {
	StoreName    string `json:"store_name" binding:"required"`
	StoreEmail   string `json:"store_email"`
	StorePhone   string `json:"store_phone"`
	StoreAddress string `json:"store_address"`

	Currency string  `json:"currency" binding:"required"`
	TaxRate  float64 `json:"tax_rate"`

	LowStockThreshold int `json:"low_stock_threshold"`

	Timezone string `json:"timezone"`

	ReceiptHeader string `json:"receipt_header"`
	ReceiptFooter string `json:"receipt_footer"`
}
