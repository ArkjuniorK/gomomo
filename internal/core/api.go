package core

import "github.com/gofiber/fiber/v2"

// API responsible to deliver communication of
// application and client via http transport.
type API struct {
	core   *fiber.App
	router fiber.Router
}

func NewApi(l *Logger) *API {
	defer l.GetCore().Info("API Initialized")

	conf := fiber.Config{
		AppName:               "go-boilerplate",
		DisableStartupMessage: true,
	}

	api := fiber.New(conf)
	router := api.Group("/api").Group("/v1")
	return &API{core: api, router: router}
}

func (r *API) GetCore() *fiber.App {
	return r.core
}

func (r *API) GetRouter() fiber.Router {
	return r.router
}
