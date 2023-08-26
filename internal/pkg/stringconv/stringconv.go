package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg"
)

func New(config *pkg.Config) {
	var (
		repo = NewRepository(config.Database)
		svc  = NewService(repo, config.Logger)
		hdl  = NewHandler(svc, repo)
		rg   = config.Router.Group("stringconv")
	)

	rg.Post("to-base64", hdl.ConvertToBase64)
}
