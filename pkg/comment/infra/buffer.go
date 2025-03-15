package infra

import (
	"crud-with-cache/pkg/comment/domain"

	"github.com/redis/go-redis/v9"
)

type redisBuffer struct {
	domain.CommentRepository
	redis redis.UniversalClient
}

func NewBuffer(repo domain.CommentRepository, buffer redis.UniversalClient) domain.CommentRepository {
	return &redisBuffer{
		CommentRepository: repo,
		redis:             buffer,
	}
}

//func (r *redisBuffer) InsertComment(ctx c.Context, e domain.Comment) error {
//
//}
//
//func (r *redisBuffer) buffering(ctx c.Context, key string) {
//}
//
//func (r *redisBuffer) flush(ctx c.Context, key string) {
//}
