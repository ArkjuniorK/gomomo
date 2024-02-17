package core

import (
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

// API responsible to deliver communication of
// application and client via http transport.
type API struct {
	core   *fiber.App
	router fiber.Router
}

func NewApi(l *Logger) *API {
	defer l.GetCore().Info("API Initialized")

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	conf := fiber.Config{
		Views:                 html.New(path.Join(wd, "api"), ".html"),
		AppName:               "go-boilerplate",
		ErrorHandler:          errorHandler,
		DisableStartupMessage: true,
	}

	api := fiber.New(conf)

	api.Static("/assets", path.Join(wd, "api", "assets"))
	api.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	router := api.Group("/api").Group("/v1")
	return &API{core: api, router: router}
}

func (r *API) GetCore() *fiber.App {
	return r.core
}

func (r *API) GetRouter() fiber.Router {
	return r.router
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	switch err.Error() {
	case
		os.ErrNotExist.Error(),
		gorm.ErrRecordNotFound.Error():
		code = fiber.StatusNotFound
	}

	rs := make(map[string]interface{})
	rs["msg"] = err.Error()
	rs["data"] = nil

	return c.Status(code).JSON(rs)
}
