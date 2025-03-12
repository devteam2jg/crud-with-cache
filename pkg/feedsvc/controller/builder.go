package controller

import "crud-with-cache/pkg/feedsvc/domain"

type controllerDtoBuilder interface {
	BuildGetFeedsResponse(feeds []domain.Feed) getFeedsResponse
}

type builder struct{}

func newDtoBuilder() controllerDtoBuilder {
	return &builder{}
}

func (b *builder) BuildGetFeedsResponse(feeds []domain.Feed) getFeedsResponse {
	return getFeedsResponse{
		Feeds: feeds,
	}
}
