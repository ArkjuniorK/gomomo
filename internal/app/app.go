package app

import (
	"context"
	"fmt"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"os"
	"os/signal"
	"syscall"
)

type App interface {
	Run()
}

// app contain all the core dependency and should
// be implemented at application startup.
type app struct {
	Addr string

	API    *core.API
	Logger *core.Logger
	PubSub *core.PubSub
}

func (app *app) Run() {
	api := app.API.GetCore()
	logger := app.Logger.GetCore()
	pubSub := app.PubSub.GetCore()

	logger.Info("App running", "address", app.Addr)
	go func() { _ = api.Listen(app.Addr) }()

	var sig os.Signal
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	sig = <-c                                       // This blocks the main thread until an interrupt is received

	logger.Info("Signal received", "signal", sig.String())
	logger.Info("Shutting down app, waiting background process to finish")
	defer logger.Info("App shutdown")

	_ = api.ShutdownWithContext(context.Background())
	_ = pubSub.Close()
}

// New would implement the App interface by
// initialize the app's core dependencies and packages.
func New(host, port string) App {
	var app = new(app)

	app.Addr = fmt.Sprintf("%s:%s", host, port)
	app.initCore()
	app.initPackages()

	return app
}
