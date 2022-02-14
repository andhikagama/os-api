package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) GetByID(ctx *utils.Context, id uint64) (dao.Inventory, error) {
	var res dao.Inventory

	tx := ctx.GetORMTransaction(r.resource.DB).
		Model(&dao.Inventory{}).
		Where("id", id)

	if err := tx.
		Take(&res).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentGetByID, err)
		return dao.Inventory{}, err
	}

	return res, nil
}
