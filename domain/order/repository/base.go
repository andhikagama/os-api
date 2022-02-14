package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

const (
	segmentCreate  = "order.repository.Create"
	segmentGetByID = "order.repository.GetByID"
)

type (
	// Repository .
	Repository interface {
		Create(ctx *utils.Context, request dao.Order) (dao.Order, error)
		GetByID(ctx *utils.Context, id uint64) (dao.Order, error)
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
