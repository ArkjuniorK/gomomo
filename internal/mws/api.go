package mws

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/schema"
	"github.com/gofiber/fiber/v2"
	recoverfiber "github.com/gofiber/fiber/v2/middleware/recover"
)

// Declare and register API middleware, make sure it returns fiber.Handler
// type otherwise it won't be added to router stack.

// InitAPIMiddleware would inject the middleware to application API,
// middleware could also pass as arguments.
func InitAPIMiddleware(api *core.API, mws ...fiber.Handler) {

	router := api.GetRouter()

	if len(mws) != 0 {
		for _, mw := range mws {
			router.Use(mw)
		}
	}

	router.Use(
		recoverfiber.New(),
		responseDispatcher(),
	)

}

// responseDispatcher useful to wrap the response to *schema.Response
// so each response have unified structure.
func responseDispatcher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return err
		}

		data := ctx.UserContext().Value("data")

		rs := new(schema.Response)
		rs.Msg = "success"
		rs.Data = data

		return ctx.JSON(rs)
	}
}
