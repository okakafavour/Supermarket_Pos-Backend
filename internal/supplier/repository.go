package supplier

import (
	"strings"

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

func (r *Repository) Create(supplier *Supplier) error {
	return r.db.Create(supplier).Error
}

// ===========================================
// GET ALL SUPPLIERS (Paginated)
// ===========================================

func (r *Repository) GetAll(
	page int,
	limit int,
	search string,
	status string,
	sort string,
) (*PaginatedSuppliers, error) {

	var suppliers []Supplier
	var total int64

	query := r.db.Model(&Supplier{})

	// -----------------------------
	// Search
	// -----------------------------

	if search != "" {

		search = "%" + strings.ToLower(search) + "%"

		query = query.Where(`
			LOWER(name) LIKE ?
			OR LOWER(email) LIKE ?
			OR LOWER(phone) LIKE ?
			OR LOWER(contact_person) LIKE ?
		`,
			search,
			search,
			search,
			search,
		)
	}

	// -----------------------------
	// Status
	// -----------------------------

	switch strings.ToLower(status) {

	case "active":
		query = query.Where("is_active = ?", true)

	case "inactive":
		query = query.Where("is_active = ?", false)
	}

	// -----------------------------
	// Count
	// -----------------------------

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// -----------------------------
	// Sorting
	// -----------------------------

	switch sort {

	case "name":
		query = query.Order("name ASC")

	case "created_at":
		query = query.Order("created_at DESC")

	default:
		query = query.Order("created_at DESC")
	}

	offset := (page - 1) * limit

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&suppliers).Error; err != nil {

		return nil, err
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &PaginatedSuppliers{
		Data: suppliers,

		Page: page,

		Limit: limit,

		Total: total,

		TotalPages: totalPages,
	}, nil
}

// ===========================================
// GET BY ID
// ===========================================

func (r *Repository) GetByID(id string) (*Supplier, error) {

	var supplier Supplier

	err := r.db.
		First(&supplier, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

// ===========================================
// GET BY EMAIL
// ===========================================

func (r *Repository) GetByEmail(email string) (*Supplier, error) {

	var supplier Supplier

	err := r.db.
		Where("email = ?", email).
		First(&supplier).Error

	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

// ===========================================
// UPDATE
// ===========================================

func (r *Repository) Update(supplier *Supplier) error {
	return r.db.Save(supplier).Error
}

// ===========================================
// SOFT DELETE
// ===========================================

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Supplier{}, "id = ?", id).Error
}

// ===========================================
// RESTORE
// ===========================================

func (r *Repository) Restore(id string) error {

	return r.db.
		Unscoped().
		Model(&Supplier{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}

// ===========================================
// PERMANENT DELETE
// ===========================================

func (r *Repository) PermanentDelete(id string) error {

	return r.db.
		Unscoped().
		Delete(&Supplier{}, "id = ?", id).Error
}

// ===========================================
// GET DELETED
// ===========================================

func (r *Repository) GetDeleted() ([]Supplier, error) {

	var suppliers []Supplier

	err := r.db.
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&suppliers).Error

	return suppliers, err
}
