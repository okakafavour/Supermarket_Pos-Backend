package customer

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email"`
	Phone     string `json:"phone" validate:"required"`
	Address   string `json:"address"`
}

type UpdateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	IsActive  bool   `json:"is_active"`
}
