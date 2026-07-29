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

func (r *Repository) GetAll(filter ProductFilter) ([]Product, int64, error) {

	var products []Product
	var total int64

	query := r.db.Model(&Product{}).
		Preload("Category").
		Preload("Supplier")

	// Search
	if filter.Search != "" {
		query = query.Where(
			"name ILIKE ? OR sku ILIKE ?",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	// Category
	if filter.Category != "" {
		query = query.Where("category_id = ?", filter.Category)
	}

	// Stock Status
	switch filter.Status {

	case "low":
		query = query.Where("quantity <= minimum_stock AND quantity > 0")

	case "out":
		query = query.Where("quantity = 0")

	case "healthy":
		query = query.Where("quantity > minimum_stock")
	}

	// Count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	sortBy := "created_at"

	if filter.SortBy != "" {
		sortBy = filter.SortBy
	}

	order := "DESC"

	if filter.Order == "asc" {
		order = "ASC"
	}

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order(sortBy + " " + order).
		Offset(offset).
		Limit(filter.Limit).
		Find(&products).Error

	return products, total, err
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
	return r.db.
		Model(&Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":          product.Name,
			"description":   product.Description,
			"sku":           product.SKU,
			"barcode":       product.Barcode,
			"category_id":   product.CategoryID,
			"supplier_id":   product.SupplierID,
			"cost_price":    product.CostPrice,
			"selling_price": product.SellingPrice,
			"quantity":      product.Quantity,
			"minimum_stock": product.MinimumStock,
			"image_url":     product.ImageURL,
			"is_active":     product.IsActive,
		}).Error
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
