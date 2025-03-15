package domain

import (
	"context"
)

type CommentUseCase interface {
	GetComments(ctx context.Context, feedID uint16) ([]Comment, error)
	PostComment(ctx context.Context, dto PostCommentDto) error
	UpdateComment(ctx context.Context, dto UpdatedCommentDto) error
	DeleteComment(ctx context.Context, dto DeleteCommentDto) error
}

type commentUseCase struct {
	repo CommentRepository
}

func NewCommentUseCase(repo CommentRepository) CommentUseCase {
	return &commentUseCase{
		repo: repo,
	}
}

func (uc *commentUseCase) GetComments(ctx context.Context, feedID uint16) ([]Comment, error) {
	return uc.repo.FindComments(ctx, feedID)
}

func (uc *commentUseCase) PostComment(ctx context.Context, dto PostCommentDto) error {
	return uc.repo.InsertComment(ctx, Comment{
		OwnerID: dto.UserID,
		FeedID:  dto.FeedID,
		Content: dto.Comment,
	})
}

func (uc *commentUseCase) UpdateComment(ctx context.Context, dto UpdatedCommentDto) error {
	return uc.repo.UpdateComment(ctx, Comment{
		ID:      dto.CommentID,
		OwnerID: dto.UserID,
		FeedID:  dto.FeedID,
	})
}

func (uc *commentUseCase) DeleteComment(ctx context.Context, dto DeleteCommentDto) error {
	return uc.repo.DeleteComment(ctx, Comment{
		ID:      dto.CommentID,
		OwnerID: dto.UserID,
		FeedID:  dto.FeedID,
	})
}
