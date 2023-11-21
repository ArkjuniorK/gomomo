package core

import "github.com/gofiber/fiber/v2"

type Router struct {
	core *fiber.App
}

// NewRouter initialize the Router's core.
func NewRouter() *Router {
	conf := fiber.Config{AppName: "go-modulith-boilerplate", DisableStartupMessage: true}
	core := fiber.New(conf)

	return &Router{core: core}
}

func (r *Router) GetCore() *fiber.App {
	return r.core
}
