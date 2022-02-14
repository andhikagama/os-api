package order

import (
	"github.com/andhikagama/os-api/domain"
	"github.com/andhikagama/os-api/model/dto/order"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Handler .
	Handler interface {
		Create(ctx *utils.Context, request order.CreateRequest) (order.Response, error)
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
