package repository

import (
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/models/models_db"
	"go-cli-mgt/pkg/store/postgres"
)

type DatabaseStore interface {
	Init(cfg models_config.DatabaseConfig) error

	// UserRepository
	UserRepository
	LoginRepository
	HistoryRepository
	RoleRepository
	NetworkElementRepository
	MmeCommandRepository
}

type UserRepository interface {
	SaveHistory(*models_api.History) error
	GetHistoryById(uint64) (*models_api.History, error)
	DeleteHistoryById(uint64) error
	GetHistoryListByMode(string) ([]models_api.History, error)

	CreateUser(*models_api.User) error
	GetUserByID(uint) (*models_api.User, error)
	ListUsers() ([]models_api.User, error)
	GetUserByUsername(string) (*models_api.User, error)
	DeleteUser(string) error
	UpdateUser(*models_api.User) error

	GetRoleByUserId(uint) ([]models_db.Role, error)

	CreateNetworkElement(*models_api.NeData) error
	DeleteNetworkElementByName(string, string) error
	GetNetworkElementByName(string, string) (*models_api.NeData, error)
	GetListNetworkElement() ([]models_api.NeData, error)
	GetNetworkElementByUserName(string) ([]models_api.NeData, error)
}

type RoleRepository interface {
}

type NetworkElementRepository interface {
}

type MmeCommandRepository interface {
}

type LoginRepository interface {
}

type HistoryRepository interface {
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
