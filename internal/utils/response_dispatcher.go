package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseDispatcher(c *fiber.Ctx, data, pagination interface{}) error {
	rs := make(fiber.Map)
	rs["msg"] = "success"
	rs["data"] = data

	if pagination != nil {
		rs["pagination"] = pagination
	}

	return c.JSON(rs)
}
