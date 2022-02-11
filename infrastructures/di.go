package infrastructures

import (
	"github.com/andhikagama/os-api/infrastructures/rest"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	container.Provide(rest.New)
	container.Provide(New)
}
