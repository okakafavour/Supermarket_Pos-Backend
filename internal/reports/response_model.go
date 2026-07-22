package reports

type SalesReport struct {
	TotalSales   int     `json:"total_sales"`
	TotalRevenue float64 `json:"total_revenue"`
}

type RevenueReport struct {
	Revenue float64 `json:"revenue"`
}

type LowStockReport struct {
	ProductID    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	Quantity     int    `json:"quantity"`
	MinimumStock int    `json:"minimum_stock"`
}

type TopProductReport struct {
	ProductID    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	QuantitySold int     `json:"quantity_sold"`
	Revenue      float64 `json:"revenue"`
}
