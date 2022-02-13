package inventory

import (
	"github.com/andhikagama/os-api/domain"
	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Handler .
	Handler interface {
		GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error)
	}

	handler struct {
		domain   domain.Domain
		resource shared.Resource
	}
)

// NewHandler .
func NewHandler(domain domain.Domain, resource shared.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
