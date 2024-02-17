package auth

import "github.com/ArkjuniorK/gofiber-boilerplate/internal/core"

func New(db *core.Database, api *core.API, l *core.Logger, ps *core.PubSub) {
	var (
		rpo = newRepository(db)
		svc = newService(rpo)
		hdl = newHandler(svc)
		rtr = newRouter(api, ps, l, hdl)
	)

	rtr.Serve()
}
