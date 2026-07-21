package inventory

import (
	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
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

// Create inventory log
func (r *Repository) CreateLog(log *InventoryLog) error {
	return r.db.Create(log).Error
}

// Get all inventory logs
func (r *Repository) GetAllLogs() ([]InventoryLog, error) {

	var logs []InventoryLog

	err := r.db.
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Supplier").
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

// Get logs for a specific product
func (r *Repository) GetProductLogs(productID string) ([]InventoryLog, error) {

	var logs []InventoryLog

	err := r.db.
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Supplier").
		Where("product_id = ?", productID).
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

// Get product by ID
func (r *Repository) GetProductByID(id string) (*product.Product, error) {

	var p product.Product

	err := r.db.
		First(&p, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Update only the stock quantity
func (r *Repository) UpdateStock(productID string, quantity int) error {

	return r.db.
		Model(&product.Product{}).
		Where("id = ?", productID).
		Update("quantity", quantity).Error
}

// Check if product exists
func (r *Repository) ProductExists(id uuid.UUID) bool {

	var count int64

	r.db.
		Model(&product.Product{}).
		Where("id = ?", id).
		Count(&count)

	return count > 0
}

func (r *Repository) GetLogByID(id string) (*InventoryLog, error) {

	var log InventoryLog

	err := r.db.
		Preload("Product.Category").
		Preload("Product.Supplier").
		First(&log, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &log, nil
}
