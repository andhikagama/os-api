package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var corsConfig = echoMiddleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
}

// SetCORS .
func (m *Middleware) SetCORS() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(corsConfig)
}

// RemoveTrailingSlash .
func (m *Middleware) RemoveTrailingSlash() echo.MiddlewareFunc {
	return echoMiddleware.RemoveTrailingSlash()
}
