package notification

import "github.com/okakafavour/supermarket-pos-backend/internal/user"

type Service struct {
	repo     *Repository
	userRepo *user.Repository
}

func NewService(
	repo *Repository,
	userRepo *user.Repository,
) *Service {
	return &Service{
		repo:     repo,
		userRepo: userRepo,
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

// Create creates a notification for a specific user.
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

// NotifyAdminsAndManagers creates a notification for
// every active admin and manager.
func (s *Service) NotifyAdminsAndManagers(
	notificationType NotificationType,
	title string,
	message string,
) error {

	users, err := s.userRepo.GetUsers()
	if err != nil {
		return err
	}

	for _, u := range users {

		// Ignore inactive users.
		if !u.IsActive {
			continue
		}

		// Only admins and managers should receive
		// operational notifications.
		if u.Role != user.Admin && u.Role != user.Manager {
			continue
		}

		err := s.Create(
			u.ID.String(),
			notificationType,
			title,
			message,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
