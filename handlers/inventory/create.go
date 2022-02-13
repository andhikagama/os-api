package inventory

import (
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared/utils"
)

func (h *handler) GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error) {
	return h.domain.Inventory.GetAllPaginated(ctx, paginatedRequest)
}
