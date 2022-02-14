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

	users := v1.Group("/users")
	users.POST("", r.controller.User.Create).Name = consts.PrivilegeTrusted
	users.POST("/login", r.controller.User.Login).Name = consts.PrivilegeTrusted

	inventories := v1.Group("/inventories")
	inventories.GET("", r.controller.Inventory.GetAllPaginated).Name = consts.PrivilegeGranted

	orders := v1.Group("/orders")
	orders.POST("", r.controller.Order.Create).Name = consts.PrivilegeGranted
	orders.PATCH("/:id/payment", r.controller.Order.UpdatePayment).Name = consts.PrivilegeGranted
}
