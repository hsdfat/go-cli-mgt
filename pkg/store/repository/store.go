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
	var err error
	cfg := config.GetDatabaseConfig()
	switch cfg.DbType {
	case "mysql":
		store, err = mysql.NewClient(cfg)
	case "postgresql":
		store, err = postgres.NewClient(cfg)
	//case "aerospike":
	//	store = aerospikes.GetInstance()
	default:
		panic("unsupported database type")
	}
	if err != nil {
		panic("cant init store")
	}
}
