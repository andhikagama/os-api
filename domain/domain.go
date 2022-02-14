package domain

import (
	"go.uber.org/dig"

	inventory "github.com/andhikagama/os-api/domain/inventory/service"
	order "github.com/andhikagama/os-api/domain/order/service"
	user "github.com/andhikagama/os-api/domain/user/service"
)

type (
	Domain struct {
		dig.In

		User      user.Service
		Inventory inventory.Service
		Order     order.Service
	}
)
