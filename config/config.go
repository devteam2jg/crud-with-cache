package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	MySQL  MySQLConfig
	Cached CacheConfig
	Buffer BufferConfig
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     int    `env:"MYSQL_PORT"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	DBName   string `env:"MYSQL_DBNAME"`
}

type CacheConfig struct {
	Host string `env:"REDIS_CACHE_HOST" `
	Port int    `env:"REDIS_CACHE_PORT"`
}

type BufferConfig struct {
	Host string `env:"REDIS_BUFFER_HOST" `
	Port int    `env:"REDIS_BUFFER_PORT"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("./.env.local/crud"); err != nil {
		log.Println("No .env file found")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
