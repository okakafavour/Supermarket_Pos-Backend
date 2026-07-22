package payment

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create
func (r *Repository) Create(payment *Payment) error {
	return r.db.Create(payment).Error
}

// Get All
func (r *Repository) GetAll() ([]Payment, error) {
	var payments []Payment

	err := r.db.
		Preload("Sale").
		Preload("Sale.Items").
		Preload("Sale.Items.Product").
		Preload("Sale.Items.Product.Category").
		Preload("Sale.Items.Product.Supplier").
		Find(&payments).Error

	return payments, err
}

// Get By ID
func (r *Repository) GetByID(id string) (*Payment, error) {
	var payment Payment

	err := r.db.
		Preload("Sale").
		Preload("Sale.Items").
		Preload("Sale.Items.Product").
		Preload("Sale.Items.Product.Category").
		Preload("Sale.Items.Product.Supplier").
		First(&payment, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// Update
func (r *Repository) Update(payment *Payment) error {
	return r.db.Save(payment).Error
}

// Soft Delete
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Payment{}, "id = ?", id).Error
}

// Restore
func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Payment{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// Permanent Delete
func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Payment{}, "id = ?", id).Error
}

// Get Deleted
func (r *Repository) GetDeleted() ([]Payment, error) {
	var payments []Payment

	err := r.db.
		Unscoped().
		Preload("Sale").
		Preload("Sale.Items").
		Preload("Sale.Items.Product").
		Preload("Sale.Items.Product.Category").
		Preload("Sale.Items.Product.Supplier").
		Where("deleted_at IS NOT NULL").
		Find(&payments).Error

	return payments, err
}

// Get By Sale ID
func (r *Repository) GetBySaleID(saleID string) (*Payment, error) {

	var payment Payment

	err := r.db.
		Where("sale_id = ?", saleID).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
