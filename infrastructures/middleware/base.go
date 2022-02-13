package middleware

import (
	"github.com/andhikagama/os-api/config"
	"github.com/labstack/echo/v4"
)

// Middleware .
type Middleware struct {
	Config *config.EnvConfig
	Echo   *echo.Echo
}

// New ...
func New() Middleware {
	return Middleware{}
}
