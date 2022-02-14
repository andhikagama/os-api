package order

import (
	"net/http"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
)

// Create .
func (c Controller) Create(ec echo.Context) error {
	ctx := utils.NewContext(ec, c.Resource.Validator)

	request, err := order.NewCreateRequest(ctx)
	if err != nil {
		return ctx.ResponseErrorJSON(http.StatusBadRequest, err)
	}

	result, err := c.Handler.Order.Create(ctx, request)
	if err != nil {
		return ctx.ResponseErrorJSONFromList(err, consts.ErrListOrder)
	}

	return ctx.ResponseCreatedJSON(result)
}
