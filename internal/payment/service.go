package payment

import (
	"errors"
	"fmt"
	"time"

	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
)

type Service struct {
	repo     *Repository
	saleRepo *sale.Repository
}

func NewService(
	repo *Repository,
	saleRepo *sale.Repository,
) *Service {
	return &Service{
		repo:     repo,
		saleRepo: saleRepo,
	}
}

func generatePaymentReference() string {
	now := time.Now()

	return fmt.Sprintf(
		"PAY-%s-%04d",
		now.Format("20060102"),
		now.Unix()%10000,
	)
}

// Create Payment
func (s *Service) Create(req CreatePaymentRequest) (*Payment, error) {

	// Check duplicate payment
	existing, _ := s.repo.GetBySaleID(req.SaleID.String())
	if existing != nil {
		return nil, errors.New("payment already exists for this sale")
	}

	// Validate amount
	if req.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	// Get sale
	saleData, err := s.saleRepo.GetByID(req.SaleID.String())
	if err != nil {
		return nil, errors.New("sale not found")
	}

	// Validate payment amount
	if req.Amount != saleData.TotalAmount {
		return nil, fmt.Errorf(
			"payment amount must be %.2f",
			saleData.TotalAmount,
		)
	}

	payment := &Payment{
		SaleID:    req.SaleID,
		Amount:    req.Amount,
		Method:    req.Method,
		Status:    Paid,
		Reference: generatePaymentReference(),
		PaidAt:    time.Now(),
	}

	// Save payment
	if err := s.repo.Create(payment); err != nil {
		return nil, err
	}

	// ==============================
	// UPDATE SALE STATUS HERE
	// ==============================
	saleData.Status = sale.SalePaid

	if err := s.saleRepo.Update(saleData); err != nil {
		return nil, err
	}

	// Return payment with sale preloaded
	return s.repo.GetByID(payment.ID.String())
}

// Get All Payments
func (s *Service) GetAll() ([]Payment, error) {
	return s.repo.GetAll()
}

// Get Payment By ID
func (s *Service) GetByID(id string) (*Payment, error) {
	return s.repo.GetByID(id)
}

// Soft Delete
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

// Restore
func (s *Service) Restore(id string) error {
	return s.repo.Restore(id)
}

// Permanent Delete
func (s *Service) PermanentDelete(id string) error {
	return s.repo.PermanentDelete(id)
}

// Get Deleted Payments
func (s *Service) GetDeleted() ([]Payment, error) {
	return s.repo.GetDeleted()
}
