package reports

import (
	"time"

	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
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

func (r *Repository) GetSalesSummary() (*SalesSummary, error) {

	var totalSales int64
	var totalRevenue float64

	// Count sales
	if err := r.db.
		Model(&sale.Sale{}).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	// Sum payments
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

	// Count today's sales
	if err := r.db.
		Model(&sale.Sale{}).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalSales).Error; err != nil {
		return nil, err
	}

	// Sum today's payments
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
