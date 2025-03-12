package app

import (
	"context"
	"crud-with-cache/config"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Infra struct {
	RDB   *gorm.DB
	Redis redis.UniversalClient
}

func NewInfra(cfg *config.Config) (*Infra, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DBName)

	// Initialize the GORM DB object
	rdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to connect to mysql"), err)
	}

	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to connect to redis"), err)
	}

	return &Infra{
		RDB:   rdb,
		Redis: redisClient,
	}, nil
}
