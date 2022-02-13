package di

import (
	"github.com/andhikagama/os-api/config"
	"github.com/andhikagama/os-api/domain"
	"github.com/andhikagama/os-api/handlers"
	"github.com/andhikagama/os-api/infrastructures"
	"github.com/andhikagama/os-api/shared"
	"go.uber.org/dig"
)

// BuildContainer .
func BuildContainer() *dig.Container {
	container := dig.New()

	config.Register(container)
	shared.Register(container)
	domain.Register(container)
	handlers.Register(container)
	infrastructures.Register(container)

	return container
}
