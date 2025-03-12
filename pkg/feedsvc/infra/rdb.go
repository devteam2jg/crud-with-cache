package infra

import (
	"crud-with-cache/pkg/feedsvc/domain"

	"gorm.io/gorm"
)

func NewFeedMySQLRepository(db *gorm.DB) domain.FeedRepository {
	return &mysqlRepo{db: db}
}

type mysqlRepo struct {
	db *gorm.DB
}
