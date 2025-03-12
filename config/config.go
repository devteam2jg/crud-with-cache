package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

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

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
