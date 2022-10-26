package model

type Comment struct {
	ID        string `json:"id"`
	UserID    *User  `json:"user_id"`
	PhotoID   *Photo `json:"photo_id"`
	Message   string `json:"message"`
	CreatedAt *int   `json:"created_at"`
	UpdatedAt *int   `json:"updated_at"`
}
