package settings

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetSettings() (*Settings, error) {
	var settings Settings

	err := r.db.First(&settings).Error

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func (r *Repository) CreateSettings(settings *Settings) error {
	return r.db.Create(settings).Error
}

func (r *Repository) UpdateSettings(settings *Settings) error {
	return r.db.Save(settings).Error
}

func (r *Repository) GetOrCreateSettings() (*Settings, error) {
	settings, err := r.GetSettings()

	if err == nil {
		return settings, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	settings = &Settings{
		StoreName:         "StockFlow",
		StoreEmail:        "",
		StorePhone:        "",
		StoreAddress:      "",
		Currency:          "NGN",
		TaxRate:           0,
		ReceiptFooter:     "Thank you for shopping with us.",
		LowStockThreshold: 5,
	}

	if err := r.CreateSettings(settings); err != nil {
		return nil, err
	}

	return settings, nil
}
