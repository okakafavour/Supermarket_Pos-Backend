package purchase

import (
	"fmt"
	"time"

	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
	"gorm.io/gorm"
)

type Service struct {
	db            *gorm.DB
	repo          *Repository
	supplierRepo  *supplier.Repository
	productRepo   *product.Repository
	inventoryRepo *inventory.Repository
}

func NewService(
	db *gorm.DB,
	repo *Repository,
	supplierRepo *supplier.Repository,
	productRepo *product.Repository,
	inventoryRepo *inventory.Repository,
) *Service {
	return &Service{
		db:            db,
		repo:          repo,
		supplierRepo:  supplierRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
	}
}

func generatePurchaseInvoice() string {
	now := time.Now()

	return fmt.Sprintf(
		"PUR-%s-%04d",
		now.Format("20060102"),
		now.Unix()%10000,
	)
}

// Create Purchase
func (s *Service) Create(
	req CreatePurchaseRequest,
	userID string,
) (*Purchase, error) {

	if len(req.Items) == 0 {
		return nil, fmt.Errorf("purchase must contain at least one item")
	}

	// Validate supplier
	_, err := s.supplierRepo.GetByID(req.SupplierID.String())
	if err != nil {
		return nil, fmt.Errorf("supplier not found")
	}

	purchase := &Purchase{
		InvoiceNumber: generatePurchaseInvoice(),
		SupplierID:    req.SupplierID,
		Status:        Pending,
		CreatedBy:     userID,
	}

	var total float64

	for _, item := range req.Items {

		productData, err := s.productRepo.GetByID(item.ProductID.String())
		if err != nil {
			return nil, fmt.Errorf("product not found")
		}

		subtotal := float64(item.Quantity) * item.UnitCost

		purchase.Items = append(
			purchase.Items,
			PurchaseItem{
				ProductID: productData.ID,
				Quantity:  item.Quantity,
				UnitCost:  item.UnitCost,
				Subtotal:  subtotal,
			},
		)

		total += subtotal
	}

	purchase.TotalAmount = total

	if err := s.repo.Create(purchase); err != nil {
		return nil, err
	}

	return s.repo.GetByID(purchase.ID.String())
}

// Receive Purchase
func (s *Service) Receive(id, userID string) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		purchase, err := s.repo.GetByID(id)
		if err != nil {
			return err
		}

		if purchase.Status == Received {
			return fmt.Errorf("purchase already received")
		}

		for _, item := range purchase.Items {

			productData, err := s.productRepo.GetByID(item.ProductID.String())
			if err != nil {
				return err
			}

			previousStock := productData.Quantity
			newStock := previousStock + item.Quantity

			if err := s.productRepo.UpdateQuantity(
				productData.ID.String(),
				newStock,
			); err != nil {
				return err
			}

			log := &inventory.InventoryLog{
				ProductID:     productData.ID,
				MovementType:  inventory.Restock,
				Quantity:      item.Quantity,
				PreviousStock: previousStock,
				NewStock:      newStock,
				Reason:        "Purchase received",
				CreatedBy:     userID,
			}

			if err := s.inventoryRepo.CreateLog(log); err != nil {
				return err
			}
		}

		now := time.Now()

		purchase.Status = Received
		purchase.ReceivedAt = &now

		if err := s.repo.Update(purchase); err != nil {
			return err
		}

		return nil
	})
}

// Get All Purchases
func (s *Service) GetAll() ([]Purchase, error) {
	return s.repo.GetAll()
}

// Get Purchase By ID
func (s *Service) GetByID(id string) (*Purchase, error) {
	return s.repo.GetByID(id)
}

// Delete Purchase
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
