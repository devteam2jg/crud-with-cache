package controller

import (
	"crud-with-cache/pkg/feedsvc/domain"
	"crud-with-cache/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FeedController interface {
	GetFeeds(c echo.Context) error
}

func NewFeedController(e router.Router, feedUseCase domain.FeedUseCase) FeedController {
	ctrl := &controller{
		feedUseCase: feedUseCase,
	}
	e.GET("/api/feeds", ctrl.GetFeeds)
	return ctrl
}

type controller struct {
	feedUseCase domain.FeedUseCase
}

func (con *controller) GetFeeds(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
