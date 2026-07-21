package sale

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
)

type Service struct {
	repo          *Repository
	productRepo   *product.Repository
	inventoryRepo *inventory.Repository
}

func NewService(
	repo *Repository,
	productRepo *product.Repository,
	inventoryRepo *inventory.Repository,
) *Service {
	return &Service{
		repo:          repo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *Service) Create(req CreateSaleRequest, userID string) (*Sale, error) {

	if len(req.Items) == 0 {
		return nil, errors.New("sale must contain at least one item")
	}

	sale := &Sale{
		InvoiceNumber: fmt.Sprintf("INV-%d", time.Now().Unix()),
		CustomerName:  req.CustomerName,
		Discount:      req.Discount,
		Tax:           req.Tax,
		PaymentMethod: req.PaymentMethod,
		SoldBy:        userID,
	}

	var total float64

	for _, item := range req.Items {

		// Get product
		productData, err := s.productRepo.GetByID(item.ProductID.String())
		if err != nil {
			return nil, errors.New("product not found")
		}

		// Check stock
		if productData.Quantity < item.Quantity {
			return nil, fmt.Errorf(
				"not enough stock for %s",
				productData.Name,
			)
		}

		// Calculate subtotal
		subtotal := float64(item.Quantity) * productData.SellingPrice

		// Add sale item
		sale.Items = append(sale.Items, SaleItem{
			ProductID: productData.ID,
			Quantity:  item.Quantity,
			UnitPrice: productData.SellingPrice,
			Subtotal:  subtotal,
		})

		total += subtotal

		// Stock movement
		previousStock := productData.Quantity
		newStock := previousStock - item.Quantity

		// Update only quantity
		if err := s.productRepo.UpdateQuantity(
			productData.ID.String(),
			newStock,
		); err != nil {
			return nil, err
		}

		// Keep local object updated
		productData.Quantity = newStock

		// Create inventory log
		log := &inventory.InventoryLog{
			ProductID:     productData.ID,
			MovementType:  inventory.Sale,
			Quantity:      -item.Quantity,
			PreviousStock: previousStock,
			NewStock:      newStock,
			Reason:        "Product sold",
			CreatedBy:     userID,
		}

		if err := s.inventoryRepo.CreateLog(log); err != nil {
			return nil, err
		}
	}

	// Calculate final total
	sale.TotalAmount = total - sale.Discount + sale.Tax

	// Save sale
	if err := s.repo.Create(sale); err != nil {
		return nil, err
	}

	// Return sale with items
	return s.repo.GetByID(sale.ID.String())
}

func (s *Service) GetAll() ([]Sale, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (*Sale, error) {
	return s.repo.GetByID(id)
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

func (s *Service) GetDeleted() ([]Sale, error) {
	return s.repo.GetDeleted()
}

// Optional helper
func parseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
