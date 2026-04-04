package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5"
	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

func (s *PostsService) GetLastWeekPosts(
	ctx context.Context,
	pageStr,
	pageSizeStr string,
) ([]domain.Post, *domain.Pagination, error) {
	paginationParams, err := domain.ParsePaginationParams(pageStr, pageSizeStr)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid pagination params: %w", err)
	}

	total, err := s.postsRepo.GetCountOfLastWeekPosts(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	offset := (paginationParams.Page - 1) * paginationParams.PageSize
	totalPages := int64(math.Ceil(float64(total) / float64(paginationParams.PageSize)))

	posts, err := s.postsRepo.GetLastWeekPosts(ctx, offset, paginationParams.PageSize)
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
