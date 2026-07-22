package dashboard

import (
	"fmt"
	"time"

	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/customer"
	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
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

// Count Products
func (r *Repository) CountProducts() (int64, error) {
	var count int64

	err := r.db.Model(&product.Product{}).Count(&count).Error

	return count, err
}

// Count Categories
func (r *Repository) CountCategories() (int64, error) {
	var count int64

	err := r.db.Model(&category.Category{}).Count(&count).Error

	return count, err
}

// Count Suppliers
func (r *Repository) CountSuppliers() (int64, error) {
	var count int64

	err := r.db.Model(&supplier.Supplier{}).Count(&count).Error

	return count, err
}

// Count Customers
func (r *Repository) CountCustomers() (int64, error) {
	var count int64

	err := r.db.Model(&customer.Customer{}).Count(&count).Error

	return count, err
}

// Count Sales
func (r *Repository) CountSales() (int64, error) {
	var count int64

	err := r.db.Model(&sale.Sale{}).Count(&count).Error

	return count, err
}

// Count Today's Sales (DEBUG)
func (r *Repository) CountTodaySales() (int64, error) {
	var count int64

	now := time.Now()

	start := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	end := start.Add(24 * time.Hour)

	fmt.Println("NOW   :", now)
	fmt.Println("START :", start)
	fmt.Println("END   :", end)

	err := r.db.
		Model(&sale.Sale{}).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&count).Error

	fmt.Println("COUNT :", count)

	return count, err
}

// Total Revenue
func (r *Repository) TotalRevenue() (float64, error) {
	var total float64

	err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error

	return total, err
}

// Today's Revenue
func (r *Repository) TodayRevenue() (float64, error) {
	var total float64

	now := time.Now().UTC()

	start := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	end := start.Add(24 * time.Hour)

	err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("created_at >= ? AND created_at < ?", start, end).
		Scan(&total).Error

	return total, err
}

// Count Low Stock Products
func (r *Repository) CountLowStockProducts() (int64, error) {
	var count int64

	err := r.db.
		Model(&product.Product{}).
		Where("quantity <= minimum_stock").
		Count(&count).Error

	return count, err
}
