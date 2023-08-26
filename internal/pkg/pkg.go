package pkg

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/gofiber/fiber/v2"
)

// Config take all the config required by packages as single source
type Config struct {
	Router   fiber.Router
	Logger   *config.Logger
	Database *config.Database
}
