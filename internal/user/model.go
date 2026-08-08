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

	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`

	Email string `gorm:"uniqueIndex;not null" json:"email"`

	Phone string `json:"phone"`

	Password string `gorm:"not null" json:"-"`

	Role Role `gorm:"type:varchar(20);default:'cashier'" json:"role"`

	IsActive bool `gorm:"default:true" json:"is_active"`
}
