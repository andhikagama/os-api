package order

import (
	"github.com/andhikagama/os-api/config/consts/orderEnum"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/orderDetail"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	CreateRequest struct {
		UserID       uint64                        `json:"-" validate:"required"`
		PaymentType  orderEnum.PaymentType         `json:"paymentType" validate:"required"`
		OrderDetails orderDetail.BulkCreateRequest `json:"orderDetails" validate:"required,dive"`
	}
)

func NewCreateRequest(ctx *utils.Context) (CreateRequest, error) {
	request := CreateRequest{}

	if err := ctx.Bind(&request); err != nil {
		return CreateRequest{}, err
	}

	request.UserID = ctx.GetUserID()
	if err := ctx.Validate(&request); err != nil {
		return CreateRequest{}, err
	}

	return request, nil
}

func (req CreateRequest) ToDao() dao.Order {
	return dao.Order{
		UserID:        req.UserID,
		PaymentType:   req.PaymentType,
		PaymentStatus: orderEnum.PaymentStatusPending,
		OrderDetails:  req.OrderDetails.ToDao(),
	}
}
