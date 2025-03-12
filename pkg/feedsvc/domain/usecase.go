package domain

import "context"

type FeedUseCase interface {
	GetFeeds(ctx context.Context, userID uint16) ([]Feed, error)
	CreateFeed(ctx context.Context, feed Feed) error
	UpdateFeed(ctx context.Context, feed Feed) error
	DeleteFeed(ctx context.Context, feedID uint16) error
}

type feedUseCase struct {
	repo FeedRepository
}

func NewFeedUseCase(repo FeedRepository) FeedUseCase {
	return &feedUseCase{
		repo: repo,
	}
}

func (uc *feedUseCase) GetFeeds(ctx context.Context, userID uint16) ([]Feed, error) {
	return uc.repo.FindAll(ctx, userID)
}

func (uc *feedUseCase) CreateFeed(ctx context.Context, feed Feed) error {
	return uc.repo.Insert(ctx, feed)
}

func (uc *feedUseCase) UpdateFeed(ctx context.Context, feed Feed) error {
	return uc.repo.Update(ctx, feed)
}

func (uc *feedUseCase) DeleteFeed(ctx context.Context, feedID uint16) error {
	return uc.repo.Delete(ctx, feedID)
}
