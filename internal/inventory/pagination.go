package inventory

type PaginatedInventoryLogs struct {
	Data       []InventoryLog `json:"data"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	Total      int64          `json:"total"`
	TotalPages int            `json:"total_pages"`
}
