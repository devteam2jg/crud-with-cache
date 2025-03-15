package domain

import "time"

type Comment struct {
	ID         uint   `json:"id"`
	OwnerID    uint16 `json:"owner_id"`
	FeedID     uint16 `json:"feed_id"`
	Content    string `json:"content"`
	BufferedAt time.Time
}
