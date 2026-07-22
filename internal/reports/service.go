package reports

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
