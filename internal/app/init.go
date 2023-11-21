package app

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/handlers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/api/routers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/strcon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

// initCore initialize core packages of application such as router, logger, database client, etc.
// This function could also handle middleware registration for application router
// before registering each package to handlers and routers.
func (app *app) initCore() {

	logger := core.NewLogger()
	app.Logger = logger

	router := core.NewRouter()
	rtrCore := router.GetCore()
	rtrCore.Use(slogfiber.New(logger.GetCore()))
	rtrCore.Use(recover.New())
	app.Router = router

	defer logger.GetCore().Info("Setup core done!")
}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all packages should only
// initialize here
func (app *app) initPackages() {
	// init common package from inside out
	// starting from the service, handler then router
	// that would be registered to application router
	logger := app.Logger.GetCore()
	router := app.Router.GetCore().Group("/api").Group("/v1")

	{
		svc := strcon.New()
		hdl := handlers.NewStrconHandler(svc)
		rtr := routers.NewStrconRouter(hdl)
		router.Mount("/strcon", rtr)
	}

	defer logger.Info("Setup packages done!")
}
