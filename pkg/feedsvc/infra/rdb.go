package infra

import (
	"crud-with-cache/pkg/feedsvc/domain"

	"gorm.io/gorm"
)

func NewFeedMySQLRepository(db *gorm.DB) domain.FeedRepository {
	return &mysqlRepo{db: db}
}

type mysqlRepo struct {
	db *gorm.DB
}

type Feed struct {
	ID      uint16 `gorm:"primaryKey"`
	UserID  uint16
	Title   string
	Content string
	ImgURL  []string
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

func (r *mysqlRepo) FindAllByUserID(id uint16) ([]domain.Feed, error) {
	var records []Feed
	if err := r.db.Where("user_id = ?", id).Find(&records).Error; err != nil {
		return nil, err
	}
	return toEntities(records), nil
}
