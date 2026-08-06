package supplier

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name string `json:"name"`

	ContactPerson string `json:"contact_person"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Address string `json:"address"`

	City string `json:"city"`

	State string `json:"state"`

	Country string `json:"country"`

	IsActive bool `json:"is_active"`
}
