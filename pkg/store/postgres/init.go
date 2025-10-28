package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_config "github.com/hsdfat/go-cli-mgt/pkg/models/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgClient struct {
	pool *pgxpool.Pool
	cfg  models_config.PostgresConfig
}

func NewClient(cfg models_config.DatabaseConfig) (*PgClient, error) {
	// Initialize PostgreSQL database connection
	logger.Logger.Debugf("Initializing PostgreSQL database connection with config %v", cfg)
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Pgsql.User, cfg.Pgsql.Password,
		cfg.Pgsql.Host, cfg.Pgsql.Port,
		cfg.Pgsql.DbName)
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	c := &PgClient{
		pool: dbPool,
		cfg:  cfg.Pgsql,
	}
	return c, nil
}

var (
	client *PgClient
)

func GetInstance() *PgClient {
	if client == nil {
		client = &PgClient{}
	}
	return client
}
