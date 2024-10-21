package mysql

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	Db  *gorm.DB
	cfg models_config.MySqlConfig
}

var (
	client *Client
)

func GetInstance() *Client {
	if client == nil {
		client = &Client{}
	}
	return client
}

func (c *Client) Init(cfg models_config.DatabaseConfig) error {
	var err error
	var (
		DbUsername = cfg.Mysql.User
		DbPassword = cfg.Mysql.Password
		DbHost     = cfg.Mysql.Host
		DbPort     = cfg.Mysql.Port
		DbName     = cfg.Mysql.Name
	)
	gormLogger := logger.NewGormLogger()
	gormLogger.LogMode(1)
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logger.Logger.Debugf("Error connecting to database : error=%v", err)
		return err
	}
	logger.Logger.Info("Connect to database: ", dsn)
	c.Db = db

	return nil
}

func (c *Client) Ping() error {
	sql, err := c.Db.DB()
	if err != nil {
		return err
	}
	return sql.Ping()
}
