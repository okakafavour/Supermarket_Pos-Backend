package purchase

import (
	"strings"
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
// ===========================================
// GET ALL PURCHASES (Paginated)
// ===========================================

func (r *Repository) GetAll(
	filter PurchaseFilter,
) (*PaginatedPurchases, error) {

	var purchases []Purchase
	var total int64

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	query := r.db.Model(&Purchase{}).
		Preload("Supplier").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Preload("Items.Product.Supplier")

	//----------------------------------
	// Search
	//----------------------------------

	if filter.Search != "" {

		search := "%" + strings.ToLower(filter.Search) + "%"

		query = query.Joins("LEFT JOIN suppliers ON suppliers.id = purchases.supplier_id").
			Where(`
				LOWER(purchases.invoice_number) LIKE ?
				OR LOWER(suppliers.name) LIKE ?
			`,
				search,
				search,
			)
	}

	//----------------------------------
	// Status
	//----------------------------------

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	//----------------------------------
	// Supplier
	//----------------------------------

	if filter.SupplierID != "" {
		query = query.Where("supplier_id = ?", filter.SupplierID)
	}

	//----------------------------------
	// Date Range
	//----------------------------------

	if filter.From != nil {
		query = query.Where("created_at >= ?", *filter.From)
	}

	if filter.To != nil {
		query = query.Where("created_at <= ?", *filter.To)
	}

	//----------------------------------
	// Count
	//----------------------------------

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	//----------------------------------
	// Sorting
	//----------------------------------

	switch filter.Sort {

	case "oldest":
		query = query.Order("created_at ASC")

	case "amount":
		query = query.Order("total_amount DESC")

	case "invoice":
		query = query.Order("invoice_number ASC")

	default:
		query = query.Order("created_at DESC")
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.
		Limit(filter.Limit).
		Offset(offset).
		Find(&purchases).Error; err != nil {

		return nil, err
	}

	totalPages := int((total + int64(filter.Limit) - 1) / int64(filter.Limit))

	return &PaginatedPurchases{
		Data: purchases,

		Page: filter.Page,

		Limit: filter.Limit,

		Total: total,

		TotalPages: totalPages,
	}, nil
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

// ===========================================
// PURCHASE STATS
// ===========================================

func (r *Repository) GetStats() (*PurchaseStats, error) {

	stats := &PurchaseStats{}

	r.db.Model(&Purchase{}).
		Count(&stats.TotalPurchases)

	r.db.Model(&Purchase{}).
		Where("status = ?", Pending).
		Count(&stats.Pending)

	r.db.Model(&Purchase{}).
		Where("status = ?", Received).
		Count(&stats.Received)

	r.db.Model(&Purchase{}).
		Where("status = ?", Cancelled).
		Count(&stats.Cancelled)

	r.db.Model(&Purchase{}).
		Select("COALESCE(SUM(total_amount),0)").
		Scan(&stats.TotalAmount)

	return stats, nil
}

// ===========================================
// CANCEL PURCHASE
// ===========================================

func (r *Repository) Cancel(purchase *Purchase) error {

	return r.db.
		Model(&Purchase{}).
		Where("id = ?", purchase.ID).
		Updates(map[string]interface{}{
			"status":     Cancelled,
			"updated_at": time.Now(),
		}).Error
}

// Delete Purchase
func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Purchase{}, "id = ?", id).Error
}
