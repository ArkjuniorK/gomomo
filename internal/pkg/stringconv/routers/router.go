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

// Route would contain all the route for each sub router for example the 'converter'
// is sub route of 'stringconv' package. This is the only function that should be
// called at the package init file.
func (r *Router) Route(handler *handlers.Handler) {
	routeConverter(r.router.Group("converter"), handler)
}
