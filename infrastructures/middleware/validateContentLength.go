package middleware

import (
	"fmt"
	"net/http"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/shared/utils"

	"github.com/labstack/echo/v4"
)

// ValidateContentLength .
func (m *Middleware) ValidateContentLength(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		fmt.Println(1)
		switch req.Method {
		case echo.PATCH, echo.PUT, echo.POST:
			if req.ContentLength != 0 {
				return next(c)
			}
			break
		case echo.GET, echo.DELETE:
			return next(c)
		}

		fmt.Println(2)
		ctx := utils.NewContext(c, nil)
		return ctx.ResponseErrorJSON(http.StatusBadRequest, consts.ErrEmptyRequestBody)
	}
}
