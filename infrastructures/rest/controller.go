package rest

import (
	"github.com/andhikagama/os-api/infrastructures/rest/health"
	"go.uber.org/dig"
)

type (
	// Controller .
	Controller struct {
		dig.In

		Health health.Controller
	}
)
