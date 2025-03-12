package domain

type FeedUseCase interface {
}

type feedUseCase struct {
}

func NewFeedUseCase(repo FeedRepository) FeedUseCase {
	return &feedUseCase{}
}
