package routers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func NewStrconRouter(handler *handlers.StrconHandler) *fiber.App {
	router := fiber.New()
	router.Post("/convert/:type", handler.Convert)
	return router
}
