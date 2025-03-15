package infra

import (
	"context"
	"crud-with-cache/pkg/feed/domain"
	"errors"

	"gorm.io/gorm"
)

func NewMySQLRepository(db *gorm.DB) domain.FeedRepository {
	return &mysqlRepo{db: db}
}

type mysqlRepo struct {
	db *gorm.DB
}

type Feed struct {
	ID      uint16   `gorm:"primaryKey;column:id"`
	UserID  uint16   `gorm:"column:user_id"`
	Title   string   `gorm:"column:title"`
	Content string   `gorm:"column:content"`
	ImgURL  []string `gorm:"column:img_url"`
}

func toEntities(records []Feed) []domain.Feed {
	entities := make([]domain.Feed, 0, len(records))
	for _, r := range records {
		entities = append(entities, domain.Feed{
			ID:      r.ID,
			UserID:  r.UserID,
			Title:   r.Title,
			Content: r.Content,
			ImgURL:  r.ImgURL,
		})
	}
	return entities
}

func (Feed) TableName() string {
	return "feeds"
}

func (f Feed) toEntity() *domain.Feed {
	return &domain.Feed{
		ID:      f.ID,
		UserID:  f.UserID,
		Title:   f.Title,
		Content: f.Content,
		ImgURL:  f.ImgURL,
	}
}

func (r *mysqlRepo) FindOneByID(ctx context.Context, id uint16) (*domain.Feed, error) {
	var record Feed
	if err := r.db.WithContext(ctx).First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrFeedNotFound
		}
		return nil, err
	}
	return record.toEntity(), nil
}

func (r *mysqlRepo) FindAllByUserID(ctx context.Context, userID uint16) ([]domain.Feed, error) {
	var records []Feed
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}
	return toEntities(records), nil
}

func (r *mysqlRepo) Insert(ctx context.Context, feed domain.Feed) error {
	return r.db.WithContext(ctx).Create(&Feed{
		UserID:  feed.UserID,
		Title:   feed.Title,
		Content: feed.Content,
		ImgURL:  feed.ImgURL,
	}).Error
}

func (r *mysqlRepo) Update(ctx context.Context, feed domain.Feed) error {
	return r.db.WithContext(ctx).Save(&Feed{
		ID:      feed.ID,
		UserID:  feed.UserID,
		Title:   feed.Title,
		Content: feed.Content,
		ImgURL:  feed.ImgURL,
	}).Error
}

func (r *mysqlRepo) Delete(ctx context.Context, feedID uint16) error {
	return r.db.WithContext(ctx).Delete(&Feed{}, feedID).Error
}
