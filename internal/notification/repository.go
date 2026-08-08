package notification

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(notification *Notification) error {
	return r.db.Create(notification).Error
}

func (r *Repository) GetByUserID(
	userID string,
) ([]Notification, error) {

	var notifications []Notification

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error

	return notifications, err
}

func (r *Repository) GetUnreadCount(
	userID string,
) (int64, error) {

	var count int64

	err := r.db.
		Model(&Notification{}).
		Where(
			"user_id = ? AND is_read = ?",
			userID,
			false,
		).
		Count(&count).Error

	return count, err
}

func (r *Repository) MarkAsRead(
	id string,
	userID string,
) error {

	return r.db.
		Model(&Notification{}).
		Where(
			"id = ? AND user_id = ?",
			id,
			userID,
		).
		Update("is_read", true).Error
}

func (r *Repository) MarkAllAsRead(
	userID string,
) error {

	return r.db.
		Model(&Notification{}).
		Where(
			"user_id = ? AND is_read = ?",
			userID,
			false,
		).
		Update("is_read", true).Error
}

func (r *Repository) Delete(
	id string,
	userID string,
) error {

	return r.db.
		Where(
			"id = ? AND user_id = ?",
			id,
			userID,
		).
		Delete(&Notification{}).Error
}

func (r *Repository) DeleteAll(
	userID string,
) error {

	return r.db.
		Where("user_id = ?", userID).
		Delete(&Notification{}).Error
}
