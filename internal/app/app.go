package app

import (
	"fmt"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv"
	"go.uber.org/zap"
)

// New would initialize  database connection, logger, engine and the packages.
// This function would be exported and used by the cmd file at '/cmd/cmd.go'.
//
// TODO: Implement graceful shutdown
func New(host, port string) {
	var (
		addr     = fmt.Sprintf("%v:%v", host, port)
		logger   = config.NewLogger()
		database = config.NewDatabase(logger)
		engine   = config.NewEngine(logger)
	)

	initPackages(engine, database, logger)

	err := engine.App.Listen(addr)
	if err != nil {
		logger.Core.Fatal("Unable to start app", zap.Error(err))
	}
}

func initPackages(engine *config.Engine, database *config.Database, logger *config.Logger) {
	logger.Core.Sugar().Infoln("Initializing packages...")
	defer logger.Core.Sugar().Infoln("Packages initialized")

	stringconv.New("stringconv", engine, database, logger)
}
