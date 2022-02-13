package inventory

import (
	"github.com/andhikagama/os-api/domain/inventory/repository"
	"github.com/andhikagama/os-api/domain/inventory/service"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	container.Provide(repository.New)
	container.Provide(service.New)
}
