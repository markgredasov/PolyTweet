package domain

import (
	"fmt"
	"strconv"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 15
	MaxPageSize     = 30
)

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func ParsePaginationParams(pageStr, pageSizeStr string) (*PaginationParams, error) {
	page, err := parsePage(pageStr)
	if err != nil {
		return nil, err
	}

	pageSize, err := parsePageSize(pageSizeStr)
	if err != nil {
		return nil, err
	}

	return &PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func parsePage(pageStr string) (int, error) {
	if pageStr == "" {
		return DefaultPage, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, fmt.Errorf("page must be a number: %w", INVALID_REQUEST)
	}

	if page < 1 {
		return 0, fmt.Errorf("page must be at least 1: %w", INVALID_REQUEST)
	}

	return page, nil
}

func parsePageSize(pageSizeStr string) (int, error) {
	if pageSizeStr == "" {
		return DefaultPageSize, nil
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, fmt.Errorf("page_size must be a number: %w", INVALID_REQUEST)
	}

	if pageSize < 1 {
		return 0, fmt.Errorf("page_size must be at least 1: %w", INVALID_REQUEST)
	}

	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	return pageSize, nil
}
