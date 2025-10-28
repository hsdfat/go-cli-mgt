package mysql

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_config "github.com/hsdfat/go-cli-mgt/pkg/models/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	Db  *gorm.DB
	cfg models_config.MySqlConfig
}

func NewClient(cfg models_config.DatabaseConfig) (*Client, error) {
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
		return nil, err
	}
	logger.Logger.Info("Connect to database: ", dsn)
	c := &Client{
		Db:  db,
		cfg: cfg.Mysql,
	}

	return c, nil
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

func (c *Client) Ping() error {
	sql, err := c.Db.DB()
	if err != nil {
		return err
	}
	return sql.Ping()
}
