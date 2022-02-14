package inventory

import (
	"net/http"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
)

// GetAllPaginated .
func (c Controller) GetAllPaginated(ec echo.Context) error {
	ctx := utils.NewContext(ec, c.Resource.Validator)

	request, err := inventory.NewPaginatedRequest(ctx)
	if err != nil {
		return ctx.ResponseErrorJSON(http.StatusBadRequest, err)
	}

	result, err := c.Handler.Inventory.GetAllPaginated(ctx, request)
	if err != nil {
		return ctx.ResponseErrorJSONFromList(err, consts.ErrListInventory)
	}

	return ctx.ResponseSuccessJSON(result)
}
