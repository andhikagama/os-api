package service

import (
	"testing"

	"github.com/andhikagama/os-api/model/dao"

	"github.com/stretchr/testify/assert"
)

func Test_Service_GetByID_Positive(t *testing.T) {
	var (
		id   uint64 = 1
		want        = dao.Inventory{}
	)
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().GetByID(m.Ctx, id).Return(want, nil)

		got, err := m.service.GetByID(m.Ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}

func Test_Service_GetByID_Negative(t *testing.T) {
	var (
		id   uint64 = 1
		want        = dao.Inventory{}
	)
	t.Run("when failed on repository.GetByID error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().GetByID(m.Ctx, id).Return(want, m.ErrTesting)

		got, err := m.service.GetByID(m.Ctx, id)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
}
