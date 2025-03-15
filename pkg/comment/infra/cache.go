package infra

import (
	c "context"
	"time"

	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/redis/go-redis/v9"

	"crud-with-cache/pkg/comment/domain"
)

const (
	cacheTTL = 10 * time.Second
)

type redisCache struct {
	domain.CommentRepository
	redis redis.UniversalClient
}

func NewCache(repo domain.CommentRepository, redis redis.UniversalClient) domain.CommentRepository {
	return &redisCache{
		CommentRepository: repo,
		redis:             redis,
	}
}

func (c *redisCache) FindComments(ctx c.Context, feedID uint16) ([]domain.Comment, error) {
	key := c.makeKey(feedID)
	if cached := c.fetch(ctx, key); cached != nil {
		fmt.Println("Cache hit")
		return cached, nil
	}
	comments, e := c.CommentRepository.FindComments(ctx, feedID)
	if e != nil {
		return nil, e
	}
	c.store(ctx, key, comments)
	fmt.Println("Cache miss")
	return comments, nil
}

func (c *redisCache) fetch(ctx c.Context, key string) []domain.Comment {
	data, err := c.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}
	var comments []domain.Comment
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&comments); err != nil {
		//todo log error
	}
	return comments
}

func (c *redisCache) store(ctx c.Context, key string, comments []domain.Comment) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(comments); err != nil {
		//todo log error
	}
	c.redis.Set(ctx, key, buf.Bytes(), cacheTTL)
}

func (c *redisCache) makeKey(feedID uint16) string {
	return fmt.Sprintf("feed:%d:comments", feedID)
}
