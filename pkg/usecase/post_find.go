package usecase

import (
	"context"
	"errors"

	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/query"
	"github.com/trungluongwww/auth/pkg/model/response"
)

// GetPost implementation for find operations
func (u *postImpl) GetPost(context context.Context, userID uint32, postID uint32) (*response.PostResponse, error) {
	post, err := u.Repository.NewPost().FirstRaw(&model.Post{ID: postID})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("post not found")
	}

	// Increment view count
	//u.Repository.NewPost().IncrementViewCount(postID)

	user, err := u.Repository.NewUser().FirstRaw(&model.User{ID: post.UserID})
	if err != nil {
		return nil, err
	}

	// Check if user liked the post
	isLiked, _ := u.Repository.NewPostLike().Exists(postID, userID)

	return u.PostService.ConvertToPostResponse(post, user, isLiked), nil
}

// GetPosts implementation for find operations
func (u *postImpl) GetPosts(context context.Context, userID uint32, page, limit int, search string) (*response.PostListResponse, error) {
	condition := query.PostCondition{
		CommonCondition: query.CommonCondition{
			Pagination: &query.Pagination{
				Page:  page,
				Limit: limit,
			},
			Order: &query.Order{
				OrderBy:    "created_at",
				OrderValue: "desc",
			},
		},
		IsPublic: &[]bool{true}[0], // Only public posts
		Search:   search,
	}

	posts, err := u.Repository.NewPost().FindByCondition(condition, userID)
	if err != nil {
		return nil, err
	}

	// Convert to response
	postResponses := make([]response.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = *u.PostService.ConvertToPostResponse(&post.Post, post.User, false)
	}

	return &response.PostListResponse{
		Posts:      postResponses,
		TotalCount: len(postResponses),
		Page:       page,
		Limit:      limit,
	}, nil
}

// GetUserPosts implementation for find operations
func (u *postImpl) GetUserPosts(context context.Context, currentUserID, targetUserID uint32, page, limit int) (*response.PostListResponse, error) {
	condition := query.PostCondition{
		CommonCondition: query.CommonCondition{
			Pagination: &query.Pagination{
				Page:  page,
				Limit: limit,
			},
			Order: &query.Order{
				OrderBy:    "created_at",
				OrderValue: "desc",
			},
		},
		UserID: targetUserID,
	}

	// If viewing own posts, show all. If viewing others, show only public
	if currentUserID != targetUserID {
		isPublic := true
		condition.IsPublic = &isPublic
	}

	posts, err := u.Repository.NewPost().FindByCondition(condition, currentUserID)
	if err != nil {
		return nil, err
	}

	// Convert to response
	postResponses := make([]response.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = *u.PostService.ConvertToPostResponse(&post.Post, post.User, false)
	}

	return &response.PostListResponse{
		Posts:      postResponses,
		TotalCount: len(postResponses),
		Page:       page,
		Limit:      limit,
	}, nil
}
