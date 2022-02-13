package repository

import (
	"github.com/andhikagama/os-api/model/dao"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared/utils"
)

func (r repository) GetAllPaginated(ctx *utils.Context, paginatedRequest *inventory.PaginatedRequest) (dao.Inventories, error) {
	var res dao.Inventories

	tx := ctx.GetORMTransaction(r.resource.DB).
		Model(&dao.Inventory{})

	paginatedRequest.Filter.Apply(tx)
	paginatedRequest.ToPaginatedQuery(tx, paginatedRequest.AvailableSort())

	if err := tx.
		Find(&res).
		Error; err != nil {
		r.resource.Logger.Errorf("%v error %w", segmentGetAllPaginated, err)
		return dao.Inventories{}, err
	}

	return res, nil
}
