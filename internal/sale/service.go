package sale

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/notification"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
)

type Service struct {
	repo                *Repository
	productRepo         *product.Repository
	inventoryRepo       *inventory.Repository
	notificationService *notification.Service
}

func NewService(
	repo *Repository,
	productRepo *product.Repository,
	inventoryRepo *inventory.Repository,
	notificationService *notification.Service,
) *Service {
	return &Service{
		repo:                repo,
		productRepo:         productRepo,
		inventoryRepo:       inventoryRepo,
		notificationService: notificationService,
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

// =====================================================
// CREATE SALE
// =====================================================

func (s *Service) Create(
	req CreateSaleRequest,
	userID string,
) (*Sale, string, error) {

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

		// =================================================
		// GET PRODUCT
		// =================================================

		productData, err := s.productRepo.GetByID(
			item.ProductID.String(),
		)

		if err != nil {
			return nil, "", errors.New("product not found")
		}

		// =================================================
		// CHECK STOCK
		// =================================================

		if productData.Quantity < item.Quantity {
			return nil, "", fmt.Errorf(
				"not enough stock for %s",
				productData.Name,
			)
		}

		// =================================================
		// CALCULATE SUBTOTAL
		// =================================================

		subtotal := float64(item.Quantity) *
			productData.SellingPrice

		sale.Items = append(
			sale.Items,
			SaleItem{
				ProductID: productData.ID,
				Quantity:  item.Quantity,
				UnitPrice: productData.SellingPrice,
				Subtotal:  subtotal,
			},
		)

		total += subtotal

		// =================================================
		// UPDATE STOCK
		// =================================================

		previousStock := productData.Quantity
		newStock := previousStock - item.Quantity

		if err := s.productRepo.UpdateQuantity(
			productData.ID.String(),
			newStock,
		); err != nil {
			return nil, "", err
		}

		productData.Quantity = newStock

		// =================================================
		// INVENTORY NOTIFICATION
		// =================================================

		if newStock == 0 {

			warning = fmt.Sprintf(
				"%s is out of stock",
				productData.Name,
			)

			if s.notificationService != nil {

				err := s.notificationService.Create(
					userID,
					notification.OutOfStockNotification,
					"Out of Stock",
					fmt.Sprintf(
						"%s is now out of stock.",
						productData.Name,
					),
				)

				if err != nil {
					return nil, "", err
				}
			}

		} else if newStock <= productData.MinimumStock {

			warning = fmt.Sprintf(
				"%s is below minimum stock (%d left)",
				productData.Name,
				newStock,
			)

			if s.notificationService != nil {

				err := s.notificationService.Create(
					userID,
					notification.LowStockNotification,
					"Low Stock Alert",
					fmt.Sprintf(
						"%s is below minimum stock. Only %d left.",
						productData.Name,
						newStock,
					),
				)

				if err != nil {
					return nil, "", err
				}
			}
		}

		// =================================================
		// INVENTORY LOG
		// =================================================

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

	// =====================================================
	// TOTAL
	// =====================================================

	sale.TotalAmount =
		total -
			sale.Discount +
			sale.Tax

	// =====================================================
	// CREATE SALE
	// =====================================================

	if err := s.repo.Create(sale); err != nil {
		return nil, "", err
	}

	// =====================================================
	// GET CREATED SALE
	// =====================================================

	createdSale, err := s.repo.GetByID(
		sale.ID.String(),
	)

	if err != nil {
		return nil, "", err
	}

	return createdSale, warning, nil
}

// =====================================================
// GET ALL SALES
// =====================================================

func (s *Service) GetAll(
	filter SaleFilter,
) ([]Sale, Pagination, error) {

	sales, total, err := s.repo.GetAll(filter)

	if err != nil {
		return nil, Pagination{}, err
	}

	totalPages := 0

	if total > 0 && filter.Limit > 0 {
		totalPages = int(
			(total + int64(filter.Limit) - 1) /
				int64(filter.Limit),
		)
	}

	pagination := Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return sales, pagination, nil
}

// =====================================================
// GET SALE BY ID
// =====================================================

func (s *Service) GetByID(
	id string,
) (*Sale, error) {

	return s.repo.GetByID(id)
}

// =====================================================
// SOFT DELETE
// =====================================================

func (s *Service) Delete(
	id string,
) error {

	return s.repo.Delete(id)
}

// =====================================================
// RESTORE
// =====================================================

func (s *Service) Restore(
	id string,
) error {

	return s.repo.Restore(id)
}

// =====================================================
// PERMANENT DELETE
// =====================================================

func (s *Service) PermanentDelete(
	id string,
) error {

	return s.repo.PermanentDelete(id)
}

// =====================================================
// GET DELETED SALES
// =====================================================

func (s *Service) GetDeleted() ([]Sale, error) {

	return s.repo.GetDeleted()
}

// =====================================================
// UUID PARSER
// =====================================================

func parseUUID(
	id string,
) (uuid.UUID, error) {

	return uuid.Parse(id)
}

// =====================================================
// DASHBOARD
// =====================================================

func (s *Service) GetDashboard() (
	*DashboardResponse,
	error,
) {

	return s.repo.GetDashboard()
}

// =====================================================
// ANALYTICS
// =====================================================

func (s *Service) GetAnalytics() (
	*AnalyticsResponse,
	error,
) {

	return s.repo.GetAnalytics()
}
