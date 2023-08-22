package services

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
)

type Service interface{}

type service struct {
	repo   *repositories.Repository
	logger *config.Logger
}

func New(repo *repositories.Repository, logger *config.Logger) Service {
	return &service{repo, logger}
}
