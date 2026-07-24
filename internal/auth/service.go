package auth

import (
	"errors"

	"github.com/okakafavour/supermarket-pos-backend/internal/user"
	"github.com/okakafavour/supermarket-pos-backend/pkg/helpers"
	jwtutil "github.com/okakafavour/supermarket-pos-backend/pkg/jwt"
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

func (s *Service) Register(req RegisterRequest) error {

	existingUser, err := s.repo.GetUserByEmail(req.Email)

	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	role := req.Role

	if role == "" {
		role = user.Cashier
	}

	newUser := user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPassword,
		Role:      role,
		IsActive:  true,
	}

	return s.repo.CreateUser(&newUser)
}

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = helpers.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwtutil.GenerateToken(
		user.ID.String(),
		string(user.Role),
	)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *Service) Profile(userID string) (*user.User, error) {
	return s.repo.GetUserByID(userID)
}
