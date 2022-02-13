package user

import (
	"github.com/andhikagama/os-api/domain/user/repository"
	"github.com/andhikagama/os-api/domain/user/service"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	container.Provide(repository.New)
	container.Provide(service.New)
}
