package controller

import "crud-with-cache/pkg/comment/domain"

type GetCommentsRequest struct {
	FeedID uint16 `param:"feed_id"`
}

type GetCommentsResponse struct {
	Comments []domain.Comment `json:"comments"`
}

type PostCommentRequest struct {
	FeedID  uint16 `param:"feed_id"`
	UserID  uint16 `json:"user_id"`
	Comment string `json:"comment"`
}

type PutCommentRequest struct {
	FeedID    uint16 `param:"feed_id"`
	CommentID uint   `param:"comment_id"`
	UserID    uint16 `json:"user_id"`
	Comment   string `json:"comment"`
}

type DeleteCommentRequest struct {
	FeedID    uint16 `param:"feed_id"`
	CommentID uint   `param:"comment_id"`
	UserID    uint16 `json:"user_id"`
}
