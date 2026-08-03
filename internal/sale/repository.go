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
		Preload("Items")

	//---------------------------------
	// Search
	//---------------------------------

	if filter.Search != "" {
		query = query.Where(
			"invoice_number ILIKE ? OR customer_name ILIKE ?",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	//---------------------------------
	// Payment
	//---------------------------------

	if filter.PaymentMethod != "" {
		query = query.Where(
			"payment_method = ?",
			filter.PaymentMethod,
		)
	}

	//---------------------------------
	// Status
	//---------------------------------

	if filter.Status != "" {
		query = query.Where(
			"status = ?",
			filter.Status,
		)
	}

	//---------------------------------
	// Count
	//---------------------------------

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//---------------------------------
	// Sorting
	//---------------------------------

	sortColumn := "created_at"

	switch filter.SortBy {

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

func (r *Repository) GetDashboard() (*DashboardResponse, error) {

	var dashboard DashboardResponse

	// Total Sales
	if err := r.db.
		Model(&Sale{}).
		Count(&dashboard.TotalSales).Error; err != nil {
		return nil, err
	}

	// Total Revenue
	r.db.
		Model(&Sale{}).
		Select("COALESCE(SUM(total_amount),0)").
		Scan(&dashboard.TotalRevenue)

	// Average Sale
	r.db.
		Model(&Sale{}).
		Select("COALESCE(AVG(total_amount),0)").
		Scan(&dashboard.AverageSale)

	// Pending
	r.db.
		Model(&Sale{}).
		Where("status = ?", SalePending).
		Count(&dashboard.PendingSales)

	// Paid
	r.db.
		Model(&Sale{}).
		Where("status = ?", SalePaid).
		Count(&dashboard.PaidSales)

	// Today's Revenue
	r.db.
		Model(&Sale{}).
		Where("DATE(created_at) = CURRENT_DATE").
		Select("COALESCE(SUM(total_amount),0)").
		Scan(&dashboard.TodaysRevenue)

	// Today's Sales
	r.db.
		Model(&Sale{}).
		Where("DATE(created_at) = CURRENT_DATE").
		Count(&dashboard.TodaysSales)

	return &dashboard, nil
}

func (r *Repository) GetAnalytics() (*AnalyticsResponse, error) {

	response := &AnalyticsResponse{}

	//----------------------------------
	// Sales Trend (Last 30 Days)
	//----------------------------------

	if err := r.db.
		Raw(`
			SELECT
				DATE(created_at) as date,
				COUNT(*) as sales,
				COALESCE(SUM(total_amount),0) as revenue
			FROM sales
			WHERE deleted_at IS NULL
			AND created_at >= NOW() - INTERVAL '30 days'
			GROUP BY DATE(created_at)
			ORDER BY DATE(created_at)
		`).
		Scan(&response.SalesTrend).Error; err != nil {
		return nil, err
	}

	//----------------------------------
	// Payment Methods
	//----------------------------------

	if err := r.db.
		Raw(`
			SELECT
				payment_method as method,
				COUNT(*) as count,
				COALESCE(SUM(total_amount),0) as amount
			FROM sales
			WHERE deleted_at IS NULL
			GROUP BY payment_method
			ORDER BY amount DESC
		`).
		Scan(&response.PaymentMethods).Error; err != nil {
		return nil, err
	}

	//----------------------------------
	// Top Products
	//----------------------------------

	if err := r.db.
		Raw(`
			SELECT
				p.id as product_id,
				p.name,
				SUM(si.quantity) as quantity,
				SUM(si.subtotal) as revenue
			FROM sale_items si
			JOIN products p
				ON p.id = si.product_id
			GROUP BY p.id, p.name
			ORDER BY quantity DESC
			LIMIT 10
		`).
		Scan(&response.TopProducts).Error; err != nil {
		return nil, err
	}

	//----------------------------------
	// Hourly Sales (Today)
	//----------------------------------

	if err := r.db.
		Raw(`
			SELECT
				EXTRACT(HOUR FROM created_at)::INT as hour,
				COUNT(*) as sales,
				COALESCE(SUM(total_amount),0) as revenue
			FROM sales
			WHERE deleted_at IS NULL
			AND DATE(created_at) = CURRENT_DATE
			GROUP BY hour
			ORDER BY hour
		`).
		Scan(&response.HourlySales).Error; err != nil {
		return nil, err
	}

	//----------------------------------
	// Monthly Revenue (Last 12 Months)
	//----------------------------------

	if err := r.db.
		Raw(`
			SELECT
				TO_CHAR(created_at, 'Mon') as month,
				COALESCE(SUM(total_amount),0) as revenue
			FROM sales
			WHERE deleted_at IS NULL
			AND created_at >= NOW() - INTERVAL '12 months'
			GROUP BY
				DATE_PART('year', created_at),
				DATE_PART('month', created_at),
				TO_CHAR(created_at, 'Mon')
			ORDER BY
				DATE_PART('year', created_at),
				DATE_PART('month', created_at)
		`).
		Scan(&response.MonthlyRevenue).Error; err != nil {
		return nil, err
	}

	// Ensure empty arrays instead of null

	if response.SalesTrend == nil {
		response.SalesTrend = []SalesTrend{}
	}

	if response.PaymentMethods == nil {
		response.PaymentMethods = []PaymentSummary{}
	}

	if response.TopProducts == nil {
		response.TopProducts = []TopProduct{}
	}

	if response.HourlySales == nil {
		response.HourlySales = []HourlySale{}
	}

	if response.MonthlyRevenue == nil {
		response.MonthlyRevenue = []MonthlyRevenue{}
	}

	return response, nil
}
