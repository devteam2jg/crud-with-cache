package domain

type Comment struct {
	ID      uint
	OwnerID uint16
	FeedID  uint16
	Content string
}
