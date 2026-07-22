package returns

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create Return
func (r *Repository) Create(returnData *Return) error {
	return r.db.Create(returnData).Error
}

// Update Return
func (r *Repository) Update(returnData *Return) error {
	return r.db.Save(returnData).Error
}

// Delete Return (Soft Delete)
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Return{}, "id = ?", id).Error
}

// Restore Return
func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Return{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// Permanent Delete
func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Return{}, "id = ?", id).Error
}

// Get Return By ID
func (r *Repository) GetByID(id string) (*Return, error) {

	var returnData Return

	err := r.db.
		Preload("Sale").
		Preload("Items").
		Preload("Items.Product").
		First(&returnData, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &returnData, nil
}

// Get All Returns
func (r *Repository) GetAll() ([]Return, error) {

	var returns []Return

	err := r.db.
		Preload("Sale").
		Preload("Items").
		Preload("Items.Product").
		Order("created_at DESC").
		Find(&returns).Error

	return returns, err
}

// Get Deleted Returns
func (r *Repository) GetDeleted() ([]Return, error) {

	var returns []Return

	err := r.db.
		Unscoped().
		Preload("Sale").
		Preload("Items").
		Preload("Items.Product").
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&returns).Error

	return returns, err
}
