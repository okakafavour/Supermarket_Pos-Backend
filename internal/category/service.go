package category

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(req CreateCategoryRequest) (*Category, error) {
	category := &Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id string, req UpdateCategoryRequest) (*Category, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}

	if req.Description != "" {
		category.Description = req.Description
	}

	err = s.repo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
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

func (s *Service) GetDeleted() ([]Category, error) {
	return s.repo.GetDeleted()
}
