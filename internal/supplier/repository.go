package supplier

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(supplier *Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *Repository) GetAll() ([]Supplier, error) {
	var suppliers []Supplier

	err := r.db.Find(&suppliers).Error

	return suppliers, err
}

func (r *Repository) GetByID(id string) (*Supplier, error) {
	var supplier Supplier

	err := r.db.First(&supplier, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *Repository) GetByEmail(email string) (*Supplier, error) {
	var supplier Supplier

	err := r.db.Where("email = ?", email).First(&supplier).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *Repository) Update(supplier *Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Supplier{}, "id = ?", id).Error
}

func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Supplier{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Supplier{}, "id = ?", id).Error
}

func (r *Repository) GetDeleted() ([]Supplier, error) {
	var suppliers []Supplier

	err := r.db.
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Find(&suppliers).Error

	return suppliers, err
}
