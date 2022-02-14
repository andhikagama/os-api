package service

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (s service) GetByID(ctx *utils.Context, id uint64) (dao.Inventory, error) {
	return s.repository.GetByID(ctx, id)
}
