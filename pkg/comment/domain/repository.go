package domain

import (
	c "context"
)

type CommentRepository interface {
	FindComments(ctx c.Context, feedID uint16) ([]Comment, error)
	InsertComment(ctx c.Context, e Comment) error
	UpdateComment(ctx c.Context, e Comment) error
	DeleteComment(ctx c.Context, e Comment) error
}
