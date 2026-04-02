package redis

import (
	"context"

	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (r *PostsCache) SavePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	// TODO
	return nil, nil
}
