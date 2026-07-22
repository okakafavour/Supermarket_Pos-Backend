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

func generateInvoiceNumber() string {
	now := time.Now()

	return fmt.Sprintf(
		"INV-%s-%04d",
		now.Format("20060102"),
		now.Unix()%10000,
	)
}

func (s *Service) Create(req CreateSaleRequest, userID string) (*Sale, string, error) {

	if len(req.Items) == 0 {
		return nil, "", errors.New("sale must contain at least one item")
	}

	sale := &Sale{
		InvoiceNumber: generateInvoiceNumber(),
		CustomerName:  req.CustomerName,
		Discount:      req.Discount,
		Tax:           req.Tax,
		PaymentMethod: req.PaymentMethod,
		SoldBy:        userID,
	}

	var (
		total   float64
		warning string
	)

	for _, item := range req.Items {

		productData, err := s.productRepo.GetByID(item.ProductID.String())
		if err != nil {
			return nil, "", errors.New("product not found")
		}

		if productData.Quantity < item.Quantity {
			return nil, "", fmt.Errorf(
				"not enough stock for %s",
				productData.Name,
			)
		}

		subtotal := float64(item.Quantity) * productData.SellingPrice

		sale.Items = append(sale.Items, SaleItem{
			ProductID: productData.ID,
			Quantity:  item.Quantity,
			UnitPrice: productData.SellingPrice,
			Subtotal:  subtotal,
		})

		total += subtotal

		previousStock := productData.Quantity
		newStock := previousStock - item.Quantity

		if err := s.productRepo.UpdateQuantity(
			productData.ID.String(),
			newStock,
		); err != nil {
			return nil, "", err
		}

		productData.Quantity = newStock

		if newStock <= productData.MinimumStock {
			warning = fmt.Sprintf(
				"%s is below minimum stock (%d left)",
				productData.Name,
				newStock,
			)
		}

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
			return nil, "", err
		}
	}

	sale.TotalAmount = total - sale.Discount + sale.Tax

	if err := s.repo.Create(sale); err != nil {
		return nil, "", err
	}

	createdSale, err := s.repo.GetByID(sale.ID.String())
	if err != nil {
		return nil, "", err
	}

	return createdSale, warning, nil
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

func parseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
