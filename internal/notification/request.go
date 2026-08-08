package notification

type CreateNotificationRequest struct {
	UserID  string           `json:"user_id"`
	Type    NotificationType `json:"type" binding:"required"`
	Title   string           `json:"title" binding:"required"`
	Message string           `json:"message" binding:"required"`
}
