package order

import (
	"github.com/andhikagama/os-api/config/consts/orderEnum"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	UpdatePaymentRequest struct {
		ID            uint64                  `param:"id" validate:"required"`
		PaymentStatus orderEnum.PaymentStatus `json:"paymentStatus" validate:"required"`
	}
)

func NewUpdatePaymentRequest(ctx *utils.Context) (UpdatePaymentRequest, error) {
	request := UpdatePaymentRequest{}

	if err := ctx.Bind(&request); err != nil {
		return UpdatePaymentRequest{}, err
	}

	if err := ctx.Validate(&request); err != nil {
		return UpdatePaymentRequest{}, err
	}

	return request, nil
}

func (req UpdatePaymentRequest) ToMapForUpdate() map[string]interface{} {
	return map[string]interface{}{
		"payment_status": req.PaymentStatus,
	}
}
