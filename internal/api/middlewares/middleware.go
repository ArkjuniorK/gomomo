package middlewares

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/schema"
	"github.com/gofiber/fiber/v2"
)

// Define global middleware here, for package's specific middleware
// its more convenient to declared at other file with the same name

func ResponseDispatcher() fiber.Handler {
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
