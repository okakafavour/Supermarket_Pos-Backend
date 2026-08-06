package supplier

import (
	"errors"

	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// ==========================================
// CREATE
// ==========================================

func (s *Service) Create(req CreateSupplierRequest) (*Supplier, error) {

	existing, err := s.repo.GetByEmail(req.Email)

	if err == nil && existing != nil {
		return nil, errors.New("supplier with this email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	supplier := &Supplier{
		Name:          req.Name,
		ContactPerson: req.ContactPerson,
		Email:         req.Email,
		Phone:         req.Phone,
		Address:       req.Address,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		IsActive:      true,
	}

	if err := s.repo.Create(supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

// ==========================================
// GET ALL
// ==========================================

func (s *Service) GetAll(filter SupplierFilter) (*PaginatedSuppliers, error) {

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.Sort == "" {
		filter.Sort = "created_at"
	}

	return s.repo.GetAll(
		filter.Page,
		filter.Limit,
		filter.Search,
		filter.Status,
		filter.Sort,
	)
}

// ==========================================
// GET BY ID
// ==========================================

func (s *Service) GetByID(id string) (*Supplier, error) {
	return s.repo.GetByID(id)
}

// ==========================================
// UPDATE
// ==========================================

func (s *Service) Update(id string, req UpdateSupplierRequest) (*Supplier, error) {

	supplier, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		supplier.Name = *req.Name
	}

	if req.ContactPerson != nil {
		supplier.ContactPerson = *req.ContactPerson
	}

	if req.Email != nil {

		// prevent duplicate emails
		existing, err := s.repo.GetByEmail(*req.Email)

		if err == nil &&
			existing != nil &&
			existing.ID != supplier.ID {

			return nil, errors.New("supplier with this email already exists")
		}

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		supplier.Email = *req.Email
	}

	if req.Phone != nil {
		supplier.Phone = *req.Phone
	}

	if req.Address != nil {
		supplier.Address = *req.Address
	}

	if req.City != nil {
		supplier.City = *req.City
	}

	if req.State != nil {
		supplier.State = *req.State
	}

	if req.Country != nil {
		supplier.Country = *req.Country
	}

	if req.IsActive != nil {
		supplier.IsActive = *req.IsActive
	}

	if err := s.repo.Update(supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

// ==========================================
// DELETE
// ==========================================

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

// ==========================================
// RESTORE
// ==========================================

func (s *Service) Restore(id string) error {
	return s.repo.Restore(id)
}

// ==========================================
// PERMANENT DELETE
// ==========================================

func (s *Service) PermanentDelete(id string) error {
	return s.repo.PermanentDelete(id)
}

// ==========================================
// GET DELETED
// ==========================================

func (s *Service) GetDeleted() ([]Supplier, error) {
	return s.repo.GetDeleted()
}
