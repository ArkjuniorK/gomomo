package auth

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
)

type router struct {
	api    *core.API
	pubsub *core.PubSub
	logger *core.Logger

	handler *handler
}

func newRouter(a *core.API, p *core.PubSub, l *core.Logger, h *handler) *router {
	return &router{
		api:     a,
		pubsub:  p,
		handler: h,
		logger:  l,
	}
}

func (r *router) Serve() {

	var (
		api     = r.api
		handler = r.handler
	)

	// API routes definition
	{
		apiRouter := api.GetRouter()
		apiRouter.Post("/login", handler.Login)
		apiRouter.Post("/register", handler.Register)
		apiRouter.Delete("/logout", handler.Logout)
	}

	// Pub/Sub routes definition
}
