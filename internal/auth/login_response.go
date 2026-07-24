package auth

import "github.com/okakafavour/supermarket-pos-backend/internal/user"

type LoginResponse struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}
