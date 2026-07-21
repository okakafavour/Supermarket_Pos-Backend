package category

import "github.com/okakafavour/supermarket-pos-backend/internal/common"

type Category struct {
	common.BaseModel

	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `json:"description"`
}
