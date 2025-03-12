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
	e.POST("/api/feed", ctrl.CreateFeed)
	e.PUT("/api/feed", ctrl.UpdateFeed)
	e.DELETE("/api/feed", ctrl.DeleteFeed)

	return ctrl
}

type controller struct {
	feedUseCase domain.FeedUseCase
}

func (con *controller) GetFeeds(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (con *controller) CreateFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (con *controller) UpdateFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (con *controller) DeleteFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
