package domain

type PostCommentDto struct {
	UserID  uint16 `json:"user_id"`
	FeedID  uint16 `json:"feed_id"`
	Content string `json:"comment"`
}

type UpdatedCommentDto struct {
	CommentID uint   `json:"comment_id"`
	UserID    uint16 `json:"user_id"`
	FeedID    uint16 `json:"feed_id"`
	Content   string `json:"comment"`
}

type DeleteCommentDto struct {
	CommentID uint   `json:"comment_id"`
	UserID    uint16 `json:"user_id"`
	FeedID    uint16 `json:"feed_id"`
}
