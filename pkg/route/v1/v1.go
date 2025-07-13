package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/config"
	middleware "github.com/trungluongwww/auth/internal/midleware"
	"github.com/trungluongwww/auth/pkg/handler"
	"github.com/trungluongwww/auth/register"
)

func New(e *echo.Group, r *register.Register, cfg config.Env) {
	auth := middleware.Jwt(cfg.SecretUserJWTToken)
	userHandler := handler.NewUser(r.NewUsecaseUser())
	postHandler := handler.NewPost(r.NewUsecasePost())

	public(e, userHandler)
	post(e, auth, postHandler)
	user(e, auth, userHandler)
}

func public(e *echo.Group, userHandler *handler.User) {
	e.GET("/ping", userHandler.Ping)
	e.POST("/login", userHandler.Login)
	e.POST("/facebook", userHandler.LoginWithFacebook)
	e.POST("/register", userHandler.Register)
	e.POST("/refresh-token", userHandler.RefreshToken)
}
