package subscriber

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
	RDB    *gorm.DB
	Buffer redis.UniversalClient
}

func NewInfra(cfg *config.Config) (*Infra, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DBName)

	// Initialize the GORM DB object
	rdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to connect to mysql"), err)
	}

	buffer := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Buffer.Host, cfg.Buffer.Port),
	})
	_, err = buffer.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to connect to redis buffer"), err)
	}

	return &Infra{
		RDB:    rdb,
		Buffer: buffer,
	}, nil
}
