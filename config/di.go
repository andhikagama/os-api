package config

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) error {
	if err := container.Provide(New); err != nil {
		return errors.Wrap(err, "failed to provide config")
	}

	return nil
}
