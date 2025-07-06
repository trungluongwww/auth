package service

import (
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
)

type PostService interface {
	ConvertCreatePostPayloadToModel(p request.CreatePostPayload, userID uint32) *model.Post
	ConvertUpdatePostPayloadToModel(p request.UpdatePostPayload, post *model.Post) *model.Post
	ConvertToPostResponse(post *model.Post, user *model.User, isLiked bool) *response.PostResponse
}

type postService struct {
}

func NewPostService() PostService {
	return &postService{}
}

func (postService) ConvertCreatePostPayloadToModel(p request.CreatePostPayload, userID uint32) *model.Post {
	return &model.Post{
		UserID:       userID,
		Title:        p.Title,
		Content:      p.Content,
		ImageURL:     &p.ImageURL,
		IsPublic:     p.IsPublic,
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
	}
}

func (postService) ConvertUpdatePostPayloadToModel(p request.UpdatePostPayload, post *model.Post) *model.Post {
	if p.Title != "" {
		post.Title = p.Title
	}
	if p.Content != "" {
		post.Content = p.Content
	}
	if p.ImageURL != "" {
		post.ImageURL = &p.ImageURL
	}
	if p.IsPublic != nil {
		post.IsPublic = *p.IsPublic
	}
	return post
}

func (postService) ConvertToPostResponse(post *model.Post, user *model.User, isLiked bool) *response.PostResponse {
	userRes := &response.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return &response.PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		Content:      post.Content,
		ImageURL:     post.ImageURL,
		IsPublic:     post.IsPublic,
		ViewCount:    int(post.ViewCount),
		LikeCount:    int(post.LikeCount),
		CommentCount: int(post.CommentCount),
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		User:         *userRes,
		IsLiked:      isLiked,
	}
}
