package config

import (
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/helper"
	"github.com/ArkjuniorK/gofiber-boilerplate/internal/app/schema"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Engine struct {
	App *fiber.App
}

func NewEngine(logger *Logger) *Engine {
	logger.Core.Sugar().Infoln("Initializing engine...")
	defer logger.Core.Sugar().Infoln("Engine initialized!")

	conf := fiber.Config{}
	conf.ErrorHandler = errorHandler
	conf.ReadTimeout = 10 * time.Second
	conf.WriteTimeout = 10 * time.Second

	return &Engine{App: fiber.New(conf)}
}

// errorHandler is custom error handler that send a JSON response.
func errorHandler(ctx *fiber.Ctx, err error) error {
	resp := &schema.Response{Data: nil, Message: err.Error()}
	code := fiber.StatusInternalServerError

	switch err.Error() {
	case helper.ErrEmptyString.Error():
		code = fiber.StatusBadRequest
	}

	return ctx.Status(code).JSON(&resp)
}
