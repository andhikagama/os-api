package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

const (
	segmentGetAllPaginated = "inventory.repository.GetAllPaginated"
	segmentGetByID         = "inventory.repository.GetByID"
	segmentUpdateByID      = "inventory.repository.UpdateByID"
)

type (
	// Repository .
	Repository interface {
		GetAllPaginated(ctx *utils.Context, paginatedRequest *inventory.PaginatedRequest) (dao.Inventories, error)
		GetByID(ctx *utils.Context, id uint64) (dao.Inventory, error)
		UpdateByID(ctx *utils.Context, id uint64, request map[string]interface{}) error
	}

	repository struct {
		resource shared.Resource
	}
)

// New .
func New(resource shared.Resource) Repository {
	return &repository{
		resource: resource,
	}
}
