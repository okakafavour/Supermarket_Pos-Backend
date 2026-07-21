package supplier

import "github.com/okakafavour/supermarket-pos-backend/internal/common"

type Supplier struct {
	common.BaseModel

	Name          string `gorm:"size:150;not null"`
	ContactPerson string `gorm:"size:150"`
	Email         string `gorm:"uniqueIndex"`
	Phone         string `gorm:"size:20"`
	Address       string `gorm:"type:text"`
	City          string `gorm:"size:100"`
	State         string `gorm:"size:100"`
	Country       string `gorm:"size:100"`
	IsActive      bool   `gorm:"default:true"`
}
