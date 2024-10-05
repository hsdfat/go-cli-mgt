package server

import (
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/store/repository"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Initialize() *fiber.App {
	err := godotenv.Load("D:/Projects/go/go-cli-mgt/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &models_config.Config{
		Svr: models_config.ServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Db: models_config.DatabaseConfig{
			DbType: os.Getenv("DB_DRIVER"),
			Pgsql: models_config.PostgresConfig{
				Host:     os.Getenv("PG_HOST"),
				Port:     os.Getenv("PG_PORT"),
				User:     os.Getenv("PG_USER"),
				Password: os.Getenv("PG_PASSWORD"),
				DbName:   os.Getenv("PG_DB_NAME"),
				Schema:   os.Getenv("PG_DB_SCHEMA"),
			},
		},
		Log: models_config.LogConfig{
			Level:   os.Getenv("LOG_LEVEL"),
			DbLevel: os.Getenv("DB_LOG_LEVEL"),
		},
	}
	config.Init(cfg)
	logger.Init()
	repository.Init()

	// Initialize Server
	return NewFiber()
}
