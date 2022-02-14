package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) UpdateByID(ctx *utils.Context, id uint64, request map[string]interface{}) error {
	tx := ctx.GetORMTransaction(r.resource.DB)
	if err := tx.
		Model(&dao.Order{}).
		Where("id = ?", id).
		Updates(request).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentUpdateByID, err)
		return err
	}

	return nil
}
