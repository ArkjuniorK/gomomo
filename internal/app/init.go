package app

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/mws"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/strcon"
	slogfiber "github.com/samber/slog-fiber"
)

// initCore initialize core packages of application such as router, logger, database client, etc.
func (app *app) initCore() {

	logger := core.NewLogger()
	app.Logger = logger

	api := core.NewApi(logger)
	mws.InitAPIMiddleware(api, slogfiber.New(logger.GetCore()))
	app.API = api

	pubSub := core.NewPubSub(logger)
	app.PubSub = pubSub

	defer logger.GetCore().Info("Initialize core packages done!")

}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all
// packages should be initialized.
func (app *app) initPackages() {

	var (
		api    = app.API
		logger = app.Logger
		pubsub = app.PubSub
	)

	strcon.New(api, logger, pubsub)

	defer logger.GetCore().Info("Initialize common packages done!")

}
