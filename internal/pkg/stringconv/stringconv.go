package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	// Service would usually be used by other module/package so,
	// it is make sense to export the service
	Service services.Service
}

func New(router *fiber.App) *Module {
	return &Module{}
}
