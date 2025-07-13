package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/auth/pkg/handler"
)

func post(e *echo.Group, auth echo.MiddlewareFunc, postHandler *handler.Post) {
	g := e.Group("/posts")

	// public routes
	g.GET("", postHandler.GetPosts)
	g.GET("/:id", postHandler.GetPost)

	// auth routes
	g.POST("", postHandler.CreatePost, auth)
	g.PUT("/:id", postHandler.UpdatePost, auth)
	g.DELETE("/:id", postHandler.DeletePost, auth)
	g.GET("/users/:userId/posts", postHandler.GetUserPosts, auth)
	g.POST("/like", postHandler.LikePost, auth)
	g.DELETE("/like", postHandler.UnlikePost, auth)
}
