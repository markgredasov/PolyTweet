package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	ParentID  *string   `json:"parent_id"`
	ReplyTo   *string   `json:"reply_to"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) Validate() error {
	validator := validator.New()

	if err := validator.Struct(p); err != nil {
		return fmt.Errorf("validator struct: %w", err)
	}

	return nil
}
