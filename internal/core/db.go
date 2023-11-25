package core

import (
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
)

type Database struct {
	core *gorm.DB
}

func NewDatabase(l *Logger) *Database {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logger := slogGorm.New(slogGorm.WithLogger(l.GetCore()), slogGorm.WithTraceAll())
	option := &gorm.Config{Logger: logger}

	dialect := sqlite.Open(path.Join(wd, "shipping.db"))
	db, err := gorm.Open(dialect, option)
	if err != nil {
		panic(err)
	}

	return &Database{core: db}

}

func (db *Database) GetCore() *gorm.DB {
	return db.core
}
