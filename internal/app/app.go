package app

import (
	"fmt"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

// New would initialize  database connection, logger, engine and the packages.
// This function would be exported and used by the cmd file at '/cmd/cmd.go'.
func New(host, port string) {
	var (
		addr     = fmt.Sprintf("%v:%v", host, port)
		logger   = config.NewLogger()
		database = config.NewDatabase(logger)
		engine   = config.NewEngine(logger)
	)

	// register global middleware
	// engine.App.Use()

	v1 := engine.App.Group("v1")
	handlePackages(v1, database, logger)

	handleApp(addr, engine, logger)
}

func handlePackages(router fiber.Router, database *config.Database, logger *config.Logger) {
	logger.Core.Sugar().Infoln("Initializing packages...")
	defer logger.Core.Sugar().Infoln("Packages initialized")

	stringconv.New(router, database, logger)
}

func handleApp(addr string, engine *config.Engine, logger *config.Logger) {
	defer logger.Core.Sync()

	logger.Core.Sugar().Infoln("Service running at ", addr)
	go func() { engine.App.Listen(addr) }()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	sig := <-c // This blocks the main thread until an interrupt is received
	logger.Core.Sugar().Infoln("Shutting down service because", sig)
	logger.Core.Sugar().Infoln("Running cleanup tasks...")

	engine.App.Shutdown()

	logger.Core.Sugar().Infoln("Service was successful shutdown.")
}
