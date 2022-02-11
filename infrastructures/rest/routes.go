package rest

import (
	"github.com/andhikagama/os-api/config/consts"
	"github.com/andhikagama/os-api/shared"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	resource   shared.Resource
	controller Controller
}

func New(resource shared.Resource, controller Controller) Routes {
	return Routes{
		resource:   resource,
		controller: controller,
	}
}

func (r *Routes) Register(v1 *echo.Group) {
	r.resource.Echo.GET("/health", r.controller.Health.Check).Name = consts.PrivilegePublic
}
