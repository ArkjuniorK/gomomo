package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// NOTE: Before using this command make sure to setup database which would be use
// By default it would use SQLite as initial database

const DB_TYPE = "sqlite"

var (
	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "Database migration and seeder",
	}

	dbMigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate to latest",
	}

	dbMigrateCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new migration file",
		Run:   migrateCreateCmd,
	}

	dbMigrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Migrate to previous",
		Run:   migrateUpCmd,
	}

	dbMigrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Database migration using golang-migrate",
		Run:   migrateDownCmd,
	}

	dbSeedCmd = &cobra.Command{
		Use:   "seed",
		Short: "Seed database",
		Long:  "Database seeder for specific driver and database, currently only support GORM",
		Run:   seed,
	}
)

func setupMigration(d *core.Database) (*string, database.Driver, error) {
	db, err := d.GetCore().DB()
	if err != nil {
		return nil, nil, err
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	p := path.Join(wd, "migrations")
	f := "file://" + p

	return &f, driver, nil
}

func migrateCreateCmd(cmd *cobra.Command, args []string) {
	l := core.NewLogger()
	logger := l.GetCore().WithGroup("migration").WithGroup("create")

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logger.Error("err", err)
		return
	}

	if len(name) == 0 {
		logger.Error("Specify name of migration")
		return
	}

	logger.Info("Creating migration file(s)")

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("err", err)
		return
	}

	ext = strings.TrimPrefix(ext, ".")
	dir := path.Join(wd, "migrations")
	ver := strconv.FormatInt(time.Now().Unix(), 10)
	name = strings.Join(strings.Split(name, " "), "_")

	verglob := filepath.Join(dir, ver+"_*."+ext)
	matches, err := filepath.Glob(verglob)
	if err != nil {
		logger.Error("err", err)
		return
	}

	if len(matches) > 0 {
		logger.Error("err", "duplicate migration version: %s", ver)
		return
	}

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%s_%s.%s.%s", ver, name, direction, ext)
		filename := filepath.Join(dir, basename)

		// create exclusive (fails if file already exists)
		// os.Create() specifies 0666 as the FileMode, so we're doing the same
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			logger.Error("err", err)
			return
		}

		if err = f.Close(); err != nil {
			logger.Error("err", err)
			return
		}
	}

	logger.Info("Migration file(s) created")
}

func migrateUpCmd(cmd *cobra.Command, args []string) {
	l := core.NewLogger()
	d := core.NewDatabase(l)

	logger := l.GetCore().WithGroup("migration").WithGroup("up")

	logger.Info("Starting migration")
	defer logger.Info("Migration succeed")

	file, driver, err := setupMigration(d)
	if err != nil {
		logger.Error("err", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(*file, DB_TYPE, driver)
	if err != nil {
		logger.Error("error", err)
		return
	}

	m.Up()
}

func migrateDownCmd(cmd *cobra.Command, args []string) {
	l := core.NewLogger()
	d := core.NewDatabase(l)

	logger := l.GetCore().WithGroup("migration").WithGroup("down")

	logger.Info("Starting migration")
	defer logger.Info("Migration succeed")

	file, driver, err := setupMigration(d)
	if err != nil {
		logger.Error("err", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(*file, DB_TYPE, driver)
	if err != nil {
		logger.Error("err", err)
		return
	}

	m.Down()
}

func seed(cmd *cobra.Command, args []string) {}
