package domain

type Feed struct {
	ID      uint16   `json:"id"`
	UserID  uint16   `json:"user_id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	ImgURL  []string `json:"img_url"`
}
