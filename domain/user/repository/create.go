package repository

import (
	"fmt"

	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) Create(ctx *utils.Context, request dao.User) (dao.User, error) {
	tx := ctx.GetORMTransaction(r.resource.DB)
	if err := tx.
		Create(&request).
		Error; err != nil {
		fmt.Println(err)
		return dao.User{}, err
	}

	return request, nil
}
