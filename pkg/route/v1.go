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
	postHandler := handler.NewPost(r.NewUsecasePost())

	// guest
	e.GET("/ping", userHandler.Ping)
	e.POST("/login", userHandler.Login)
	e.POST("/facebook", userHandler.LoginWithFacebook)
	e.POST("/register", userHandler.Register)
	e.POST("/refresh-token", userHandler.RefreshToken)

	// user
	e.GET("/users/me", userHandler.Me, auth)

	// posts (authenticated)
	e.POST("/posts", postHandler.CreatePost, auth)
	e.PUT("/posts/:id", postHandler.UpdatePost, auth)
	e.DELETE("/posts/:id", postHandler.DeletePost, auth)
	e.GET("/posts/:id", postHandler.GetPost, auth)
	e.GET("/posts", postHandler.GetPosts, auth)
	e.GET("/users/:userId/posts", postHandler.GetUserPosts, auth)
	e.POST("/posts/like", postHandler.LikePost, auth)
	e.DELETE("/posts/like", postHandler.UnlikePost, auth)
}
