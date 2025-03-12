package controller

import "crud-with-cache/pkg/feedsvc/domain"

type feedDtoBase struct {
	userID uint16 `param:"user_id"`
}

type getFeedsResponse struct {
	Feeds []domain.Feed `json:"feeds"`
}
