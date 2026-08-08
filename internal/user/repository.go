package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUsers() ([]User, error) {
	var users []User

	err := r.db.Order("created_at DESC").Find(&users).Error

	return users, err
}

func (r *Repository) GetUserByID(id string) (*User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *Repository) UpdateStatus(id string, active bool) error {
	return r.db.Model(&User{}).
		Where("id = ?", id).
		Update("is_active", active).Error
}

func (r *Repository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func (r *Repository) GetUserByEmailExceptID(
	email string,
	id string,
) (*User, error) {
	var user User

	err := r.db.
		Where("email = ? AND id != ?", email, id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *Repository) GetActiveUsersByRoles(roles ...Role) ([]User, error) {
	var users []User

	err := r.db.
		Where("role IN ?", roles).
		Where("is_active = ?", true).
		Order("created_at DESC").
		Find(&users).Error

	return users, err
}
