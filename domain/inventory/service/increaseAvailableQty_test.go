package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Service_IncreaseAvailableQty_Positive(t *testing.T) {
	var (
		id        uint64  = 1
		qty       float64 = 2
		mapUpdate         = map[string]interface{}{
			"available_qty":  gorm.Expr("available_qty + ?", qty),
			"quarantine_qty": gorm.Expr("quarantine_qty - ?", qty),
		}
	)
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().UpdateByID(m.Ctx, id, mapUpdate).Return(nil)

		err := m.service.IncreaseAvailableQty(m.Ctx, id, qty)
		assert.Nil(t, err)
	})
}

func Test_Service_IncreaseAvailableQty_Negative(t *testing.T) {
	var (
		id        uint64  = 1
		qty       float64 = 2
		mapUpdate         = map[string]interface{}{
			"available_qty":  gorm.Expr("available_qty + ?", qty),
			"quarantine_qty": gorm.Expr("quarantine_qty - ?", qty),
		}
	)
	t.Run("when failed on repository.IncreaseAvailableQty error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		m.repository.EXPECT().UpdateByID(m.Ctx, id, mapUpdate).Return(m.ErrTesting)

		err := m.service.IncreaseAvailableQty(m.Ctx, id, qty)
		assert.Error(t, err)
	})
}
