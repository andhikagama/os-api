package domain

import (
	"go.uber.org/dig"

	user "github.com/andhikagama/os-api/domain/user/service"
)

type (
	Domain struct {
		dig.In

		User user.Service
	}
)
