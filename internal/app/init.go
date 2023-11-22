package app

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/handlers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/middlewares"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/routers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/strcon"
	recoverfiber "github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

// initCore initialize core packages of application such as router, logger, database client, etc.
// This function could also handle middleware registration for application router
// before registering each package to handlers and routers.
func (app *app) initCore() {

	logger := core.NewLogger()
	app.Logger = logger

	router := core.NewRouter()
	rcore := router.GetCore()

	// official and third-party middlewares
	rcore.Use(slogfiber.New(logger.GetCore()))
	rcore.Use(recoverfiber.New())

	// application middlewares
	rcore.Use(middlewares.ResponseDispatcher())

	app.Router = router

	defer logger.GetCore().Info("Initializing core packages done!")
}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all
// packages should be initialized.
func (app *app) initPackages() {

	logger := app.Logger.GetCore()
	router := app.Router.GetCore().Group("/api").Group("/v1")

	// init packages from inside out starting from
	// service, handler and router then mount
	// the router with package prefix
	{
		svc := strcon.New()
		hdl := handlers.NewStrconvHandler(svc)
		rtr := routers.NewStrconvRouter(hdl)
		router.Mount("/strconv", rtr)
	}

	defer logger.Info("Initializing common packages done!")
}
