package routers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func NewStrconvRouter(handler *handlers.StrconvHandler) *fiber.App {
	router := fiber.New()
	router.Post("/convert/:type", handler.Convert)
	return router
}
