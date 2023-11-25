package strcon

import "github.com/ArkjuniorK/gofiber-boilerplate/internal/core"

func New(api *core.API, l *core.Logger, ps *core.PubSub) {
	var (
		rpo = newRepository()
		svc = newService(rpo)
		hdl = newHandler(svc)
		rtr = newRouter(api, ps, l, hdl)
	)

	rtr.Serve()
}
