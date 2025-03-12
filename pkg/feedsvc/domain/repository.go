package domain

import "context"

type FeedRepository interface {
	FindOneByID(ctx context.Context, feedID uint16) (*Feed, error)
	FindAllByUserID(ctx context.Context, userID uint16) ([]Feed, error)
	Insert(ctx context.Context, feed Feed) error
	Update(ctx context.Context, feed Feed) error
	Delete(ctx context.Context, feedID uint16) error
}

type FeedRepositoryFindOption struct {
	UserID uint16
}
