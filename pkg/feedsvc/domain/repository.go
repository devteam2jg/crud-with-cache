package domain

import "context"

type FeedRepository interface {
	FindAllByUserID(ctx context.Context, userID uint16) ([]Feed, error)
}
