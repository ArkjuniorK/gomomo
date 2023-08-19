package handlers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
	"github.com/gofiber/fiber/v2"
)

type Converter struct {
	Service    *services.Service
	Repository *repositories.Repository
}

func NewConverterHandler(svc *services.Service, repo *repositories.Repository) *Converter {
	return &Converter{
		Service:    svc,
		Repository: repo,
	}
}

func (c *Converter) ConvertToBase64(ctx *fiber.Ctx) error {
	return nil
}
