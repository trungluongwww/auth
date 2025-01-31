package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func (i *RegisterPayload) Validate() error {
	return validator.New().Struct(i)
}
func (i *RegisterPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}

	return i.Validate()
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (i *LoginPayload) Validate() error {
	return validator.New().Struct(i)
}
func (i *LoginPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}

	return i.Validate()
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

func (i *RefreshTokenPayload) Validate() error {
	return validator.New().Struct(i)
}
func (i *RefreshTokenPayload) Bind(ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return err
	}

	return i.Validate()
}
