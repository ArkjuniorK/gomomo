package strcon

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
)

type router struct {
	api    *core.API
	pubSub *core.PubSub
	logger *core.Logger

	handler *handler
}

func newRouter(a *core.API, p *core.PubSub, l *core.Logger, h *handler) *router {
	return &router{
		api:     a,
		pubSub:  p,
		handler: h,
		logger:  l,
	}
}

func (r *router) Serve() {

	var (
		api    = r.api
		pubSub = r.pubSub

		handler = r.handler
	)

	// API routes definition
	{
		apiRouter := api.GetRouter().Group("/strconv")
		apiRouter.Post("/convert/:type", handler.Convert())
	}

	// Pub/Sub routes definition
	{
		psCore := pubSub.GetCore()
		psRouter := pubSub.GetRouter()
		psRouter.AddHandler(
			"example_handler",
			"convert",
			psCore,
			"convert",
			psCore,
			nil)
	}
}
