package repository

import (
	"github.com/hsdfat/go-cli-mgt/pkg/config"
	"github.com/hsdfat/go-cli-mgt/pkg/store/postgres"
)

var (
	store IDatabaseStore
)

func GetSingleton() IDatabaseStore {
	return store
}

func Init() {
	var err error
	cfg := config.GetDatabaseConfig()
	switch cfg.DbType {
	case "mysql":
		// store, err = mysql.NewClient(cfg)
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
