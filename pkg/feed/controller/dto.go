package controller

import "crud-with-cache/pkg/feed/domain"

type feedDtoBase struct {
	UserID int16 `param:"user_id"`
}

type getFeedsResponse struct {
	Feeds []domain.Feed `json:"feeds"`
}
