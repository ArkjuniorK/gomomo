package handlers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
	"github.com/gofiber/fiber/v2"
)

type Converter struct {
	service    services.Service
	repository *repositories.Repository
}

func NewConverterHandler(svc services.Service, repo *repositories.Repository) *Converter {
	return &Converter{
		service:    svc,
		repository: repo,
	}
}

func (c *Converter) ConvertToBase64(ctx *fiber.Ctx) error {
	return nil
}
