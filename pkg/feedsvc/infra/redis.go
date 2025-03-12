package infra

import (
	"bytes"
	"context"
	"crud-with-cache/pkg/feedsvc/domain"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type cache struct {
	domain.FeedRepository
	redis redis.UniversalClient
}

func NewFeedCache(repo domain.FeedRepository, redis redis.UniversalClient) domain.FeedRepository {
	return &cache{
		FeedRepository: repo,
		redis:          redis,
	}
}

func (c *cache) FindAllByUserID(ctx context.Context, userID uint16) ([]domain.Feed, error) {
	key := c.makeKey(userID)
	if cached := c.fetch(ctx, key); cached != nil {
		return cached, nil
	}
	feeds, e := c.FeedRepository.FindAllByUserID(ctx, userID)
	if e != nil {
		return nil, e
	}
	c.store(ctx, key, feeds)
	return feeds, nil
}

func (c *cache) fetch(ctx context.Context, key string) []domain.Feed {
	b, e := c.redis.Get(ctx, key).Bytes()
	if errors.Is(e, redis.Nil) {
		return nil
	}
	if e != nil {
		//todo log
		return nil
	}
	var feeds []domain.Feed
	if e := gob.NewDecoder(bytes.NewReader(b)).Decode(&feeds); e != nil {
		//todo log
		return nil
	}
	return feeds
}

func (c *cache) store(ctx context.Context, key string, feeds []domain.Feed) {
	var buf bytes.Buffer
	if e := gob.NewEncoder(&buf).Encode(feeds); e != nil {
		//todo log
	}
	if e := c.redis.SetNX(ctx, key, buf.String(), 0).Err(); e != nil {
		//todo log
	}
}

func (c *cache) makeKey(userID uint16) string {
	return fmt.Sprintf("feed:userid:%d", userID)
}
