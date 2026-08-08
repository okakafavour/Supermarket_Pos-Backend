package user

import (
	"errors"

	"github.com/okakafavour/supermarket-pos-backend/pkg/helpers"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *Service) GetUserByID(id string) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) CreateUser(req CreateUserRequest) error {

	existing, err := s.repo.GetUserByEmail(req.Email)

	if err == nil && existing != nil {
		return errors.New("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	password, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  password,
		Role:      req.Role,
		IsActive:  true,
	}

	return s.repo.CreateUser(&user)
}

func (s *Service) UpdateUser(id string, req UpdateUserRequest) error {

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone
	user.Role = req.Role

	return s.repo.UpdateUser(user)
}

func (s *Service) UpdateStatus(id string, req UpdateUserStatusRequest) error {

	_, err := s.repo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.repo.UpdateStatus(id, req.IsActive)
}

func (s *Service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *Service) UpdateProfile(
	id string,
	req UpdateProfileRequest,
) (*User, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	// Check whether another user already has this email.
	existing, err := s.repo.GetUserByEmailExceptID(
		req.Email,
		id,
	)

	if err == nil && existing != nil {
		return nil, errors.New("email already exists")
	}

	if err != nil &&
		!errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
