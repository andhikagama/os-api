package inventory

import (
	"testing"

	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/stretchr/testify/assert"
)

func Test_Handler_GetAllPaginated_Positive(t *testing.T) {
	var (
		req  = inventory.PaginatedRequest{}
		want = inventory.PaginatedResponse{}
	)
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.invService.EXPECT().GetAllPaginated(m.Ctx, req).Return(want, nil)

		got, err := m.handler.GetAllPaginated(m.Ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}

func Test_Handler_GetAllPaginated_Negative(t *testing.T) {
	var (
		req  = inventory.PaginatedRequest{}
		want = inventory.PaginatedResponse{}
	)
	t.Run("when failed on repository.GetAllPaginated error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.invService.EXPECT().GetAllPaginated(m.Ctx, req).Return(want, m.ErrTesting)

		got, err := m.handler.GetAllPaginated(m.Ctx, req)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
}
