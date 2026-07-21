package product

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *Repository) GetAll() ([]Product, error) {
	var products []Product

	err := r.db.
		Preload("Category").
		Preload("Supplier").
		Find(&products).Error

	return products, err
}

func (r *Repository) GetByID(id string) (*Product, error) {
	var product Product

	err := r.db.
		Preload("Category").
		Preload("Supplier").
		First(&product, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}

func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Product{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

func (r *Repository) PermanentDelete(id string) error {
	return r.db.
		Unscoped().
		Delete(&Product{}, "id = ?", id).Error
}

func (r *Repository) GetDeleted() ([]Product, error) {
	var products []Product

	err := r.db.
		Unscoped().
		Preload("Category").
		Preload("Supplier").
		Where("deleted_at IS NOT NULL").
		Find(&products).Error

	return products, err
}

func (r *Repository) GetBySKU(sku string) (*Product, error) {

	var product Product

	err := r.db.
		Where("sku = ?", sku).
		First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repository) GetByBarcode(barcode string) (*Product, error) {

	var product Product

	err := r.db.
		Where("barcode = ?", barcode).
		First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repository) UpdateQuantity(id string, quantity int) error {
	return r.db.
		Model(&Product{}).
		Where("id = ?", id).
		Update("quantity", quantity).Error
}
