package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
)

type Service struct {
	repo         *Repository
	categoryRepo *category.Repository
	supplierRepo *supplier.Repository
}

func NewService(
	repo *Repository,
	categoryRepo *category.Repository,
	supplierRepo *supplier.Repository,
) *Service {

	return &Service{
		repo:         repo,
		categoryRepo: categoryRepo,
		supplierRepo: supplierRepo,
	}
}

func (s *Service) Create(req CreateProductRequest) (*Product, error) {

	// Validate Category
	_, err := s.categoryRepo.GetByID(req.CategoryID.String())
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Validate Supplier
	_, err = s.supplierRepo.GetByID(req.SupplierID.String())
	if err != nil {
		return nil, errors.New("supplier not found")
	}

	// Validate Prices
	if req.SellingPrice < req.CostPrice {
		return nil, errors.New("selling price cannot be less than cost price")
	}

	// Validate SKU
	existingSKU, _ := s.repo.GetBySKU(req.SKU)
	if existingSKU != nil {
		return nil, errors.New("sku already exists")
	}

	// Validate Barcode
	if req.Barcode != "" {
		existingBarcode, _ := s.repo.GetByBarcode(req.Barcode)
		if existingBarcode != nil {
			return nil, errors.New("barcode already exists")
		}
	}

	product := &Product{
		Name:         req.Name,
		Description:  req.Description,
		SKU:          req.SKU,
		Barcode:      req.Barcode,
		CategoryID:   req.CategoryID,
		SupplierID:   req.SupplierID,
		CostPrice:    req.CostPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		MinimumStock: req.MinimumStock,
		ImageURL:     req.ImageURL,
		IsActive:     true,
	}

	if product.MinimumStock <= 0 {
		product.MinimumStock = 5
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	// Return product with Category and Supplier loaded
	return s.repo.GetByID(product.ID.String())
}

func (s *Service) GetAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id string, req UpdateProductRequest) (*Product, error) {

	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}

	if req.Description != "" {
		product.Description = req.Description
	}

	if req.SKU != "" && req.SKU != product.SKU {

		existingSKU, _ := s.repo.GetBySKU(req.SKU)
		if existingSKU != nil {
			return nil, errors.New("sku already exists")
		}

		product.SKU = req.SKU
	}

	if req.Barcode != "" && req.Barcode != product.Barcode {

		existingBarcode, _ := s.repo.GetByBarcode(req.Barcode)
		if existingBarcode != nil {
			return nil, errors.New("barcode already exists")
		}

		product.Barcode = req.Barcode
	}

	// Validate Category
	if req.CategoryID != uuid.Nil {

		_, err := s.categoryRepo.GetByID(req.CategoryID.String())
		if err != nil {
			return nil, errors.New("category not found")
		}

		product.CategoryID = req.CategoryID
	}

	// Validate Supplier
	if req.SupplierID != uuid.Nil {

		_, err := s.supplierRepo.GetByID(req.SupplierID.String())
		if err != nil {
			return nil, errors.New("supplier not found")
		}

		product.SupplierID = req.SupplierID
	}

	if req.CostPrice > 0 {
		product.CostPrice = req.CostPrice
	}

	if req.SellingPrice > 0 {
		product.SellingPrice = req.SellingPrice
	}

	if product.SellingPrice < product.CostPrice {
		return nil, errors.New("selling price cannot be less than cost price")
	}

	if req.Quantity >= 0 {
		product.Quantity = req.Quantity
	}

	if req.MinimumStock > 0 {
		product.MinimumStock = req.MinimumStock
	}

	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}

	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	// Return updated product with relations loaded
	return s.repo.GetByID(product.ID.String())
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

func (s *Service) GetDeleted() ([]Product, error) {
	return s.repo.GetDeleted()
}
