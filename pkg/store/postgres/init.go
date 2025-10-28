package postgres

import (
	"fmt"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_config "github.com/hsdfat/go-cli-mgt/pkg/models/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgClient struct {
	Db  *gorm.DB
	cfg models_config.PostgresConfig
}

// NewClient initializes a new PostgreSQL client using GORM
func NewClient(cfg models_config.DatabaseConfig) (*PgClient, error) {
	logger.Logger.Debugf("Initializing PostgreSQL database connection with GORM, config: %v", cfg)

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.Pgsql.Host,
		cfg.Pgsql.Port,
		cfg.Pgsql.User,
		cfg.Pgsql.Password,
		cfg.Pgsql.DbName,
	)

	// Initialize GORM logger
	gormLogger := logger.NewGormLogger()
	gormLogger.LogMode(1)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logger.Logger.Errorf("Failed to connect to PostgreSQL database: %v", err)
		return nil, err
	}

	// Get underlying SQL DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Logger.Errorf("Failed to get underlying SQL DB: %v", err)
		return nil, err
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(25)                     // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                      // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(5 * 60 * 1000000000) // 5 minutes in nanoseconds
	sqlDB.SetConnMaxIdleTime(1 * 60 * 1000000000) // 1 minute in nanoseconds

	logger.Logger.Info("Successfully connected to PostgreSQL database using GORM: ", dsn)

	c := &PgClient{
		Db:  db,
		cfg: cfg.Pgsql,
	}

	return c, nil
}

var (
	client *PgClient
)

// GetInstance returns singleton instance of PgClient
func GetInstance() *PgClient {
	if client == nil {
		client = &PgClient{}
	}
	return client
}

// Ping checks if database connection is alive
func (c *PgClient) Ping() error {
	sqlDB, err := c.Db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
