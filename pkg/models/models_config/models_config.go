package models_config

type Config struct {
	Db     DatabaseConfig
	Svr    ServerConfig
	Log    LogConfig
	Token  TokenConfig
	Router RouterConfig
}

type ServerConfig struct {
	ServerName string
	Host       string
	Port       string
	TcpPort    string
}

type RouterConfig struct {
	BasePath string
	Origins  string
	Methods  string
	Headers  string
}

type LogConfig struct {
	Level   string
	DbLevel string
}

type DatabaseConfig struct {
	DbType string
	Mysql  MySqlConfig
	Pgsql  PostgresConfig
}

type TokenConfig struct{}

var DatabaseConfigInit Config
