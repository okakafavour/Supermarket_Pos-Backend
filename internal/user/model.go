package user

import "github.com/okakafavour/supermarket-pos-backend/internal/common"

type Role string

const (
	Admin   Role = "admin"
	Manager Role = "manager"
	Cashier Role = "cashier"
)

type User struct {
	common.BaseModel

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`

	Email string `gorm:"uniqueIndex;not null"`

	Phone string

	Password string `gorm:"not null" json:"-"`

	Role Role `gorm:"type:varchar(20);default:'cashier'"`

	IsActive bool `gorm:"default:true"`
}
