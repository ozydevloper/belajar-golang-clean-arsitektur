package config

import (
	"github.com/gofiber/fiber/v3"
)

func NewApp(cfg ConfigApp) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cfg.Name,
		ReadTimeout:  cfg.ReadTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	return app
}
