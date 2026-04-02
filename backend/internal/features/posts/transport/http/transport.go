package http

import (
	"context"
	"net/http"

	"github.com/tryingmyb3st/PolyTweet/internal/core/domain"
	"github.com/tryingmyb3st/PolyTweet/internal/core/middleware"
	"github.com/tryingmyb3st/PolyTweet/internal/core/transport/server"
)

type PostsHTTPHandler struct {
	PostsService PostsService
}

type PostsService interface {
	CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error)
}

func NewPostsHandler(postsService PostsService) *PostsHTTPHandler {
	return &PostsHTTPHandler{
		PostsService: postsService,
	}
}

func (h *PostsHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:               http.MethodPost,
			URL:                  "/posts/create",
			Handler:              h.CreatePost,
			AdditionalMiddleware: []middleware.Middleware{middleware.AuthMiddleware()},
		},
	}
}
