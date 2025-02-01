package route

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/config"
	middleware "github.com/trungluongwww/auth/internal/midleware"
	"github.com/trungluongwww/auth/pkg/handler"
	"github.com/trungluongwww/auth/register"
)

func newV1(e *echo.Group, r *register.Register, cfg config.Env) {
	auth := middleware.Jwt(cfg.SecretUserJWTToken)
	userHandler := handler.NewUser(r.NewUsecaseUser())

	// guest
	e.GET("/ping", userHandler.Ping)
	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)
	e.POST("/refresh-token", userHandler.RefreshToken)

	// user
	e.GET("/users/me", userHandler.Me, auth)
}
