package server

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/config"
	"github.com/trungluongwww/auth/pkg/route"
	"github.com/trungluongwww/auth/register"
)

func Bootstrap(r *register.Register, cfg config.Env) *echo.Echo {
	e := echo.New()

	route.New(e, r, cfg)
	return e
}
