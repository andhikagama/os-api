package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) Check(ec echo.Context) error {
	return ec.String(http.StatusOK, "ok!")
}
