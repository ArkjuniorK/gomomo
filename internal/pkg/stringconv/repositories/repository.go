package repositories

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
)

type Repository struct {
}

func New(db *config.Database) *Repository {
	return &Repository{}
}
