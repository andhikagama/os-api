package config

import (
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) {
	container.Provide(New)
}
