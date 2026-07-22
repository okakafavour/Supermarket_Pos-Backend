package customer

import (
	"fmt"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create Customer
func (s *Service) Create(req CreateCustomerRequest) (*Customer, error) {

	customer := &Customer{
		FirstName: strings.TrimSpace(req.FirstName),
		LastName:  strings.TrimSpace(req.LastName),
		Email:     strings.TrimSpace(req.Email),
		Phone:     strings.TrimSpace(req.Phone),
		Address:   strings.TrimSpace(req.Address),
	}

	if err := s.repo.Create(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// Update Customer
func (s *Service) Update(id string, req UpdateCustomerRequest) (*Customer, error) {

	customer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("customer not found")
	}

	if req.FirstName != "" {
		customer.FirstName = strings.TrimSpace(req.FirstName)
	}

	if req.LastName != "" {
		customer.LastName = strings.TrimSpace(req.LastName)
	}

	if req.Email != "" {
		customer.Email = strings.TrimSpace(req.Email)
	}

	if req.Phone != "" {
		customer.Phone = strings.TrimSpace(req.Phone)
	}

	if req.Address != "" {
		customer.Address = strings.TrimSpace(req.Address)
	}

	if req.IsActive != nil {
		customer.IsActive = *req.IsActive
	}

	if err := s.repo.Update(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// Delete Customer
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

// Restore Customer
func (s *Service) Restore(id string) error {
	return s.repo.Restore(id)
}

// Permanent Delete
func (s *Service) PermanentDelete(id string) error {
	return s.repo.PermanentDelete(id)
}

// Get Customer By ID
func (s *Service) GetByID(id string) (*Customer, error) {
	return s.repo.GetByID(id)
}

// Get All Customers
func (s *Service) GetAll() ([]Customer, error) {
	return s.repo.GetAll()
}

// Get Deleted Customers
func (s *Service) GetDeleted() ([]Customer, error) {
	return s.repo.GetDeleted()
}

// Search Customers
func (s *Service) Search(query string) ([]Customer, error) {

	query = strings.TrimSpace(query)

	if query == "" {
		return s.repo.GetAll()
	}

	return s.repo.Search(query)
}

// Add Loyalty Points
func (s *Service) AddLoyaltyPoints(id string, points int64) (*Customer, error) {

	customer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("customer not found")
	}

	customer.LoyaltyPoints += points

	if err := s.repo.Update(customer); err != nil {
		return nil, err
	}

	return customer, nil
}
