package app

import (
	"fmt"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv"
	"go.uber.org/zap"
)

// New initialize the app at this file such as database connection, logger, engine and packages.
// This function would be exported and used by the cmd file at '/cmd/cmd.go'.
// TODO: Implement graceful shutdown
func New(host, port string) {
	var (
		logger   = config.NewLogger()
		database = config.NewDatabase(logger)
		engine   = config.NewEngine(logger)
	)

	// initialized the packages
	stringconv.New("string-conv", engine, database, logger)

	addr := fmt.Sprintf("%v+%v", host, port)
	err := engine.App.Listen(addr)
	if err != nil {
		logger.Core.Fatal("Unable to start app", zap.Error(err))
	}
}
