package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

const (
	segmentCreate             = "user.repository.Create"
	segmentGetByPhonePassword = "user.repository.GetByPhonePassword"
	segmentUpdateByID         = "user.repository.UpdateByID"
)

type (
	// Repository .
	Repository interface {
		Create(ctx *utils.Context, request dao.User) (dao.User, error)
		GetByPhonePassword(ctx *utils.Context, phone string, password string) (dao.User, error)
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
