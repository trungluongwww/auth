package service

import (
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
)

type CommentService interface {
	ConvertCreateCommentPayloadToModel(p request.CreateCommentPayload, userID uint32) *model.Comment
	ConvertUpdateCommentPayloadToModel(p request.UpdateCommentPayload, comment *model.Comment) *model.Comment
	ConvertToCommentResponse(comment *model.Comment, user *model.User, isLiked bool) *response.CommentResponse
}

type commentService struct {
}

func NewCommentService() CommentService {
	return &commentService{}
}

func (commentService) ConvertCreateCommentPayloadToModel(p request.CreateCommentPayload, userID uint32) *model.Comment {
	return &model.Comment{
		PostID:    p.PostID,
		UserID:    userID,
		ParentID:  p.ParentID,
		Content:   p.Content,
		LikeCount: 0,
	}
}

func (commentService) ConvertUpdateCommentPayloadToModel(p request.UpdateCommentPayload, comment *model.Comment) *model.Comment {
	comment.Content = p.Content
	return comment
}

func (commentService) ConvertToCommentResponse(comment *model.Comment, user *model.User, isLiked bool) *response.CommentResponse {
	userRes := &response.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return &response.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		LikeCount: int(comment.LikeCount),
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User:      *userRes,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
		IsLiked:   isLiked,
		Replies:   []response.CommentResponse{},
	}
}
