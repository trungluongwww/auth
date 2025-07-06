package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CreatePostPayload struct {
	Title    string `json:"title" validate:"required,min=1,max=255"`
	Content  string `json:"content" validate:"required,min=1"`
	ImageURL string `json:"imageUrl" validate:"omitempty,url"`
	IsPublic bool   `json:"isPublic"`
}

func (i *CreatePostPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *CreatePostPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type UpdatePostPayload struct {
	Title    string `json:"title" validate:"omitempty,min=1,max=255"`
	Content  string `json:"content" validate:"omitempty,min=1"`
	ImageURL string `json:"imageUrl" validate:"omitempty,url"`
	IsPublic *bool  `json:"isPublic"`
}

func (i *UpdatePostPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *UpdatePostPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type CreateCommentPayload struct {
	PostID   uint32  `json:"postId" validate:"required"`
	ParentID *uint32 `json:"parentId"`
	Content  string  `json:"content" validate:"required,min=1"`
}

func (i *CreateCommentPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *CreateCommentPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type UpdateCommentPayload struct {
	Content string `json:"content" validate:"required,min=1"`
}

func (i *UpdateCommentPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *UpdateCommentPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type LikePostPayload struct {
	PostID uint32 `json:"postId" validate:"required"`
}

func (i *LikePostPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *LikePostPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type LikeCommentPayload struct {
	CommentID uint32 `json:"commentId" validate:"required"`
}

func (i *LikeCommentPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *LikeCommentPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}

type FollowUserPayload struct {
	UserID uint32 `json:"userId" validate:"required"`
}

func (i *FollowUserPayload) Validate() error {
	return validator.New().Struct(i)
}

func (i *FollowUserPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}
	return i.Validate()
}
