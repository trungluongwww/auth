package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/pkg/handler"
)

func user(e *echo.Group, auth echo.MiddlewareFunc, userHandler *handler.User) {
	g := e.Group("/users")

	g.GET("/me", userHandler.Me, auth)
}
