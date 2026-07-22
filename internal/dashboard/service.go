package dashboard

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSummary() (*Summary, error) {

	totalProducts, err := s.repo.CountProducts()
	if err != nil {
		return nil, err
	}

	totalCategories, err := s.repo.CountCategories()
	if err != nil {
		return nil, err
	}

	totalSuppliers, err := s.repo.CountSuppliers()
	if err != nil {
		return nil, err
	}

	totalCustomers, err := s.repo.CountCustomers()
	if err != nil {
		return nil, err
	}

	totalSales, err := s.repo.CountSales()
	if err != nil {
		return nil, err
	}

	todaySales, err := s.repo.CountTodaySales()
	if err != nil {
		return nil, err
	}

	totalRevenue, err := s.repo.TotalRevenue()
	if err != nil {
		return nil, err
	}

	todayRevenue, err := s.repo.TodayRevenue()
	if err != nil {
		return nil, err
	}

	lowStockProducts, err := s.repo.CountLowStockProducts()
	if err != nil {
		return nil, err
	}

	summary := &Summary{
		TotalProducts:    int(totalProducts),
		TotalCategories:  int(totalCategories),
		TotalSuppliers:   int(totalSuppliers),
		TotalCustomers:   int(totalCustomers),
		TotalSales:       int(totalSales),
		TodaySales:       int(todaySales),
		TotalRevenue:     totalRevenue,
		TodayRevenue:     todayRevenue,
		LowStockProducts: int(lowStockProducts),
	}

	return summary, nil
}
