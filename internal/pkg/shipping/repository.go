package shipping

import "github.com/ArkjuniorK/gofiber-boilerplate/internal/core"

type repository struct {
	db *core.Database
}

func newRepository(db *core.Database) *repository {
	return &repository{db}
}
