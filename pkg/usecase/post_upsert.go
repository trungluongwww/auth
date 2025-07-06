package usecase

import (
	"context"
	"errors"

	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
)

// CreatePost implementation for upsert operations
func (u *postImpl) CreatePost(context context.Context, userID uint32, p request.CreatePostPayload) (*response.PostResponse, error) {
	post := u.PostService.ConvertCreatePostPayloadToModel(p, userID)

	err := u.Repository.NewPost().Insert(post)
	if err != nil {
		return nil, err
	}

	user, err := u.Repository.NewUser().FirstRaw(&model.User{ID: userID})
	if err != nil {
		return nil, err
	}

	return u.PostService.ConvertToPostResponse(post, user, false), nil
}

// UpdatePost implementation for upsert operations
func (u *postImpl) UpdatePost(context context.Context, userID uint32, postID uint32, p request.UpdatePostPayload) (*response.PostResponse, error) {
	post, err := u.Repository.NewPost().FirstRaw(&model.Post{ID: postID})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("post not found")
	}
	if post.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	updatedPost := u.PostService.ConvertUpdatePostPayloadToModel(p, post)
	err = u.Repository.NewPost().Update(updatedPost)
	if err != nil {
		return nil, err
	}

	user, err := u.Repository.NewUser().FirstRaw(&model.User{ID: userID})
	if err != nil {
		return nil, err
	}

	// Check if user liked the post
	isLiked, _ := u.Repository.NewPostLike().Exists(postID, userID)

	return u.PostService.ConvertToPostResponse(updatedPost, user, isLiked), nil
}

// DeletePost implementation for upsert operations
func (u *postImpl) DeletePost(context context.Context, userID uint32, postID uint32) error {
	post, err := u.Repository.NewPost().FirstRaw(&model.Post{ID: postID})
	if err != nil {
		return err
	}
	if post == nil {
		return errors.New("post not found")
	}
	if post.UserID != userID {
		return errors.New("unauthorized")
	}

	return u.Repository.NewPost().Delete(post)
}

// LikePost implementation for upsert operations
func (u *postImpl) LikePost(context context.Context, userID uint32, p request.LikePostPayload) error {
	// Check if already liked
	exists, err := u.Repository.NewPostLike().Exists(p.PostID, userID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("post already liked")
	}

	// Create like record
	like := &model.PostLike{
		PostID: p.PostID,
		UserID: userID,
	}

	err = u.Repository.NewPostLike().Insert(like)
	if err != nil {
		return err
	}

	// Increment like count
	return u.Repository.NewPost().IncrementLikeCount(p.PostID)
}

// UnlikePost implementation for upsert operations
func (u *postImpl) UnlikePost(context context.Context, userID uint32, p request.LikePostPayload) error {
	// Check if liked
	exists, err := u.Repository.NewPostLike().Exists(p.PostID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("post not liked")
	}

	// Delete like record
	like, err := u.Repository.NewPostLike().FirstRaw(&model.PostLike{
		PostID: p.PostID,
		UserID: userID,
	})
	if err != nil {
		return err
	}

	err = u.Repository.NewPostLike().Delete(like)
	if err != nil {
		return err
	}

	// Decrement like count
	return u.Repository.NewPost().DecrementLikeCount(p.PostID)
}
