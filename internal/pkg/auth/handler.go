package auth

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/utils"
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

func (h *handler) Login(c *fiber.Ctx) error {
	body := new(LoginRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	user, token, err := h.service.Login(c.UserContext(), body, c.IP())
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Path:     "/",
		Name:     "token",
		Value:    *token,
		Secure:   true,
		HTTPOnly: true,
	})

	data := &LoginResponse{user, *token}
	return utils.ResponseDispatcher(c, data, nil)
}

func (h *handler) Register(c *fiber.Ctx) error {
	return nil
}

func (h *handler) Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return nil
}
