package reports

import "time"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSalesSummary() (*SalesSummary, error) {
	return s.repo.GetSalesSummary()
}

func (s *Service) GetDailySalesReport() (*DailySalesReport, error) {
	return s.repo.GetDailySalesReport()
}

func (s *Service) GetMonthlySalesReport() (*MonthlySalesReport, error) {
	return s.repo.GetMonthlySalesReport()
}

func (s *Service) GetLowStockReport() ([]LowStockReport, error) {
	return s.repo.GetLowStockReport()
}

func (s *Service) GetTopSellingProducts() ([]TopProductReport, error) {
	return s.repo.GetTopSellingProducts()
}

func (s *Service) GetDateRangeReport(start, end time.Time) (*DateRangeReport, error) {
	return s.repo.GetDateRangeReport(start, end)
}
