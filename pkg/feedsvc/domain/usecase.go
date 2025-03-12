package domain

type FeedUseCase interface {
}

type feedUseCase struct {
	repo FeedRepository
}

func NewFeedUseCase(repo FeedRepository) FeedUseCase {
	return &feedUseCase{
		repo: repo,
	}
}
