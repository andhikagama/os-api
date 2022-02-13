package handlers

import (
	"github.com/andhikagama/os-api/handlers/user"
	"go.uber.org/dig"
)

type (
	// Handler .
	Handler struct {
		dig.In

		User user.Handler
	}
)
