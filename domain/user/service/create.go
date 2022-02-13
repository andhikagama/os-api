package service

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (s service) Create(ctx *utils.Context, request dao.User) (dao.User, error) {
	return s.repository.Create(ctx, request)
}
