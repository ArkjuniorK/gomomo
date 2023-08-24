package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/handlers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/routers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
)

func New(name string, engine *config.Engine, db *config.Database, logger *config.Logger) {
	var (
		repo    = repositories.New(db)
		svc     = services.New(repo, logger)
		handler = handlers.New(svc, repo)
		router  = routers.New(name, engine)
	)

	router.Handle(handler)
}
