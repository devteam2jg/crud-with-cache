package domain

type FeedRepository interface {
	FindAllByUserID(id uint16) ([]Feed, error)
}
