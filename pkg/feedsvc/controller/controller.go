package controller

import "crud-with-cache/pkg/feedsvc/domain"

type FeedController interface {
}

func NewFeedController(repo domain.FeedRepository) FeedController {
	return &controller{
		repo: repo,
	}
}

type controller struct {
	repo domain.FeedRepository
}
