package reports

type SalesSummary struct {
	TotalSales   int64   `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type DailySalesReport struct {
	Date         string  `json:"date"`
	TotalSales   int64   `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type MonthlySalesReport struct {
	Month        string  `json:"month"`
	Year         int     `json:"year"`
	TotalSales   int64   `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type DateRangeReport struct {
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	TotalSales   int64   `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type LowStockReport struct {
	ProductID    string `json:"product_id" gorm:"column:product_id"`
	ProductName  string `json:"product_name" gorm:"column:product_name"`
	Quantity     int    `json:"quantity"`
	MinimumStock int    `json:"minimum_stock"`
}

type TopProductReport struct {
	ProductID    string  `json:"product_id" gorm:"column:product_id"`
	ProductName  string  `json:"product_name" gorm:"column:product_name"`
	QuantitySold int64   `json:"quantity_sold" gorm:"column:quantity_sold"`
	Revenue      float64 `json:"revenue"`
}
