package user

import (
	"github.com/andhikagama/os-api/domain"
	"github.com/andhikagama/os-api/model/dto/user"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type (
	// Handler .
	Handler interface {
		Create(ctx *utils.Context, request user.CreateRequest) (user.Response, error)
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
