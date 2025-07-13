package route

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/config"
	middleware "github.com/trungluongwww/auth/internal/midleware"
	"github.com/trungluongwww/auth/pkg/route/v1"
	"github.com/trungluongwww/auth/register"
)

func New(e *echo.Echo, r *register.Register, cfg config.Env) {
	e.Use(middleware.AppMiddleware(cfg)...)

	v1Group := e.Group("/api/v1")
	v1.New(v1Group, r, cfg)
}
