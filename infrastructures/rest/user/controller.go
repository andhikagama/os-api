package user

import (
	"github.com/andhikagama/os-api/handlers"
	"github.com/andhikagama/os-api/shared"
	"go.uber.org/dig"
)

type (
	// Controller .
	Controller struct {
		dig.In

		Handler  handlers.Handler
		Resource shared.Resource
	}
)
