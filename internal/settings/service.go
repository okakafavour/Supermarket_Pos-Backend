package settings

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetSettings returns the current settings.
// If no settings exist, the repository creates the default settings.
func (s *Service) GetSettings() (*Settings, error) {
	return s.repo.GetOrCreateSettings()
}

// UpdateSettings updates the current store settings.
func (s *Service) UpdateSettings(req UpdateSettingsRequest) (*Settings, error) {
	settings, err := s.repo.GetOrCreateSettings()
	if err != nil {
		return nil, err
	}

	settings.StoreName = req.StoreName
	settings.StoreEmail = req.StoreEmail
	settings.StorePhone = req.StorePhone
	settings.StoreAddress = req.StoreAddress

	settings.Currency = req.Currency
	settings.TaxRate = req.TaxRate
	settings.LowStockThreshold = req.LowStockThreshold

	settings.Timezone = req.Timezone

	settings.ReceiptHeader = req.ReceiptHeader
	settings.ReceiptFooter = req.ReceiptFooter

	if err := s.repo.UpdateSettings(settings); err != nil {
		return nil, err
	}

	return settings, nil
}
