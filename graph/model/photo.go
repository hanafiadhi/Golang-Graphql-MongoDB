package model

type Photo struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    *User      `json:"user_id"`
	Comments  []*Comment `json:"comments"`
	CreatedAt *int       `json:"created_at"`
	UpdatedAt *int       `json:"updated_at"`
}
