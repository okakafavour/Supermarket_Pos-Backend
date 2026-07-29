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

/*
========================================
Revenue Chart
========================================
*/

type RevenuePoint struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

/*
========================================
Sales Chart
========================================
*/

type SalesPoint struct {
	Day   string `json:"day"`
	Sales int    `json:"sales"`
}

/*
========================================
Recent Sales
========================================
*/

type RecentSale struct {
	InvoiceNumber string  `json:"invoice_number"`
	Customer      string  `json:"customer"`
	Cashier       string  `json:"cashier"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
}

/*
========================================
Top Products
========================================
*/

type TopProduct struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Revenue     float64 `json:"revenue"`
}

/*
========================================
Low Stock Products
========================================
*/

type LowStockProduct struct {
	ProductID    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	Category     string `json:"category"`
	Supplier     string `json:"supplier"`
	Quantity     int    `json:"quantity"`
	MinimumStock int    `json:"minimum_stock"`
}
