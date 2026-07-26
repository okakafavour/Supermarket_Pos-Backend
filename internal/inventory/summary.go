package inventory

type InventorySummary struct {
	TotalProducts   int     `json:"total_products"`
	InStock         int     `json:"in_stock"`
	LowStock        int     `json:"low_stock"`
	OutOfStock      int     `json:"out_of_stock"`
	TotalStockValue float64 `json:"total_stock_value"`
}
