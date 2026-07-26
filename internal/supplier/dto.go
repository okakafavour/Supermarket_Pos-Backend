package supplier

type CreateSupplierRequest struct {
	Name          string `json:"name" validate:"required"`
	ContactPerson string `json:"contact_person"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
}

type UpdateSupplierRequest struct {
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	IsActive      bool   `json:"is_active"`
}

// ==========================================
// Supplier Filters
// ==========================================

type SupplierFilter struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
	Status string `form:"status"`
	Sort   string `form:"sort"`
}

// ==========================================
// Paginated Response
// ==========================================

type PaginatedSuppliers struct {
	Data       []Supplier `json:"data"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Total      int64      `json:"total"`
	TotalPages int        `json:"total_pages"`
}
