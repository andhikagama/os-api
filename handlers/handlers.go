package handlers

import (
	"github.com/andhikagama/os-api/handlers/inventory"
	"github.com/andhikagama/os-api/handlers/order"
	"github.com/andhikagama/os-api/handlers/user"
	"go.uber.org/dig"
)

type (
	// Handler .
	Handler struct {
		dig.In

		User      user.Handler
		Inventory inventory.Handler
		Order     order.Handler
	}
)
