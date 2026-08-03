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
func (r *Repository) GetAll(
	filter SaleFilter,
) ([]Sale, int64, error) {

	var sales []Sale
	var total int64

	query := r.db.
		Model(&Sale{}).
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier")

	// ------------------------
	// Search
	// ------------------------

	if filter.Search != "" {
		query = query.Where(
			"invoice_number ILIKE ? OR customer_name ILIKE ?",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	// ------------------------
	// Payment Method
	// ------------------------

	if filter.PaymentMethod != "" {
		query = query.Where(
			"payment_method = ?",
			filter.PaymentMethod,
		)
	}

	// ------------------------
	// Status
	// ------------------------

	if filter.Status != "" {
		query = query.Where(
			"status = ?",
			filter.Status,
		)
	}

	// ------------------------
	// Count
	// ------------------------

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ------------------------
	// Safe Sorting
	// ------------------------

	sortColumn := "created_at"

	switch filter.SortBy {

	case "latest":
		sortColumn = "created_at"

	case "oldest":
		sortColumn = "created_at"

	case "invoice":
		sortColumn = "invoice_number"

	case "customer":
		sortColumn = "customer_name"

	case "amount":
		sortColumn = "total_amount"

	case "payment":
		sortColumn = "payment_method"

	case "status":
		sortColumn = "status"
	}

	order := "DESC"

	if filter.SortBy == "oldest" {
		order = "ASC"
	}

	if filter.Order == "asc" {
		order = "ASC"
	}

	if filter.Order == "desc" {
		order = "DESC"
	}

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order(sortColumn + " " + order).
		Offset(offset).
		Limit(filter.Limit).
		Find(&sales).Error

	if err != nil {
		return nil, 0, err
	}

	return sales, total, nil
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

// Restore
func (r *Repository) Restore(id string) error {
	return r.db.
		Unscoped().
		Model(&Sale{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// Permanent Delete
func (r *Repository) PermanentDelete(id string) error {

	tx := r.db.Begin()

	if err := tx.
		Unscoped().
		Where("sale_id = ?", id).
		Delete(&SaleItem{}).Error; err != nil {

		tx.Rollback()
		return err
	}

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

	if err != nil {
		return nil, err
	}

	return sales, nil
}
