package service

import (
	"github.com/andhikagama/os-api/domain/order/repository"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Service .
	Service interface {
		Create(ctx *utils.Context, request dao.Order) (dao.Order, error)
		GetByID(ctx *utils.Context, id uint64) (dao.Order, error)
		UpdatePayment(ctx *utils.Context, request order.UpdatePaymentRequest) error
	}

	service struct {
		repository repository.Repository
		resource   shared.Resource
	}
)

// New .
func New(
	repository repository.Repository,
	resource shared.Resource,
) (Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}
