package category

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r *Repository) GetAll() ([]Category, error) {
	var categories []Category

	err := r.db.Find(&categories).Error

	return categories, err
}

func (r *Repository) GetByID(id string) (*Category, error) {
	var category Category

	err := r.db.First(&category, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) Update(category *Category) error {
	return r.db.Save(category).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Category{}, "id = ?", id).Error
}

func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Category{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Category{}, "id = ?", id).Error
}

func (r *Repository) GetDeleted() ([]Category, error) {
	var categories []Category

	err := r.db.
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Find(&categories).Error

	return categories, err
}
