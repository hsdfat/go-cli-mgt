package server

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
)

func ListenAndServe(app *fiber.App) {
	cfg := config.GetServerConfig()

	err := app.Listen(cfg.Host + ":" + cfg.Port)
	if err != nil {
		logger.Logger.Fatalf("Can't listen server: %v", err)
	}
}
