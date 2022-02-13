package service

import (
	"github.com/andhikagama/os-api/domain/inventory/repository"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	Service interface {
		GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error)
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
