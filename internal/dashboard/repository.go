package dashboard

import (
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

/*
========================================
SUMMARY
========================================
*/

func (r *Repository) CountProducts() (int64, error) {

	var count int64

	err := r.db.
		Model(&product.Product{}).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) CountCategories() (int64, error) {

	var count int64

	err := r.db.
		Model(&category.Category{}).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) CountSuppliers() (int64, error) {

	var count int64

	err := r.db.
		Model(&supplier.Supplier{}).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) CountCustomers() (int64, error) {

	var count int64

	err := r.db.
		Model(&customer.Customer{}).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) CountSales() (int64, error) {

	var count int64

	err := r.db.
		Model(&sale.Sale{}).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) CountTodaySales() (int64, error) {

	var count int64

	start := time.Now().
		Truncate(24 * time.Hour)

	end := start.Add(24 * time.Hour)

	err := r.db.
		Model(&sale.Sale{}).
		Where(
			"created_at >= ? AND created_at < ?",
			start,
			end,
		).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) TotalRevenue() (float64, error) {

	var total float64

	err := r.db.
		Model(&payment.Payment{}).
		Select(
			"COALESCE(SUM(amount),0)",
		).
		Scan(&total).
		Error

	return total, err
}

func (r *Repository) TodayRevenue() (float64, error) {

	var total float64

	start := time.Now().
		Truncate(24 * time.Hour)

	end := start.Add(24 * time.Hour)

	err := r.db.
		Model(&payment.Payment{}).
		Select(
			"COALESCE(SUM(amount),0)",
		).
		Where(
			"created_at >= ? AND created_at < ?",
			start,
			end,
		).
		Scan(&total).
		Error

	return total, err
}

func (r *Repository) CountLowStockProducts() (int64, error) {

	var count int64

	err := r.db.
		Model(&product.Product{}).
		Where(
			"quantity <= minimum_stock",
		).
		Count(&count).
		Error

	return count, err
}

/*
========================================
REVENUE ANALYTICS
LAST 12 MONTHS
========================================
*/

func (r *Repository) GetRevenueChart() (
	[]RevenuePoint, error,
) {

	var result []RevenuePoint

	err := r.db.
		Table("payments").
		Select(
			"TO_CHAR(created_at,'Mon') as month, SUM(amount) as revenue",
		).
		Group(
			"month",
		).
		Order(
			"MIN(created_at)",
		).
		Scan(&result).
		Error

	return result, err
}

/*
========================================
SALES ANALYTICS
LAST 7 DAYS
========================================
*/

func (r *Repository) GetSalesChart() (
	[]SalesPoint, error,
) {

	var result []SalesPoint

	err := r.db.
		Table("sales").
		Select(
			"TO_CHAR(created_at,'Dy') as day, COUNT(id) as sales",
		).
		Where(
			"created_at >= ?",
			time.Now().AddDate(
				0,
				0,
				-7,
			),
		).
		Group(
			"day",
		).
		Order(
			"MIN(created_at)",
		).
		Scan(&result).
		Error

	return result, err
}

/*
========================================
RECENT SALES
========================================
*/

func (r *Repository) GetRecentSales() (
	[]RecentSale, error,
) {

	var sales []RecentSale

	err := r.db.
		Table("sales").
		Select(
			`
			sales.invoice_number,
			sales.customer_name as customer,
			sales.sold_by as cashier,
			sales.total_amount as amount,
			sales.payment_method,
			sales.status,
			sales.created_at
			`,
		).
		Order(
			"sales.created_at DESC",
		).
		Limit(10).
		Scan(&sales).
		Error

	return sales, err
}

/*
========================================
TOP SELLING PRODUCTS
========================================
*/

func (r *Repository) GetTopProducts() (
	[]TopProduct, error,
) {

	var products []TopProduct

	err := r.db.
		Table("sale_items").
		Select(
			`
			products.id as product_id,
			products.name as product_name,
			SUM(sale_items.quantity) as quantity,
			SUM(sale_items.subtotal) as revenue
			`,
		).
		Joins(
			"JOIN products ON products.id = sale_items.product_id",
		).
		Group(
			"products.id, products.name",
		).
		Order(
			"quantity DESC",
		).
		Limit(5).
		Scan(&products).
		Error

	return products, err
}

/*
========================================
LOW STOCK PRODUCTS LIST
========================================
*/

func (r *Repository) GetLowStockProducts() (
	[]LowStockProduct, error,
) {

	var products []LowStockProduct

	err := r.db.
		Table("products").
		Select(
			`
			products.id as product_id,
			products.name as product_name,
			products.quantity,
			products.minimum_stock,
			categories.name as category,
			suppliers.name as supplier
			`,
		).
		Joins(
			"LEFT JOIN categories ON categories.id = products.category_id",
		).
		Joins(
			"LEFT JOIN suppliers ON suppliers.id = products.supplier_id",
		).
		Where(
			"products.quantity <= products.minimum_stock",
		).
		Order(
			"products.quantity ASC",
		).
		Scan(&products).
		Error

	return products, err
}
