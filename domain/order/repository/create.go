package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) Create(ctx *utils.Context, request dao.Order) (dao.Order, error) {
	tx := ctx.GetORMTransaction(r.resource.DB)
	if err := tx.
		Create(&request).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentCreate, err)
		return dao.Order{}, err
	}

	return request, nil
}
