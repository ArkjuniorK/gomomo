package routers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/handlers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	router fiber.Router
}

func New(name string, engine *config.Engine) *Router {
	router := engine.App.Group(name)
	return &Router{router}
}

// Handle would contain all the route for each sub router for example the 'converter'
// is sub route of 'stringconv' package. This is the only function that should be
// called at the package init file.
func (r *Router) Handle(handler *handlers.Handler) {

	{
		var (
			hdl = handler.Converter
			rtr = r.router.Group("converter")
		)

		rtr.Post("to-base64", hdl.ConvertToBase64)
	}

}
