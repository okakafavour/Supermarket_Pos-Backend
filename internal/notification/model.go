package notification

import "github.com/okakafavour/supermarket-pos-backend/internal/common"

type NotificationType string

const (
	LowStockNotification   NotificationType = "low_stock"
	OutOfStockNotification NotificationType = "out_of_stock"
	SaleNotification       NotificationType = "sale"
	RestockNotification    NotificationType = "restock"
	AdjustmentNotification NotificationType = "adjustment"
	UserNotification       NotificationType = "user"
)

type Notification struct {
	common.BaseModel

	UserID string `gorm:"size:100;index" json:"user_id"`

	Type NotificationType `gorm:"type:varchar(30);index" json:"type"`

	Title string `gorm:"size:200;not null" json:"title"`

	Message string `gorm:"type:text;not null" json:"message"`

	IsRead bool `gorm:"default:false;index" json:"is_read"`
}
