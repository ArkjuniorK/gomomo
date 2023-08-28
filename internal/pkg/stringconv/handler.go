package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/helper"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/schema"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service    *Service
	repository *Repository
}

func NewHandler(svc *Service, repo *Repository) *Handler {
	return &Handler{service: svc, repository: repo}
}

func (h *Handler) ConvertToBase64(ctx *fiber.Ctx) error {
	var body *schema.Request
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	payload := body.Payload.(string)
	rs, err := h.service.Convert(ctx.Context(), "base64", payload)
	if err != nil {
		return err
	}

	return helper.SendOk(ctx, rs)
}

func (h *Handler) ConvertToBase32(ctx *fiber.Ctx) error {
	var body *schema.Request
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	payload := body.Payload.(string)
	rs, err := h.service.Convert(ctx.Context(), "base32", payload)
	if err != nil {
		return err
	}

	return helper.SendOk(ctx, rs)
}
