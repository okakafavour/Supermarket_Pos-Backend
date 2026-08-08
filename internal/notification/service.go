package notification

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetNotifications(
	userID string,
) ([]Notification, error) {

	return s.repo.GetByUserID(userID)
}

func (s *Service) GetUnreadCount(
	userID string,
) (int64, error) {

	return s.repo.GetUnreadCount(userID)
}

func (s *Service) MarkAsRead(
	id string,
	userID string,
) error {

	return s.repo.MarkAsRead(id, userID)
}

func (s *Service) MarkAllAsRead(
	userID string,
) error {

	return s.repo.MarkAllAsRead(userID)
}

func (s *Service) Delete(
	id string,
	userID string,
) error {

	return s.repo.Delete(id, userID)
}

func (s *Service) DeleteAll(
	userID string,
) error {

	return s.repo.DeleteAll(userID)
}

// Create a notification from another backend service.
func (s *Service) Create(
	userID string,
	notificationType NotificationType,
	title string,
	message string,
) error {

	notification := &Notification{
		UserID:  userID,
		Type:    notificationType,
		Title:   title,
		Message: message,
		IsRead:  false,
	}

	return s.repo.Create(notification)
}
