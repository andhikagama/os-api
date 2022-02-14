package service

import (
	"testing"

	"github.com/andhikagama/os-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_Service_Create_Positive(t *testing.T) {
	var (
		req  = dao.User{}
		want = dao.User{}
	)
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().Create(m.Ctx, req).Return(want, nil)

		got, err := m.service.Create(m.Ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}

func Test_Service_Create_Negative(t *testing.T) {
	var (
		req  = dao.User{}
		want = dao.User{}
	)
	t.Run("when failed on repository.Create error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().Create(m.Ctx, req).Return(want, m.ErrTesting)

		got, err := m.service.Create(m.Ctx, req)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
}
