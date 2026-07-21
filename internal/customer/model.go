package customer

import "github.com/okakafavour/supermarket-pos-backend/internal/common"

type Customer struct {
	common.BaseModel

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`

	Email string `gorm:"uniqueIndex"`
	Phone string `gorm:"size:20;uniqueIndex;not null"`

	Address string `gorm:"type:text"`

	LoyaltyPoints int64 `gorm:"default:0"`

	IsActive bool `gorm:"default:true"`
}
