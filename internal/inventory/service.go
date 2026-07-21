package inventory

import (
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Restock a product
func (s *Service) Restock(req RestockRequest, createdBy string) (*InventoryLog, error) {

	if req.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	product, err := s.repo.GetProductByID(req.ProductID.String())
	if err != nil {
		return nil, errors.New("product not found")
	}

	previousStock := product.Quantity
	newStock := previousStock + req.Quantity

	product.Quantity = newStock

	if err := s.repo.UpdateStock(product.ID.String(), newStock); err != nil {
		return nil, err
	}

	log := &InventoryLog{
		ProductID:     product.ID,
		MovementType:  Restock,
		Quantity:      req.Quantity,
		PreviousStock: previousStock,
		NewStock:      newStock,
		Reason:        req.Reason,
		CreatedBy:     createdBy,
	}

	if err := s.repo.CreateLog(log); err != nil {
		return nil, err
	}

	return s.repo.GetLogByID(log.ID.String())
}

// Adjust stock
func (s *Service) Adjust(req AdjustmentRequest, createdBy string) (*InventoryLog, error) {

	product, err := s.repo.GetProductByID(req.ProductID.String())
	if err != nil {
		return nil, errors.New("product not found")
	}

	previousStock := product.Quantity
	newStock := previousStock + req.Quantity

	if newStock < 0 {
		return nil, errors.New("insufficient stock")
	}

	product.Quantity = newStock

	if err := s.repo.UpdateStock(product.ID.String(), newStock); err != nil {
		return nil, err
	}

	log := &InventoryLog{
		ProductID:     product.ID,
		MovementType:  Adjustment,
		Quantity:      req.Quantity,
		PreviousStock: previousStock,
		NewStock:      newStock,
		Reason:        req.Reason,
		CreatedBy:     createdBy,
	}

	if err := s.repo.CreateLog(log); err != nil {
		return nil, err
	}

	return s.repo.GetLogByID(log.ID.String())
}

// Get every inventory log
func (s *Service) GetLogs() ([]InventoryLog, error) {
	return s.repo.GetAllLogs()
}

// Get logs for a single product
func (s *Service) GetProductLogs(productID string) ([]InventoryLog, error) {
	return s.repo.GetProductLogs(productID)
}
