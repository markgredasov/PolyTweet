package postgres

import (
	"context"
	"fmt"

	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (r *PostsRepository) GetLastWeekPosts(ctx context.Context, offset int, limit int) ([]domain.Post, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, user_id, content, parent_id, reply_to, image_url, created_at
    FROM posts
    WHERE created_at > current_timestamp - interval '1 week'
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`

	var posts []domain.Post

	rows, err := r.ConnPool.Query(ctxTimeout, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get recent posts: %w", err)
	}

	for rows.Next() {
		var post domain.Post
		err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.ParentID,
			&post.ReplyTo,
			&post.ImageURL,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("get recent posts: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}
