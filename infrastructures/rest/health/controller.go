package health

import (
	"github.com/andhikagama/os-api/shared"
	"go.uber.org/dig"
)

type (
	Controller struct {
		dig.In

		Resource shared.Resource
	}
)
