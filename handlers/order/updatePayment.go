package order

import (
	"github.com/andhikagama/os-api/config/consts/orderEnum"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared/utils"
)

func (h *handler) UpdatePayment(ctx *utils.Context, request order.UpdatePaymentRequest) (order.Response, error) {
	o, err := h.domain.Order.GetByID(ctx, request.ID)
	if err != nil {
		return order.Response{}, err
	}

	ctx.SetORMTransaction(h.resource.DB)

	if err := h.domain.Order.UpdatePayment(ctx, request); err != nil {
		ctx.RollBackORMTransaction()
		return order.Response{}, err
	}

	switch request.PaymentStatus {
	case orderEnum.PaymentStatusExpired:
		if err := h.domain.Inventory.UpdateExpired(ctx, o.OrderDetails); err != nil {
			ctx.RollBackORMTransaction()
			return order.Response{}, err
		}
		break
	case orderEnum.PaymentStatusPaid:
		if err := h.domain.Inventory.UpdatePaid(ctx, o.OrderDetails); err != nil {
			ctx.RollBackORMTransaction()
			return order.Response{}, err
		}
		break
	}

	ctx.CommitORMTransaction()

	res, err := h.domain.Order.GetByID(ctx, request.ID)
	if err != nil {
		return order.Response{}, err
	}

	return order.NewResponse(res), nil
}
