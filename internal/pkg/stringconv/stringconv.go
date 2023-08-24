package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/handlers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
	"github.com/gofiber/fiber/v2"
)

const name = "stringconv"

func New(router fiber.Router, db *config.Database, logger *config.Logger) {
	var (
		repo    = repositories.New(db)
		svc     = services.New(repo, logger)
		handler = handlers.New(svc, repo)
	)

	r := router.Group(name)
	r.Route("converter", func(router fiber.Router) {
		router.Post("to-base64", handler.Converter.ConvertToBase64)
	})

}
