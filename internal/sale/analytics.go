package sale

type SalesTrend struct {
	Date    string  `json:"date"`
	Sales   int64   `json:"sales"`
	Revenue float64 `json:"revenue"`
}

type PaymentSummary struct {
	Method string  `json:"method"`
	Count  int64   `json:"count"`
	Amount float64 `json:"amount"`
}

type TopProduct struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	Revenue   float64 `json:"revenue"`
}

type HourlySale struct {
	Hour    int     `json:"hour"`
	Sales   int64   `json:"sales"`
	Revenue float64 `json:"revenue"`
}

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

type AnalyticsResponse struct {
	SalesTrend     []SalesTrend     `json:"sales_trend"`
	PaymentMethods []PaymentSummary `json:"payment_methods"`
	TopProducts    []TopProduct     `json:"top_products"`
	HourlySales    []HourlySale     `json:"hourly_sales"`
	MonthlyRevenue []MonthlyRevenue `json:"monthly_revenue"`
}
