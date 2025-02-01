package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/usecase"
	"github.com/trungluongwww/auth/util/custom"
	"net/http"
)

type User struct {
	UsecaseUser usecase.User
}

func NewUser(UsecaseUser usecase.User) *User {
	return &User{
		UsecaseUser: UsecaseUser,
	}
}

func (*User) Ping(context echo.Context) error {
	ctx := custom.NewEchoCustom(context)

	fmt.Println(ctx.GetCurrentUserID())
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "pong",
	})
}

func (h *User) Register(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.RegisterPayload{}
	)

	err := input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	err = h.UsecaseUser.Register(ctx.CurrentCtx(), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, nil)
}

func (h *User) Login(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.LoginPayload{}
	)

	err := input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	res, err := h.UsecaseUser.Login(ctx.CurrentCtx(), *input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *User) Me(context echo.Context) error {
	var (
		ctx = custom.NewEchoCustom(context)
	)

	id, err := ctx.GetCurrentUserID()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	res, err := h.UsecaseUser.GetMe(ctx.CurrentCtx(), id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}

func (h *User) RefreshToken(context echo.Context) error {
	var (
		ctx   = custom.NewEchoCustom(context)
		input = &request.RefreshTokenPayload{}
	)
	err := input.Bind(ctx)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	res, err := h.UsecaseUser.RefreshToken(ctx.CurrentCtx(), input)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return context.JSON(http.StatusOK, res)
}
