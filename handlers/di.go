package handlers

import (
	"github.com/andhikagama/os-api/handlers/user"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

// Register .
func Register(container *dig.Container) error {
	if err := container.Provide(user.NewHandler); err != nil {
		return errors.Wrap(err, "failed to provide user handler")
	}

	return nil
}
