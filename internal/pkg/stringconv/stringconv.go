package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/handlers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/routers"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
)

type Package struct {
	// Service would usually be used by other module/package so,
	// it is make sense to export the service
	Service services.Service
}

func New(name string, engine *config.Engine, db *config.Database, logger *config.Logger) *Package {
	var (
		repo    = repositories.New(db)
		svc     = services.New(repo, logger)
		handler = handlers.New(svc, repo)
		router  = routers.New(name, engine)
	)

	router.Route(handler)
	return &Package{Service: svc}
}
