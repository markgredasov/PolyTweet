package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (s *PostsService) CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	post.ID = uuid.NewString()

	if err := post.Validate(); err != nil {
		return nil, fmt.Errorf("validate post: %w", domain.INVALID_REQUEST)
	}

	createdPost, err := s.postsRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("create post: %w", domain.INTERNAL_ERROR)
	}

	return createdPost, nil
}
