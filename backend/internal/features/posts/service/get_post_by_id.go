package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (s *PostsService) GetPostByID(ctx context.Context, postID string) (*domain.Post, error) {
	if _, err := uuid.Parse(postID); err != nil {
		return nil, fmt.Errorf("post id is not uuid: %w", domain.INVALID_REQUEST)
	}

	post, err := s.postsRepo.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no post: %w", domain.NOT_FOUND)
		}
		return nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	return post, nil
}
