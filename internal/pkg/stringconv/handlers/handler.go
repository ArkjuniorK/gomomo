package handlers

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/repositories"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/pkg/stringconv/services"
)

type Handler struct {
	Converter *Converter
}

func New(svc services.Service, repo *repositories.Repository) *Handler {
	return &Handler{
		Converter: NewConverterHandler(svc, repo),
	}
}
