package order

import (
	"net/http"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
)

// UpdatePayment .
func (c Controller) UpdatePayment(ec echo.Context) error {
	ctx := utils.NewContext(ec, c.Resource.Validator)

	request, err := order.NewUpdatePaymentRequest(ctx)
	if err != nil {
		return ctx.ResponseErrorJSON(http.StatusBadRequest, err)
	}

	result, err := c.Handler.Order.UpdatePayment(ctx, request)
	if err != nil {
		return ctx.ResponseErrorJSONFromList(err, consts.ErrListOrder)
	}

	return ctx.ResponseCreatedJSON(result)
}
