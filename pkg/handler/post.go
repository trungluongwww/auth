package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/usecase"
	"github.com/trungluongwww/auth/util/custom"
)

type Post struct {
	UsecasePost usecase.Post
}

func NewPost(UsecasePost usecase.Post) *Post {
	return &Post{
		UsecasePost: UsecasePost,
	}
}

func (h *Post) CreatePost(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.CreatePostPayload{}
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	err = input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	res, err := h.UsecasePost.CreatePost(ctx.CurrentCtx(), uint32(userID), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusCreated, res)
}

func (h *Post) UpdatePost(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.UpdatePostPayload{}
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	postID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": "invalid post id"})
	}

	err = input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	res, err := h.UsecasePost.UpdatePost(ctx.CurrentCtx(), uint32(userID), uint32(postID), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *Post) DeletePost(context echo.Context) error {
	var (
		ctx = custom.NewEchoCustom(context)
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	postID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": "invalid post id"})
	}

	err = h.UsecasePost.DeletePost(ctx.CurrentCtx(), uint32(userID), uint32(postID))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, echo.Map{"message": "post deleted successfully"})
}

func (h *Post) GetPost(context echo.Context) error {
	var (
		ctx = custom.NewEchoCustom(context)
	)

	// Make authentication optional for guests
	userID := 0
	if id, err := ctx.GetCurrentUserID(); err == nil {
		userID = id
	}

	postID, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": "invalid post id"})
	}

	res, err := h.UsecasePost.GetPost(ctx.CurrentCtx(), uint32(userID), uint32(postID))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *Post) GetPosts(context echo.Context) error {
	var (
		ctx = custom.NewEchoCustom(context)
	)

	// Make authentication optional for guests
	userID := 0
	if id, err := ctx.GetCurrentUserID(); err == nil {
		userID = id
	}

	page, _ := strconv.Atoi(context.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(context.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	search := context.QueryParam("search")

	res, err := h.UsecasePost.GetPosts(ctx.CurrentCtx(), uint32(userID), page, limit, search)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *Post) GetUserPosts(context echo.Context) error {
	var (
		ctx = custom.NewEchoCustom(context)
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	targetUserID, err := strconv.ParseUint(context.Param("userId"), 10, 32)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": "invalid user id"})
	}

	page, _ := strconv.Atoi(context.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(context.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	res, err := h.UsecasePost.GetUserPosts(ctx.CurrentCtx(), uint32(userID), uint32(targetUserID), page, limit)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *Post) LikePost(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.LikePostPayload{}
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	err = input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = h.UsecasePost.LikePost(ctx.CurrentCtx(), uint32(userID), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, echo.Map{"message": "post liked successfully"})
}

func (h *Post) UnlikePost(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.LikePostPayload{}
	)

	userID, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	err = input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = h.UsecasePost.UnlikePost(ctx.CurrentCtx(), uint32(userID), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, echo.Map{"message": "post unliked successfully"})
}
