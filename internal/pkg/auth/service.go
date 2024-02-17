package auth

import (
	"context"
	"os"
	"time"

	"github.com/ArkjuniorK/gofiber-boilerplate/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// service should implement interfaces.StringConvService interface
type service struct {
	repo *repository
}

func newService(repo *repository) *service {
	return &service{repo: repo}
}

// Login authenticates a user by username and password. It returns the user record,
// an auth token, and an error. The token is returned if authentication succeeds,
// the error is returned if authentication fails.
func (s *service) Login(ctx context.Context, payload *LoginRequest, ip string) (*model.User, *string, error) {
	user, err := s.repo.FindUserByUsername(payload.Username)
	if err != nil {
		return nil, nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, nil, err
	}

	k, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, nil, os.ErrNotExist
	}

	claims := jwt.RegisteredClaims{
		Issuer:    ip,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES512, claims).SignedString(k)
	if err != nil {
		return nil, nil, err
	}

	return user, &token, nil
}
