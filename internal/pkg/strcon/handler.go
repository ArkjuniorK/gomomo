package strcon

import (
	"context"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/schema"
	"github.com/gofiber/fiber/v2"
)

// handler contains handling function for both http and pub-sub request,
// use the return type to differentiate whether a function is
// related to one request or not and its optional.
type handler struct {
	service *service
}

func newHandler(svc *service) *handler {
	return &handler{svc}
}

func (h *handler) Convert() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
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
}
