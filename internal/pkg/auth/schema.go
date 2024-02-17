package auth

import "github.com/ArkjuniorK/gofiber-boilerplate/internal/model"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
