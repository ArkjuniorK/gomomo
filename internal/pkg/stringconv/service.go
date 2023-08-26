package stringconv

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
)

type Service struct {
	repo   *Repository
	logger *config.Logger
}

func NewService(repo *Repository, logger *config.Logger) *Service {
	return &Service{repo, logger}
}

func (s *Service) Convert(ctx context.Context, typ, pyd string) (rs string, err error) {

	switch typ {

	case "base64":
		rs = base64.StdEncoding.EncodeToString([]byte(pyd))
		return

	}

	return "", errors.New("stringconv: no type given")
}
