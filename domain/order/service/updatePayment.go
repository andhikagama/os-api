package service

import (
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared/utils"
)

func (s service) UpdatePayment(ctx *utils.Context, request order.UpdatePaymentRequest) error {
	return s.repository.UpdateByID(ctx, request.ID, request.ToMapForUpdate())
}
