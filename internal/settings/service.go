package settings

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSettings() (*Settings, error) {
	return s.repo.GetOrCreateSettings()
}

func (s *Service) UpdateSettings(req UpdateSettingsRequest) (*Settings, error) {
	settings, err := s.repo.GetOrCreateSettings()

	if err != nil {
		return nil, err
	}

	if req.StoreName == "" {
		return nil, errors.New("store name is required")
	}

	if req.Currency == "" {
		return nil, errors.New("currency is required")
	}

	if req.TaxRate < 0 || req.TaxRate > 100 {
		return nil, errors.New("tax rate must be between 0 and 100")
	}

	if req.LowStockThreshold < 0 {
		return nil, errors.New("low stock threshold cannot be negative")
	}

	settings.StoreName = req.StoreName
	settings.StoreEmail = req.StoreEmail
	settings.StorePhone = req.StorePhone
	settings.StoreAddress = req.StoreAddress
	settings.Currency = req.Currency
	settings.TaxRate = req.TaxRate
	settings.ReceiptFooter = req.ReceiptFooter
	settings.LowStockThreshold = req.LowStockThreshold

	if err := s.repo.UpdateSettings(settings); err != nil {
		return nil, err
	}

	return settings, nil
}
