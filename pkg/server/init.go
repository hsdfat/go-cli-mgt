package server

import (
	"fmt"
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
	models_config "go-cli-mgt/pkg/models/config"
	"go-cli-mgt/pkg/store/repository"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Initialize(configFile string) *fiber.App {
	err := godotenv.Load(configFile)
	if err != nil {
		panic(fmt.Errorf("unable to find config %v, %s", configFile, err.Error()))
	}
	fmt.Printf("Loaded configs %v \n", configFile)

	cfg := &models_config.Config{
		Svr: models_config.ServerConfig{
			Host:    os.Getenv("SERVER_HOST"),
			Port:    os.Getenv("SERVER_PORT"),
			TcpPort: os.Getenv("SERVER_TCP_PORT"),
		},
		Db: models_config.DatabaseConfig{
			DbType: os.Getenv("DB_DRIVER"),
			Mysql: models_config.MySqlConfig{
				Host:     os.Getenv("MYSQL_HOST"),
				Port:     os.Getenv("MYSQL_PORT"),
				User:     os.Getenv("MYSQL_USER"),
				Password: os.Getenv("MYSQL_PASSWORD"),
				Name:     os.Getenv("MYSQL_DB_NAME"),
			},
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
