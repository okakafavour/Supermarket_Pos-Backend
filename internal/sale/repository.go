package sale

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create Sale
func (r *Repository) Create(sale *Sale) error {
	return r.db.Create(sale).Error
}

// Get All Sales
func (r *Repository) GetAll() ([]Sale, error) {
	var sales []Sale

	err := r.db.
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier").
		Order("created_at DESC").
		Find(&sales).Error

	return sales, err
}

// Get Sale By ID
func (r *Repository) GetByID(id string) (*Sale, error) {
	var sale Sale

	err := r.db.
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier").
		First(&sale, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &sale, nil
}

// Update Sale
func (r *Repository) Update(sale *Sale) error {
	return r.db.Save(sale).Error
}

// Soft Delete
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Sale{}, "id = ?", id).Error
}

// Restore Soft Deleted Sale
func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Sale{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// Permanently Delete Sale
func (r *Repository) PermanentDelete(id string) error {

	tx := r.db.Begin()

	// Delete sale items first
	if err := tx.
		Unscoped().
		Where("sale_id = ?", id).
		Delete(&SaleItem{}).Error; err != nil {

		tx.Rollback()
		return err
	}

	// Delete the sale
	if err := tx.
		Unscoped().
		Delete(&Sale{}, "id = ?", id).Error; err != nil {

		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Get Deleted Sales
func (r *Repository) GetDeleted() ([]Sale, error) {
	var sales []Sale

	err := r.db.
		Unscoped().
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier").
		Where("deleted_at IS NOT NULL").
		Order("created_at DESC").
		Find(&sales).Error

	return sales, err
}
