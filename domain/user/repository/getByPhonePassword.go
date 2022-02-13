package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) GetByPhonePassword(ctx *utils.Context, phone string, password string) (dao.User, error) {
	var res dao.User

	tx := ctx.GetORMTransaction(r.resource.DB)

	if err := tx.
		Model(&dao.User{}).
		Where("phone = ?", phone).
		Where("password = ?", password).
		First(&res).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentGetByPhonePassword, err)
		return dao.User{}, err
	}

	return res, nil
}
