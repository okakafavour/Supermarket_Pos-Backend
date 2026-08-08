package settings

type UpdateSettingsRequest struct {
	StoreName    string `json:"store_name" binding:"required"`
	StoreEmail   string `json:"store_email" binding:"omitempty,email"`
	StorePhone   string `json:"store_phone"`
	StoreAddress string `json:"store_address"`

	Currency string  `json:"currency" binding:"required"`
	TaxRate  float64 `json:"tax_rate" binding:"gte=0,lte=100"`

	ReceiptFooter string `json:"receipt_footer"`

	LowStockThreshold int `json:"low_stock_threshold" binding:"gte=0"`

	LogoURL string `json:"logo_url"`
}
