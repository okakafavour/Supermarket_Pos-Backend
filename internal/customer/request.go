package customer

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`

	Email string `json:"email"`
	Phone string `json:"phone" binding:"required"`

	Address string `json:"address"`
}

type UpdateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email string `json:"email"`
	Phone string `json:"phone"`

	Address string `json:"address"`

	IsActive *bool `json:"is_active"`
}
