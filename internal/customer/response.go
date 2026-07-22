package customer

func ToResponse(customer *Customer) CustomerResponse {

	return CustomerResponse{
		ID:            customer.ID,
		FirstName:     customer.FirstName,
		LastName:      customer.LastName,
		FullName:      customer.FirstName + " " + customer.LastName,
		Email:         customer.Email,
		Phone:         customer.Phone,
		Address:       customer.Address,
		LoyaltyPoints: customer.LoyaltyPoints,
		TotalSpent:    customer.TotalSpent,
		TotalOrders:   customer.TotalOrders,
		IsActive:      customer.IsActive,
	}
}

func ToResponses(customers []Customer) []CustomerResponse {

	responses := make([]CustomerResponse, 0, len(customers))

	for i := range customers {
		responses = append(responses, ToResponse(&customers[i]))
	}

	return responses
}
