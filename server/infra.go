package server

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Infra struct {
	RDB   *gorm.DB
	Redis redis.UniversalClient
}
