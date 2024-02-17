package auth

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/core"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/model"
)

type repository struct {
	db *core.Database
}

func newRepository(db *core.Database) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user *model.User) error {
	return nil
}

func (r *repository) FindUserByUsername(username string) (*model.User, error) {
	return nil, nil
}
