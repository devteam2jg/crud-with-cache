package domain

type Feed struct {
	ID      uint16
	UserID  uint16
	Title   string
	Content string
	ImgURL  []string
}
