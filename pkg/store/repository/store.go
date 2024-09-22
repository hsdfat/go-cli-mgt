package repository

import (
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/store/postgres"
)

type DatabaseStore interface {
	Init(cfg models_config.DatabaseConfig) error

	// UserRepository
	CreateUserRepository
}

type CreateUserRepository interface {
	CreateUser(user *models_api.User) error
	GetUserByID(id uint) (*models_api.User, error)
	ListUsers() ([]models_api.User, error)
}

var (
	store DatabaseStore
)

func GetSingleton() DatabaseStore {
	return store
}

func Init() {
	cfg := config.GetDatabaseConfig()
	switch cfg.DbType {
	case "mysql":
		// store = mysql.GetInstance()
	case "postgresql":
		store = postgres.GetInstance()
	default:
		panic("unsupported database type")
	}
	err := store.Init(cfg)
	if err != nil {
		panic("cant init store")
	}
}
