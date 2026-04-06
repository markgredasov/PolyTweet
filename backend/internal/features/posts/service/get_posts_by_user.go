package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (s *PostsService) GetPostsByUser(
	ctx context.Context,
	userID string,
	pageStr string,
	pageSizeStr string,
) ([]domain.Post, *domain.Pagination, error) {
	if _, err := uuid.Parse(userID); err != nil {
		return nil, nil, fmt.Errorf("user id is not uuid: %w", domain.INVALID_REQUEST)
	}

	paginationParams, err := domain.ParsePaginationParams(pageStr, pageSizeStr)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid pagination params: %w", err)
	}

	total, err := s.postsRepo.GetCountOfPostsByUser(ctx, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	offset := (paginationParams.Page - 1) * paginationParams.PageSize
	totalPages := int64(math.Ceil(float64(total) / float64(paginationParams.PageSize)))

	posts, err := s.postsRepo.GetPostsByUser(ctx, userID, offset, paginationParams.PageSize)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Post{}, &domain.Pagination{}, nil
		}
		return nil, nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	pagination := domain.Pagination{
		Page:       paginationParams.Page,
		PageSize:   paginationParams.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	return posts, &pagination, nil
}
