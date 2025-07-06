package usecase

import (
	"context"

	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
	"github.com/trungluongwww/auth/pkg/repository"
	"github.com/trungluongwww/auth/pkg/service"
)

// Post interface contains all post-related operations
type Post interface {
	// Find operations
	GetPost(context context.Context, userID uint32, postID uint32) (*response.PostResponse, error)
	GetPosts(context context.Context, userID uint32, page, limit int, search string) (*response.PostListResponse, error)
	GetUserPosts(context context.Context, currentUserID, targetUserID uint32, page, limit int) (*response.PostListResponse, error)

	// Upsert operations
	CreatePost(context context.Context, userID uint32, p request.CreatePostPayload) (*response.PostResponse, error)
	UpdatePost(context context.Context, userID uint32, postID uint32, p request.UpdatePostPayload) (*response.PostResponse, error)
	DeletePost(context context.Context, userID uint32, postID uint32) error
	LikePost(context context.Context, userID uint32, p request.LikePostPayload) error
	UnlikePost(context context.Context, userID uint32, p request.LikePostPayload) error
}

type postImpl struct {
	Repository  repository.Repository
	PostService service.PostService
	UserService service.UserService
}

// NewPost creates a new Post use case
func NewPost(repository repository.Repository, postService service.PostService, userService service.UserService) Post {
	return &postImpl{
		Repository:  repository,
		PostService: postService,
		UserService: userService,
	}
}
