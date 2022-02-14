package inventory

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/andhikagama/os-api/domain"
	mock_inventory "github.com/andhikagama/os-api/mocks/inventory"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type mock struct {
	utils.BaseMock
	invService *mock_inventory.MockService
	handler    Handler
}

func TestNew(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		d := domain.Domain{}
		resource := shared.Resource{}
		h, err := NewHandler(d, resource)

		assert.Nil(t, err)
		assert.Equal(t, &handler{
			domain:   d,
			resource: resource,
		}, h)
	})
}

func setupMock(t *testing.T) mock {
	m := mock{}
	m.BaseMock = utils.SetupBaseMock(t)
	m.invService = mock_inventory.NewMockService(m.Ctrl)

	d := domain.Domain{
		Inventory: m.invService,
	}

	h, _ := NewHandler(d, m.Resource)
	m.handler = h

	return m
}
