package postgres

import (
	postgres_pool "github.com/tryingmyb3st/PolyTweet/internal/core/repository/postgres"
)

type PostsRepository struct {
	ConnPool postgres_pool.Pool
}

func NewPostsRepository(connPool postgres_pool.Pool) *PostsRepository {
	return &PostsRepository{
		ConnPool: connPool,
	}
}
