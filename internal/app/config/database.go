package config

import (
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"path"
)

// Database wrapper, the Core could be replaced with any kind of database/orm
// but for now the default core is *gorm.DB.
type Database struct {
	Core *gorm.DB // or *Mongo.Database
}

func NewDatabase(logger *Logger) *Database {
	logger.Core.Sugar().Infoln("Initializing database connection...")
	defer logger.Core.Sugar().Infoln("Database initialized!")

	conf := &gorm.Config{}
	wd, _ := os.Getwd()
	dbFile := path.Join(wd, "boilerplate.db")

	core, err := gorm.Open(sqlite.Open(dbFile), conf)
	if err != nil {
		logger.Core.Fatal("Unable connecting database", zap.Error(err))
		return nil
	}

	return &Database{Core: core}
}
