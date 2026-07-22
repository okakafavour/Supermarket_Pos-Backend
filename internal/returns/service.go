package returns

import (
	"fmt"

	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"gorm.io/gorm"
)

type Service struct {
	db            *gorm.DB
	repo          *Repository
	saleRepo      *sale.Repository
	productRepo   *product.Repository
	inventoryRepo *inventory.Repository
}

func NewService(
	db *gorm.DB,
	repo *Repository,
	saleRepo *sale.Repository,
	productRepo *product.Repository,
	inventoryRepo *inventory.Repository,
) *Service {

	return &Service{
		db:            db,
		repo:          repo,
		saleRepo:      saleRepo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *Service) Create(
	req CreateReturnRequest,
	userID string,
) (*Return, error) {

	if len(req.Items) == 0 {
		return nil, fmt.Errorf("return must contain at least one item")
	}

	saleData, err := s.saleRepo.GetByID(req.SaleID.String())
	if err != nil {
		return nil, fmt.Errorf("sale not found")
	}

	var refundAmount float64

	returnData := &Return{
		SaleID:      saleData.ID,
		Reason:      req.Reason,
		Status:      ReturnCompleted,
		ProcessedBy: userID,
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {

		for _, item := range req.Items {

			var saleItem *sale.SaleItem

			for i := range saleData.Items {

				if saleData.Items[i].ID == item.SaleItemID {
					saleItem = &saleData.Items[i]
					break
				}
			}

			if saleItem == nil {
				return fmt.Errorf("sale item not found")
			}

			if item.Quantity > saleItem.Quantity {
				return fmt.Errorf("returned quantity exceeds sold quantity")
			}

			productData, err := s.productRepo.GetByID(
				saleItem.ProductID.String(),
			)
			if err != nil {
				return err
			}

			subtotal := float64(item.Quantity) * saleItem.UnitPrice

			refundAmount += subtotal

			returnData.Items = append(
				returnData.Items,
				ReturnItem{
					ProductID:  productData.ID,
					SaleItemID: saleItem.ID,
					Quantity:   item.Quantity,
					UnitPrice:  saleItem.UnitPrice,
					Subtotal:   subtotal,
				},
			)

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
				MovementType:  inventory.Return,
				Quantity:      item.Quantity,
				PreviousStock: previousStock,
				NewStock:      newStock,
				Reason:        "Customer Return",
				CreatedBy:     userID,
			}

			if err := s.inventoryRepo.CreateLog(log); err != nil {
				return err
			}
		}

		returnData.RefundAmount = refundAmount

		if err := s.repo.Create(returnData); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(returnData.ID.String())
}

func (s *Service) GetAll() ([]Return, error) {
	return s.repo.GetAll()
}

func (s *Service) GetDeleted() ([]Return, error) {
	return s.repo.GetDeleted()
}

func (s *Service) GetByID(id string) (*Return, error) {
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
