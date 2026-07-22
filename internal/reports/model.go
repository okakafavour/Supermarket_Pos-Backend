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
