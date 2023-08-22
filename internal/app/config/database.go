package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database wrapper, the Core could be replaced with any kind of database/orm
// but for now the default core is *gorm.DB.
type Database struct {
	Core *gorm.DB // or *Mongo.Database
}

func NewDatabase(logger *Logger) *Database {
	logger.Core.Sugar().Infoln("Initializing database connection...")
	defer logger.Core.Sugar().Infoln("Database initialized!")

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	conf := &gorm.Config{}

	core, err := gorm.Open(postgres.Open(dsn), conf)
	if err != nil {
		logger.Core.Fatal("Unable connecting database", zap.Error(err))
		return nil
	}

	return &Database{Core: core}
}
