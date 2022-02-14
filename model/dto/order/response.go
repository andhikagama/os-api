package order

import (
	"github.com/andhikagama/os-api/config/consts/orderEnum"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto"
	"github.com/andhikagama/os-api/model/dto/orderDetail"
	"github.com/andhikagama/os-api/model/dto/user"
)

type (
	Response struct {
		ID            uint64                  `json:"id"`
		User          user.Response           `json:"user"`
		TotalQty      float64                 `json:"totalQty"`
		TotalAmount   float64                 `json:"totalAmount"`
		PaymentType   orderEnum.PaymentType   `json:"paymentType"`
		PaymentStatus orderEnum.PaymentStatus `json:"paymentStatus"`
		OrderDetails  orderDetail.Responses   `json:"orderDetails"`

		dto.BaseModelSoftDelete
	}
)

func NewResponse(o dao.Order) Response {
	resp := Response{
		ID:            o.ID,
		User:          user.NewResponse(o.User),
		TotalQty:      o.TotalQty,
		TotalAmount:   o.TotalAmount,
		PaymentType:   o.PaymentType,
		PaymentStatus: o.PaymentStatus,
		OrderDetails:  orderDetail.NewResponses(o.OrderDetails),
	}

	resp.BaseModelSoftDelete = dto.NewBaseModelSoftDelete(o.BaseModelSoftDelete)
	return resp
}
