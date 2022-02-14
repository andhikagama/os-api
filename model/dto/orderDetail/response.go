package orderDetail

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto"
	"github.com/andhikagama/os-api/model/dto/inventory"
)

type (
	Response struct {
		OrderID   uint64             `json:"orderId"`
		Inventory inventory.Response `json:"inventory"`
		Qty       float64            `json:"qty"`
		Amount    float64            `json:"amount"`

		dto.BaseModelSoftDelete
	}

	Responses []Response
)

func NewResponse(od dao.OrderDetail) Response {
	resp := Response{
		OrderID:   od.OrderID,
		Inventory: inventory.NewResponse(od.Inventory),
		Qty:       od.Qty,
		Amount:    od.Amount,
	}

	resp.BaseModelSoftDelete = dto.NewBaseModelSoftDelete(od.BaseModelSoftDelete)
	return resp
}

func NewResponses(ods dao.OrderDetails) Responses {
	resp := make(Responses, len(ods))

	for k, v := range ods {
		resp[k] = NewResponse(v)
	}

	return resp
}
