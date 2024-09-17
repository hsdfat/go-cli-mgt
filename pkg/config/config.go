package config

import "go-cli-mgt/pkg/models/models_config"

var config *models_config.Config


func Init(cfg *models_config.Config) {
	config = cfg
}

func Get() *models_config.Config {
    return config
}

func GetServerConfig() models_config.ServerConfig {
	return config.Svr
}

func GetDatabaseConfig() models_config.DatabaseConfig {
    return config.Db
}

func GetLogConfig() models_config.LogConfig{
	return config.Log
} 

func GetJwtConfig() models_config.TokenConfig{
    return config.Token
}

func GetRouterConfig() models_config.RouterConfig{
	return config.Router
}