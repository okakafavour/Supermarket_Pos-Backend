package settings

import (
	"time"

	"github.com/google/uuid"
)

type Settings struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	StoreName    string `gorm:"size:150;not null" json:"store_name"`
	StoreEmail   string `gorm:"size:150" json:"store_email"`
	StorePhone   string `gorm:"size:50" json:"store_phone"`
	StoreAddress string `gorm:"type:text" json:"store_address"`

	Currency string  `gorm:"size:10;default:'NGN'" json:"currency"`
	TaxRate  float64 `gorm:"default:0" json:"tax_rate"`

	LowStockThreshold int `gorm:"default:5" json:"low_stock_threshold"`

	Timezone string `gorm:"size:100;default:'Africa/Lagos'" json:"timezone"`

	ReceiptHeader string `gorm:"type:text" json:"receipt_header"`
	ReceiptFooter string `gorm:"type:text" json:"receipt_footer"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
