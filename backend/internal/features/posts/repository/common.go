package repository

import "time"

type PostModel struct {
	ID        string    `redis:"id"`
	UserID    string    `redis:"user_id"`
	Content   string    `redis:"content"`
	ParentID  *string   `redis:"parent_id"`
	ReplyTo   *string   `redis:"reply_to"`
	ImageURL  string    `redis:"image_url"`
	CreatedAt time.Time `redis:"created_at"`
}
