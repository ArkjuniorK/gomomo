package middleware

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/gofiber/fiber/v2"
	recoverfiber "github.com/gofiber/fiber/v2/middleware/recover"
)

// Declare and register API middleware, make sure it returns fiber.Handler
// type otherwise it won't be added to router stack.

// InitAPIMiddleware would inject the middleware to application API,
// middleware could also pass as arguments.
func InitAPIMiddleware(api *core.API, mws ...fiber.Handler) {
	router := api.GetRouter()

	for _, mw := range mws {
		router.Use(mw)
	}

	router.Use(
		recoverfiber.New(),
	)
}
