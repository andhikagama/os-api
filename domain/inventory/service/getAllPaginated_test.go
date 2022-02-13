package service

import (
	"testing"

	"github.com/andhikagama/os-api/model/dao"

	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/stretchr/testify/assert"
)

func Test_Service_GetAllPaginated_Positive(t *testing.T) {
	var (
		req   = inventory.PaginatedRequest{}
		items = dao.Inventories{}
		want  = inventory.PaginatedResponse{
			Items: inventory.NewResponses(items),
		}
	)
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().GetAllPaginated(m.Ctx, &req).Return(items, nil)

		got, err := m.service.GetAllPaginated(m.Ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}

func Test_Service_GetAllPaginated_Negative(t *testing.T) {
	var (
		req  = inventory.PaginatedRequest{}
		want = inventory.PaginatedResponse{}
	)
	t.Run("when failed on repository.GetAllPaginated error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().GetAllPaginated(m.Ctx, &req).Return(dao.Inventories{}, m.ErrTesting)

		got, err := m.service.GetAllPaginated(m.Ctx, req)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
}
