package repository

import (
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/store/mysql"
	"go-cli-mgt/pkg/store/postgres"
)

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
		store = mysql.GetInstance()
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
