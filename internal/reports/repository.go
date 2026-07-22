package reports

import (
	"time"

	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
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

// ===========================
// Sales Summary
// ===========================

func (r *Repository) GetSalesSummary() (*SalesSummary, error) {

	var totalSales int64
	var totalRevenue float64

	if err := r.db.
		Model(&sale.Sale{}).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}

	return &SalesSummary{
		TotalSales:   totalSales,
		TotalRevenue: totalRevenue,
	}, nil
}

// ===========================
// Daily Sales Report
// ===========================

func (r *Repository) GetDailySalesReport() (*DailySalesReport, error) {

	var totalSales int64
	var totalRevenue float64

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

	if err := r.db.
		Model(&sale.Sale{}).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Where("created_at >= ? AND created_at < ?", start, end).
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}

	return &DailySalesReport{
		Date:         start.Format("2006-01-02"),
		TotalSales:   totalSales,
		TotalRevenue: totalRevenue,
	}, nil
}

// ===========================
// Monthly Sales Report
// ===========================

func (r *Repository) GetMonthlySalesReport() (*MonthlySalesReport, error) {

	var totalSales int64
	var totalRevenue float64

	now := time.Now().UTC()

	start := time.Date(
		now.Year(),
		now.Month(),
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	end := start.AddDate(0, 1, 0)

	if err := r.db.
		Model(&sale.Sale{}).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Where("created_at >= ? AND created_at < ?", start, end).
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}

	return &MonthlySalesReport{
		Month:        start.Month().String(),
		Year:         start.Year(),
		TotalSales:   totalSales,
		TotalRevenue: totalRevenue,
	}, nil
}

// ===========================
// Date Range Report
// ===========================

func (r *Repository) GetDateRangeReport(start, end time.Time) (*DateRangeReport, error) {

	var totalSales int64
	var totalRevenue float64

	if err := r.db.
		Model(&sale.Sale{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Model(&payment.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Where("created_at BETWEEN ? AND ?", start, end).
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}

	return &DateRangeReport{
		StartDate:    start.Format("2006-01-02"),
		EndDate:      end.Format("2006-01-02"),
		TotalSales:   totalSales,
		TotalRevenue: totalRevenue,
	}, nil
}

// ===========================
// Low Stock Report
// ===========================

func (r *Repository) GetLowStockReport() ([]LowStockReport, error) {

	var report []LowStockReport

	err := r.db.
		Model(&product.Product{}).
		Select(`
			id as product_id,
			name as product_name,
			quantity,
			minimum_stock
		`).
		Where("quantity <= minimum_stock").
		Order("quantity ASC").
		Scan(&report).Error

	return report, err
}

// ===========================
// Top Selling Products
// ===========================

func (r *Repository) GetTopSellingProducts() ([]TopProductReport, error) {

	var report []TopProductReport

	err := r.db.
		Table("sale_items").
		Select(`
			products.id AS product_id,
			products.name AS product_name,
			SUM(sale_items.quantity) AS quantity_sold,
			SUM(sale_items.quantity * sale_items.unit_price) AS revenue
		`).
		Joins("JOIN products ON products.id = sale_items.product_id").
		Group("products.id, products.name").
		Order("quantity_sold DESC").
		Scan(&report).Error

	return report, err
}
