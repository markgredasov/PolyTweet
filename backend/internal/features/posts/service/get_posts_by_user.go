package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
)

const (
	defaultPage     = 1
	defaultPageSize = 15
	maxPageSize     = 30
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

	page, err := parsePage(pageStr)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid pagination: %w", err)
	}

	pageSize, err := parsePageSize(pageSizeStr)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid pagination: %w", err)
	}

	if page <= 0 || pageSize <= 0 {
		return nil, nil, fmt.Errorf("page and pageSize must be positive integers: %w", domain.INVALID_REQUEST)
	}

	if pageSize > 30 {
		pageSize = maxPageSize
	}

	total, err := s.postsRepo.GetCountOfPostsByUser(ctx, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	offset := (page - 1) * pageSize
	totalPages := int64(math.Ceil(float64(total) / float64(pageSize)))

	posts, err := s.postsRepo.GetPostsByUser(ctx, userID, offset, pageSize)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Post{}, nil, nil
		}
		return nil, nil, fmt.Errorf("%s: %w", err, domain.INTERNAL_ERROR)
	}

	pagination := domain.Pagination{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	return posts, &pagination, nil
}

func parsePage(pageStr string) (int, error) {
	if pageStr == "" {
		return defaultPage, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, fmt.Errorf("page must be a number: %w", domain.INVALID_REQUEST)
	}

	if page < 1 {
		return 0, fmt.Errorf("page must be at least 1: %w", domain.INVALID_REQUEST)
	}

	return page, nil
}

func parsePageSize(pageSizeStr string) (int, error) {
	if pageSizeStr == "" {
		return defaultPageSize, nil
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, fmt.Errorf("page_size must be a number: %w", domain.INVALID_REQUEST)
	}

	if pageSize < 1 {
		return 0, fmt.Errorf("page_size must be at least 1: %w", domain.INVALID_REQUEST)
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	return pageSize, nil
}
