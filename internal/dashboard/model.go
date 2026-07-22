package dashboard

type Summary struct {
	TotalProducts   int `json:"total_products"`
	TotalCategories int `json:"total_categories"`
	TotalSuppliers  int `json:"total_suppliers"`
	TotalCustomers  int `json:"total_customers"`

	TotalSales int `json:"total_sales"`
	TodaySales int `json:"today_sales"`

	TotalRevenue float64 `json:"total_revenue"`
	TodayRevenue float64 `json:"today_revenue"`

	LowStockProducts int `json:"low_stock_products"`
}
