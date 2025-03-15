package controller

import (
	"crud-with-cache/pkg/comment/domain"
	"crud-with-cache/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController interface {
	GetComments(c echo.Context) error
	PostComment(c echo.Context) error
	PutComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}

type commentController struct {
	useCase domain.CommentUseCase
	binder  echo.Binder
}

func NewCommentController(e router.Router, useCase domain.CommentUseCase) CommentController {
	ctrl := &commentController{
		useCase: useCase,
		binder:  &echo.DefaultBinder{},
	}

	e.GET("/api/feed/:feed_id/comments", ctrl.GetComments)
	e.POST("/api/feed/:feed_id/comment", ctrl.PostComment)
	e.PUT("/api/feed/:feed_id/comment/:comment_id", ctrl.PutComment)
	e.DELETE("/api/feed/:feed_id/comment/:comment_id", ctrl.DeleteComment)

	return ctrl
}

func (con *commentController) GetComments(c echo.Context) error {
	ctx := c.Request().Context()
	req := &GetCommentsRequest{}
	if err := con.binder.Bind(req, c); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if req.FeedID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: feed_id"}
	}
	comments, err := con.useCase.GetComments(ctx, req.FeedID)
	//todo make response code more specific
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError,
			Message: "Unexpected error: failed to get comments"}
	}
	return c.JSON(http.StatusOK, &GetCommentsResponse{Comments: comments})
}

func (con *commentController) PostComment(c echo.Context) error {
	ctx := c.Request().Context()
	req := &PostCommentRequest{}
	if err := con.binder.Bind(req, c); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if req.FeedID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: feed_id"}
	}
	if req.UserID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing json: user_id"}
	}
	if req.Comment == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing json: comment"}
	}
	err := con.useCase.PostComment(ctx, domain.PostCommentDto{
		FeedID:  req.FeedID,
		UserID:  req.UserID,
		Comment: req.Comment,
	})
	//todo make response code more specific
	if err != nil {
		//todo log error
		return &echo.HTTPError{Code: http.StatusInternalServerError,
			Message: "Unexpected error: failed to post comment"}
	}
	return c.NoContent(http.StatusNoContent)
}

func (con *commentController) PutComment(c echo.Context) error {
	ctx := c.Request().Context()
	req := &PutCommentRequest{}
	if err := con.binder.Bind(req, c); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if req.FeedID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: feed_id"}
	}
	if req.CommentID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: comment_id"}
	}
	if req.UserID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing json: user_id"}
	}
	if req.Comment == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing json: comment"}
	}
	err := con.useCase.UpdateComment(ctx, domain.UpdatedCommentDto{
		FeedID:    req.FeedID,
		CommentID: req.CommentID,
		UserID:    req.UserID,
		Comment:   req.Comment,
	})
	//todo make response code more specific
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError,
			Message: "Unexpected error: failed to update comment"}
	}
	return c.NoContent(http.StatusNoContent)
}

func (con *commentController) DeleteComment(c echo.Context) error {
	ctx := c.Request().Context()
	req := &DeleteCommentRequest{}
	if err := con.binder.Bind(req, c); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}
	}
	if req.FeedID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: feed_id"}
	}
	if req.CommentID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing param: comment_id"}
	}
	if req.UserID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Missing json: user_id"}
	}
	err := con.useCase.DeleteComment(ctx, domain.DeleteCommentDto{
		FeedID:    req.FeedID,
		CommentID: req.CommentID,
		UserID:    req.UserID,
	})
	//todo make response code more specific
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError,
			Message: "Unexpected error: failed to delete comment"}
	}
	return c.NoContent(http.StatusNoContent)
}
