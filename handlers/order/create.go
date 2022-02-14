package order

import (
	"fmt"

	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared/utils"
)

func (h *handler) Create(ctx *utils.Context, request order.CreateRequest) (order.Response, error) {
	orderDao := request.ToDao()

	ctx.SetORMTransaction(h.resource.DB)

	for k, v := range orderDao.OrderDetails {
		inv, err := h.domain.Inventory.GetByID(ctx, v.InventoryID)
		if err != nil {
			return order.Response{}, err
		}

		orderDao.OrderDetails[k].WAC = inv.WAC
		orderDao.OrderDetails[k].Amount = v.Qty * inv.WAC
		orderDao.TotalQty += v.Qty
		orderDao.TotalAmount += orderDao.OrderDetails[k].Amount

		if err := h.domain.Inventory.DecreaseAvailableQty(ctx, v.InventoryID, v.Qty); err != nil {
			ctx.RollBackORMTransaction()
			if utils.IsInsufficientError(err) {
				return order.Response{}, fmt.Errorf("%w for sku_code %v, sku_name %v", consts.ErrOrderInsufficientQty, inv.SkuCode, inv.SkuName)
			}

			return order.Response{}, err
		}
	}

	o, err := h.domain.Order.Create(ctx, orderDao)
	if err != nil {
		ctx.RollBackORMTransaction()

		return order.Response{}, err
	}

	ctx.CommitORMTransaction()

	res, err := h.domain.Order.GetByID(ctx, o.ID)
	if err != nil {
		return order.Response{}, err
	}

	return order.NewResponse(res), nil
}
