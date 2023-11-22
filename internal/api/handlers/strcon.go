package handlers

import (
	"context"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/schema"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/strcon"
	"github.com/gofiber/fiber/v2"
)

type StrconvHandler struct {
	service strcon.Service
}

func (h *StrconvHandler) Convert(ctx *fiber.Ctx) error {
	var body *schema.Request
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	typ := ctx.Params("type")
	pyd := body.Payload.(string)

	rs, err := h.service.Convert(ctx.Context(), typ, pyd)
	if err != nil {
		return err
	}

	c := context.WithValue(context.Background(), "data", rs)
	ctx.SetUserContext(c)
	return nil
}

func NewStrconvHandler(service strcon.Service) *StrconvHandler {
	return &StrconvHandler{service: service}
}
