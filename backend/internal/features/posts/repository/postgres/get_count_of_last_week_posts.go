package postgres

import (
	"context"
	"fmt"
)

func (r *PostsRepository) GetCountOfLastWeekPosts(ctx context.Context) (int64, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
	SELECT COUNT(*)
    FROM posts
    WHERE created_at > current_timestamp - interval '1 week'
	`

	row := r.ConnPool.QueryRow(ctxTimeout, query)

	var count int64
	err := row.Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("get count: %w", err)
	}

	return count, nil
}
