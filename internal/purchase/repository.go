package purchase

import (
	"time"

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

// Create Purchase
func (r *Repository) Create(purchase *Purchase) error {
	return r.db.Create(purchase).Error
}

// Update Purchase
func (r *Repository) Update(purchase *Purchase) error {
	return r.db.
		Model(&Purchase{}).
		Where("id = ?", purchase.ID).
		Updates(map[string]interface{}{
			"status":      purchase.Status,
			"received_at": purchase.ReceivedAt,
			"updated_at":  time.Now(),
		}).Error
}

// Get All Purchases
func (r *Repository) GetAll() ([]Purchase, error) {

	var purchases []Purchase

	err := r.db.
		Preload("Supplier").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier").
		Find(&purchases).Error

	return purchases, err
}

// Get Purchase By ID
func (r *Repository) GetByID(id string) (*Purchase, error) {

	var purchase Purchase

	err := r.db.
		Preload("Supplier").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier").
		First(&purchase, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

// Delete Purchase
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Purchase{}, "id = ?", id).Error
}
