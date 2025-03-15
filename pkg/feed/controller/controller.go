package controller

import (
	"crud-with-cache/pkg/feed/domain"
	"crud-with-cache/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FeedController interface {
	GetFeeds(c echo.Context) error
	CreateFeed(c echo.Context) error
	UpdateFeed(c echo.Context) error
	DeleteFeed(c echo.Context) error
}

func NewFeedController(e router.Router, feedUseCase domain.FeedUseCase) FeedController {
	ctrl := &controller{
		feedUseCase: feedUseCase,
		dtoBuilder:  newDtoBuilder(),
	}

	e.GET("/api/user/:user_id/feeds", ctrl.GetFeeds)
	e.POST("/api/user/:user_id/feed", ctrl.CreateFeed)
	e.PUT("/api/user/:user_id/feed", ctrl.UpdateFeed)
	e.DELETE("/api/user/:user_id/feed", ctrl.DeleteFeed)

	e.GET("/api/feed/test", ctrl.testFeed)

	return ctrl
}

type controller struct {
	feedUseCase domain.FeedUseCase
	dtoBuilder  controllerDtoBuilder
}

func (con *controller) GetFeeds(c echo.Context) error {
	ctx := c.Request().Context()
	binder := echo.DefaultBinder{}
	req := &feedDtoBase{}
	if err := binder.BindPathParams(c, req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	feeds, err := con.feedUseCase.GetFeeds(ctx, req.userID)
	if err != nil {
		//todo: log error
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unexpected error: failed to get feeds"}
	}
	return c.JSON(http.StatusOK, con.dtoBuilder.BuildGetFeedsResponse(feeds))
}

func (con *controller) CreateFeed(c echo.Context) error {
	ctx := c.Request().Context()
	binder := echo.DefaultBinder{}
	req := &feedDtoBase{}
	if err := binder.BindPathParams(c, req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if err := con.feedUseCase.CreateFeed(ctx, domain.Feed{UserID: req.userID}); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unexpected error: failed to create feed"}
	}
	return c.NoContent(http.StatusNoContent)
}

func (con *controller) UpdateFeed(c echo.Context) error {
	ctx := c.Request().Context()
	binder := echo.DefaultBinder{}
	req := &feedDtoBase{}
	if err := binder.BindPathParams(c, req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if err := con.feedUseCase.UpdateFeed(ctx, domain.Feed{UserID: req.userID}); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unexpected error: failed to update feed"}
	}
	return c.NoContent(http.StatusNoContent)
}

func (con *controller) DeleteFeed(c echo.Context) error {
	ctx := c.Request().Context()
	binder := echo.DefaultBinder{}
	req := &feedDtoBase{}
	if err := binder.BindPathParams(c, req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if err := con.feedUseCase.DeleteFeed(ctx, req.userID); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unexpected error: failed to delete feed"}
	}
	return c.NoContent(http.StatusNoContent)
}

func (con *controller) testFeed(c echo.Context) error {
	ctx := c.Request().Context()
	if err := con.feedUseCase.CreateFeed(ctx, domain.Feed{
		UserID:  1,
		Title:   "test",
		Content: "test",
		ImgURL:  []string{"test"},
	}); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError,
			Message: "Unexpected error: failed to post comment"}
	}
	return c.NoContent(http.StatusNoContent)
}
