package unit

import (
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/store/repository"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cfg := &models_config.Config{
		Db: models_config.DatabaseConfig{
			DbType: "postgresql",
			Pgsql: models_config.PostgresConfig{
				Host:     "127.0.0.1",
				Port:     "5432",
				User:     "postgres",
				Password: "admin",
				DbName:   "cli_db",
				Schema:   "public",
			},
		},
		Log: models_config.LogConfig{
			Level:   "DEBUG",
			DbLevel: "DEBUG",
		},
	}
	config.Init(cfg)
	logger.Init()
	repository.Init()

	os.Exit(m.Run())
}
