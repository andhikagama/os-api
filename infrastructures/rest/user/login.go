package user

import (
	"net/http"

	"github.com/andhikagama/os-api/config/consts"

	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
)

// Login .
func (c Controller) Login(ec echo.Context) error {
	ctx := utils.NewContext(ec, c.Resource.Validator)

	request, err := user.NewLoginRequest(ctx)
	if err != nil {
		return ctx.ResponseErrorJSON(http.StatusBadRequest, err)
	}

	result, err := c.Handler.User.Login(ctx, request)
	if err != nil {
		return ctx.ResponseErrorJSONFromList(err, consts.ErrListUser)
	}

	return ctx.ResponseSuccessJSON(result)
}
