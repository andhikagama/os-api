package rest

import (
	"github.com/andhikagama/os-api/infrastructures/rest/health"
	"github.com/andhikagama/os-api/infrastructures/rest/inventory"
	"github.com/andhikagama/os-api/infrastructures/rest/user"
	"go.uber.org/dig"
)

type (
	// Controller .
	Controller struct {
		dig.In

		Health    health.Controller
		User      user.Controller
		Inventory inventory.Controller
	}
)
