package product

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type ProductFilter struct {
	Page     int
	Limit    int
	Search   string
	Category string
	Status   string
	SortBy   string
	Order    string
}
