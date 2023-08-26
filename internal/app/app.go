package app

import (
	"fmt"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/helper"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Addr     string
	Router   fiber.Router
	Engine   *config.Engine
	Logger   *config.Logger
	Database *config.Database
}

// New would initialize  database connection, logger, engine and the packages.
// This function would be exported and used by the cmd file.
func New(host, port string) *App {
	var (
		addr     = fmt.Sprintf("%v:%v", host, port)
		logger   = config.NewLogger()
		database = config.NewDatabase(logger)
		engine   = config.NewEngine(logger)
		router   = engine.Core.Name(helper.AppName)
	)

	return &App{
		Addr:     addr,
		Router:   router,
		Engine:   engine,
		Logger:   logger,
		Database: database,
	}
}

func (app *App) Run() {
	app.handlePackages()
	app.handleApp()
}

func (app *App) handlePackages() {
	app.Logger.Core.Sugar().Infoln("Initializing packages...")
	defer app.Logger.Core.Sugar().Infoln("Packages initialized")

	cfg := &pkg.Config{
		Router:   app.Router.Group("v1"),
		Logger:   app.Logger,
		Database: app.Database,
	}

	stringconv.New(cfg)
}

func (app *App) handleApp() {
	defer func(Core *zap.Logger) {
		_ = Core.Sync()
	}(app.Logger.Core)

	app.Logger.Core.Sugar().Infoln("App running at", app.Addr)
	go func() { _ = app.Engine.Core.Listen(app.Addr) }()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	sig := <-c // This blocks the main thread until an interrupt is received
	app.Logger.Core.Sugar().Infoln("Shutting down app because", sig)
	app.Logger.Core.Sugar().Infoln("Running cleanup tasks...")

	_ = app.Engine.Core.Shutdown()

	app.Logger.Core.Sugar().Infoln("App successfully shutdown.")
}
