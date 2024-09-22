package main

import (
	"fmt"
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/server"
	"go-cli-mgt/pkg/store/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	httpServer := Initialize()

	go server.ListenAndServe(httpServer)

	stopOrKillServer()
}

func stopOrKillServer() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, os.Interrupt)
	sig := <-signals
	fmt.Println("Receive Signal from OS - Release resource")
	fmt.Println(sig)
	os.Exit(1)
}

func Initialize() *fiber.App {
	err := godotenv.Load()
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
	return server.NewFiber()
}
