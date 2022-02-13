package service

import (
	"github.com/andhikagama/os-api/domain/user/repository"
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Service .
	Service interface {
		Create(ctx *utils.Context, request dao.User) (dao.User, error)
		Login(ctx *utils.Context, request user.LoginRequest) (dao.User, error)
	}

	service struct {
		resource   shared.Resource
		repository repository.Repository
	}
)

// New .
func New(
	resource shared.Resource,
	repository repository.Repository) Service {
	return &service{
		resource:   resource,
		repository: repository,
	}
}
