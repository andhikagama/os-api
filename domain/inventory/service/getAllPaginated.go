package service

import (
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared/utils"
)

func (s service) GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error) {
	items, err := s.repository.GetAllPaginated(ctx, &paginatedRequest)
	if err != nil {
		return inventory.PaginatedResponse{}, err
	}

	return inventory.PaginatedResponse{
		PaginatedRequest: paginatedRequest,
		Items:            inventory.NewResponses(items),
	}, nil
}
