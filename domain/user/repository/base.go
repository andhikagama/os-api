package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Repository .
	Repository interface {
		Create(ctx *utils.Context, request dao.User) (dao.User, error)
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
