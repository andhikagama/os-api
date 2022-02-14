package order

import (
	"github.com/andhikagama/os-api/domain/order/repository"
	"github.com/andhikagama/os-api/domain/order/service"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	container.Provide(repository.New)
	container.Provide(service.New)
}
