package infra

import (
	c "context"
	"time"

	"crud-with-cache/pkg/comment/domain"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	FeedID    uint16    `gorm:"column:feed_id"`
	OwnerID   uint16    `gorm:"column:owner_id"`
	Content   string    `gorm:"column:content;size:200;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Comment) TableName() string {
	return "comment"
}

func toEntities(records []Comment) []domain.Comment {
	entities := make([]domain.Comment, 0, len(records))
	for _, r := range records {
		entities = append(entities, domain.Comment{
			ID:      r.ID,
			FeedID:  r.FeedID,
			OwnerID: r.OwnerID,
			Content: r.Content,
		})
	}
	return entities
}

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) domain.CommentRepository {
	return &mysqlRepo{
		db: db,
	}
}

func NewBufferRepository(db *gorm.DB) domain.BufferRepository {
	return &mysqlRepo{
		db: db,
	}
}

func (r *mysqlRepo) FindComments(ctx c.Context, feedID uint16) ([]domain.Comment, error) {
	var records []Comment
	err := r.db.WithContext(ctx).Where("feed_id = ?", feedID).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return toEntities(records), nil
}

func (r *mysqlRepo) InsertComment(ctx c.Context, e domain.Comment) error {
	return r.db.WithContext(ctx).Create(&Comment{
		FeedID:  e.FeedID,
		OwnerID: e.OwnerID,
		Content: e.Content,
	}).Error
}

func (r *mysqlRepo) UpdateComment(ctx c.Context, e domain.Comment) error {
	return r.db.WithContext(ctx).Model(&Comment{}).Where("id = ?", e.ID).Updates(Comment{
		Content: e.Content,
	}).Error
}

func (r *mysqlRepo) DeleteComment(ctx c.Context, e domain.Comment) error {
	return r.db.WithContext(ctx).Delete(&Comment{}, e.ID).Error
}

func (r *mysqlRepo) InsertCommentWithTransAction(ctx c.Context, e []domain.Comment) error {
	tx := r.db.WithContext(ctx).Begin()
	for _, comment := range e {
		if err := tx.Create(&Comment{
			FeedID:  comment.FeedID,
			OwnerID: comment.OwnerID,
			Content: comment.Content,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
