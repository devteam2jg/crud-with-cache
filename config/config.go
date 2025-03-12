package config

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     int    `env:"MYSQL_PORT"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	DBName   string `env:"MYSQL_DBNAME"`
}

type RedisConfig struct {
	Host string `env:"REDIS_HOST" `
	Port int    `env:"REDIS_PORT"`
}
