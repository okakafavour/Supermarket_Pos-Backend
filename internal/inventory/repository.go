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

func (r *Repository) GetInventorySummary() (*InventorySummary, error) {

	var summary InventorySummary

	var totalProducts int64
	var inStock int64
	var lowStock int64
	var outOfStock int64

	// Total Products
	r.db.
		Model(&product.Product{}).
		Count(&totalProducts)

	// In Stock
	r.db.
		Model(&product.Product{}).
		Where("quantity > minimum_stock").
		Count(&inStock)

	// Low Stock
	r.db.
		Model(&product.Product{}).
		Where("quantity > 0 AND quantity <= minimum_stock").
		Count(&lowStock)

	// Out Of Stock
	r.db.
		Model(&product.Product{}).
		Where("quantity = 0").
		Count(&outOfStock)

	summary.TotalProducts = int(totalProducts)
	summary.InStock = int(inStock)
	summary.LowStock = int(lowStock)
	summary.OutOfStock = int(outOfStock)

	// Calculate total inventory value
	type Result struct {
		Total float64
	}

	var result Result

	if err := r.db.
		Model(&product.Product{}).
		Select("COALESCE(SUM(quantity * cost_price), 0) AS total").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	summary.TotalStockValue = result.Total

	return &summary, nil
}

func (r *Repository) GetAllLogs(
	page int,
	limit int,
	search string,
	movement string,
	reason string,
	startDate string,
	endDate string,
) (*PaginatedInventoryLogs, error) {

	var logs []InventoryLog
	var total int64

	query := r.db.
		Model(&InventoryLog{}).
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Supplier")

	// ===================================
	// Filters
	// ===================================

	if movement != "" {
		query = query.Where("movement_type = ?", movement)
	}

	if reason != "" {
		query = query.Where("reason = ?", reason)
	}

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if search != "" {

		query = query.
			Joins(
				"JOIN products ON products.id = inventory_logs.product_id",
			).
			Where(
				`products.name ILIKE ?
				OR inventory_logs.created_by ILIKE ?
				OR inventory_logs.reference ILIKE ?`,
				"%"+search+"%",
				"%"+search+"%",
				"%"+search+"%",
			)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	if err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error; err != nil {

		return nil, err
	}

	totalPages := 0

	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return &PaginatedInventoryLogs{
		Data:       logs,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (r *Repository) Transaction(fn func(tx *Repository) error) error {

	return r.db.Transaction(func(tx *gorm.DB) error {

		repo := &Repository{
			db: tx,
		}

		return fn(repo)
	})
}

func (r *Repository) GetInventoryAnalytics() (*InventoryAnalytics, error) {

	analytics := &InventoryAnalytics{
		FastMoving:     []ProductMovement{},
		SlowMoving:     []ProductMovement{},
		LowStock:       []ProductMovement{},
		WeeklyMovement: []StockMovementChart{},
	}

	// =====================================
	// Products
	// =====================================

	var products []product.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	totalValue := 0.0

	for _, p := range products {

		totalValue += float64(p.Quantity) * p.CostPrice

		if p.Quantity <= p.MinimumStock {

			analytics.LowStock = append(
				analytics.LowStock,
				ProductMovement{
					Product:  p.Name,
					Quantity: p.Quantity,
				},
			)
		}
	}

	analytics.TotalInventoryValue = totalValue

	// =====================================
	// Fast / Slow Moving Products
	// =====================================

	type Result struct {
		ProductID uuid.UUID
		Name      string
		Total     int
	}

	var movement []Result

	r.db.
		Table("inventory_logs").
		Select(`
			products.id AS product_id,
			products.name,
			SUM(ABS(inventory_logs.quantity)) AS total
		`).
		Joins("JOIN products ON products.id = inventory_logs.product_id").
		Group("products.id, products.name").
		Order("total DESC").
		Scan(&movement)

	for i, item := range movement {

		pm := ProductMovement{
			Product:  item.Name,
			Quantity: item.Total,
		}

		if i < 5 {
			analytics.FastMoving = append(
				analytics.FastMoving,
				pm,
			)
		}
	}

	for i := len(movement) - 1; i >= 0 && len(analytics.SlowMoving) < 5; i-- {

		analytics.SlowMoving = append(
			analytics.SlowMoving,
			ProductMovement{
				Product:  movement[i].Name,
				Quantity: movement[i].Total,
			},
		)
	}

	// =====================================
	// Weekly Movement
	// =====================================

	var logs []InventoryLog

	if err := r.db.
		Order("created_at ASC").
		Find(&logs).Error; err != nil {

		return nil, err
	}

	days := map[string]*StockMovementChart{
		"Mon": {Day: "Mon"},
		"Tue": {Day: "Tue"},
		"Wed": {Day: "Wed"},
		"Thu": {Day: "Thu"},
		"Fri": {Day: "Fri"},
		"Sat": {Day: "Sat"},
		"Sun": {Day: "Sun"},
	}

	for _, log := range logs {

		day := log.CreatedAt.Format("Mon")

		entry := days[day]

		switch log.MovementType {

		case Restock, StockIn, Return:

			qty := log.Quantity

			if qty < 0 {
				qty = -qty
			}

			entry.StockIn += qty

		case Sale:

			qty := log.Quantity

			if qty < 0 {
				qty = -qty
			}

			entry.StockOut += qty
		}
	}

	order := []string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}

	for _, day := range order {

		analytics.WeeklyMovement = append(
			analytics.WeeklyMovement,
			*days[day],
		)
	}

	return analytics, nil
}
