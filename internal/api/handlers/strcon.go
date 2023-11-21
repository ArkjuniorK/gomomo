package handlers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/schema"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/strcon"
	"github.com/gofiber/fiber/v2"
)

type StrconHandler struct {
	service strcon.Service
}

func (h *StrconHandler) Convert(ctx *fiber.Ctx) error {
	var body *schema.Request
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	typ := ctx.Route().Path
	pyd := body.Payload.(string)

	rs, err := h.service.Convert(ctx.Context(), typ, pyd)
	if err != nil {
		return err
	}

	return ctx.JSON(rs)
}

func NewStrconHandler(service strcon.Service) *StrconHandler {
	return &StrconHandler{service: service}
}
