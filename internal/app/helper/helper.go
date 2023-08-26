package helper

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/schema"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// This file would contain all the helper function that used by the entire app

func SendOk(ctx *fiber.Ctx, payload interface{}) error {
	var res = &schema.Response{
		Data:    payload,
		Message: "Ok",
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func SendCreated(ctx *fiber.Ctx, payload interface{}) error {
	var res = &schema.Response{
		Data:    payload,
		Message: "Created",
	}

	return ctx.Status(http.StatusCreated).JSON(res)
}
