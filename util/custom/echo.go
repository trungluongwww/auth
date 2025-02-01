package custom

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strconv"
)

func NewEchoCustom(c echo.Context) *EchoCustom {
	return &EchoCustom{c}
}

type EchoCustom struct {
	echo.Context
}

func (c *EchoCustom) CurrentCtx() context.Context {
	return c.Request().Context()
}

func (c *EchoCustom) GetHeaderByKey(key string) string {
	return c.Request().Header.Get(key)
}

func (c *EchoCustom) GetCurrentUserID() (int, error) {
	token := c.Get("user")
	if token == nil {
		return 0, errors.New("invalid token")
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token")
	}

	claims, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid {
		strID := claims["sub"].(string)
		id, err := strconv.Atoi(strID)
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, errors.New("invalid token")
}
