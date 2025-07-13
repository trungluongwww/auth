package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trungluongwww/auth/config"
	"github.com/trungluongwww/auth/util/custom"
	"net/http"
)

func AppMiddleware(env config.Env) []echo.MiddlewareFunc {
	loggerConfig := middleware.LoggerConfig{
		Format: "\n👉 ${time_rfc3339_nano} ${remote_ip} -> ${uri} ${method} ${status}, latency: ${latency_human} ${error}",
	}
	return []echo.MiddlewareFunc{
		cors,
		middleware.LoggerWithConfig(loggerConfig),
		middleware.Recover(),
	}
}

var allowHeaders []string

var allowMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
}

var cors = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: allowMethods,
	AllowHeaders: allowHeaders,
	MaxAge:       600,
})

func Jwt(key string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			ctx := custom.NewEchoCustom(c)
			authHeader := ctx.GetHeaderByKey(config.HeaderAuthorization)
			return authHeader == "" || authHeader == "Bearer" || authHeader == "Bearer "
		},
	})
}
