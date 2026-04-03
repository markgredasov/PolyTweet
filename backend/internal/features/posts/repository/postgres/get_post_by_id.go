package postgres

import (
	"context"
	"fmt"

	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (r *PostsRepository) GetPostByID(ctx context.Context, postID string) (*domain.Post, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, user_id, content, parent_id, reply_to, image_url, created_at
	FROM posts
	WHERE id = $1;
	`

	row := r.ConnPool.QueryRow(ctxTimeout, query, postID)

	var post domain.Post
	err := row.Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
		&post.ParentID,
		&post.ReplyTo,
		&post.ImageURL,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("scan returning post: %w", err)
	}

	return &domain.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Content:   post.Content,
		ParentID:  post.ParentID,
		ReplyTo:   post.ReplyTo,
		ImageURL:  post.ImageURL,
		CreatedAt: post.CreatedAt,
	}, nil
}
