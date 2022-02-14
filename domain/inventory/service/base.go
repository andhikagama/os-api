package service

import (
	"github.com/andhikagama/os-api/domain/inventory/repository"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	Service interface {
		GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error)
		GetByID(ctx *utils.Context, id uint64) (dao.Inventory, error)
		DecreaseAvailableQty(ctx *utils.Context, id uint64, qty float64) error
		IncreaseAvailableQty(ctx *utils.Context, id uint64, qty float64) error
		UpdateExpired(ctx *utils.Context, req dao.OrderDetails) error
		UpdatePaid(ctx *utils.Context, req dao.OrderDetails) error
	}

	service struct {
		resource   shared.Resource
		repository repository.Repository
	}
)

func New(
	resource shared.Resource,
	repository repository.Repository) (Service, error) {
	return &service{
		resource:   resource,
		repository: repository,
	}, nil
}
