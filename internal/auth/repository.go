package auth

import (
	"github.com/okakafavour/supermarket-pos-backend/internal/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByEmail(email string) (*user.User, error) {

	var user user.User

	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id string) (*user.User, error) {
	var u user.User

	err := r.db.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
