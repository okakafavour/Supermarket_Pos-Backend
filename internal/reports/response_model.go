package reports

type SalesReport struct {
	TotalSales   int     `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type RevenueReport struct {
	Revenue float64 `json:"revenue"`
}
