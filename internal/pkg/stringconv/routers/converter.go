package routers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/handlers"
	"github.com/gofiber/fiber/v2"
)

func routeConverter(router fiber.Router, handler *handlers.Handler) {
	router.Post("/convert-to-base64", handler.Converter.ConvertToBase64)
}
