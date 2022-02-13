package domain

import (
	"github.com/andhikagama/os-api/domain/user"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	user.Register(container)
}
