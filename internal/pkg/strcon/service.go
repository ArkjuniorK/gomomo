package strcon

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"errors"
)

type Service interface {
	// Convert the payload based on given type and return the result and an error.
	Convert(ctx context.Context, typ, pyd string) (rs string, err error)
}

// service should implement interfaces.StringConvService interface
type service struct {
	repo *repository
}

// Convert the payload based on given type and return the result and an error.
func (s *service) Convert(ctx context.Context, typ, pyd string) (rs string, err error) {

	switch typ {

	case "base64":
		rs = base64.StdEncoding.EncodeToString([]byte(pyd))
		return

	case "base32":
		rs = base32.StdEncoding.EncodeToString([]byte(pyd))
		return

	}

	return "", errors.New("strcon: no type given")
}

func NewService(repo *repository) Service {
	return &service{repo: repo}
}
