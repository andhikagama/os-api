package inventory

import (
	"testing"

	"github.com/andhikagama/os-api/handlers"
	mock_inventory "github.com/andhikagama/os-api/mocks/inventory"
	"github.com/andhikagama/os-api/shared/utils"
)

type mock struct {
	utils.BaseMock
	handler    *mock_inventory.MockHandler
	controller Controller
}

func setupMock(t *testing.T) mock {
	m := mock{}

	m.BaseMock = utils.SetupBaseMock(t)
	m.handler = mock_inventory.NewMockHandler(m.Ctrl)
	m.controller = Controller{
		Handler: handlers.Handler{
			Inventory: m.handler,
		},
	}

	return m
}
