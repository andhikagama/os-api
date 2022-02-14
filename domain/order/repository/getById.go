package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) GetByID(ctx *utils.Context, id uint64) (dao.Order, error) {
	var res dao.Order

	tx := ctx.GetORMTransaction(r.resource.DB).
		Model(&dao.Order{}).
		Preload("User").
		Preload("OrderDetails.Inventory").
		Where("id", id)

	if err := tx.
		Take(&res).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentGetByID, err)
		return dao.Order{}, err
	}

	return res, nil
}
