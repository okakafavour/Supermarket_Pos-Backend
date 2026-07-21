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

	err = s.repo.Create(supplier)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) GetAll() ([]Supplier, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (*Supplier, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id string, req UpdateSupplierRequest) (*Supplier, error) {

	supplier, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		supplier.Name = req.Name
	}

	if req.ContactPerson != "" {
		supplier.ContactPerson = req.ContactPerson
	}

	if req.Email != "" {
		supplier.Email = req.Email
	}

	if req.Phone != "" {
		supplier.Phone = req.Phone
	}

	if req.Address != "" {
		supplier.Address = req.Address
	}

	if req.City != "" {
		supplier.City = req.City
	}

	if req.State != "" {
		supplier.State = req.State
	}

	if req.Country != "" {
		supplier.Country = req.Country
	}

	supplier.IsActive = req.IsActive

	err = s.repo.Update(supplier)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) Restore(id string) error {
	return s.repo.Restore(id)
}

func (s *Service) PermanentDelete(id string) error {
	return s.repo.PermanentDelete(id)
}

func (s *Service) GetDeleted() ([]Supplier, error) {
	return s.repo.GetDeleted()
}
