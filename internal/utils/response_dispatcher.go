package utils

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/schema"
	"github.com/gofiber/fiber/v2"
)

func ResponseDispatcher(data, pagination interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rs := new(schema.Response)
		rs.Msg = "success"
		rs.Data = data

		if pagination != nil {
			rs.Pagination = pagination
		}

		return c.JSON(rs)
	}
}
