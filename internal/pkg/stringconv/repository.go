package stringconv

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/config"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *config.Database) *Repository {
	core := db.Core
	return &Repository{db: core}
}
