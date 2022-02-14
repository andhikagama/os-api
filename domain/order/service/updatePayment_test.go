package service

import (
	"testing"

	"github.com/andhikagama/os-api/config/consts/orderEnum"
	"github.com/andhikagama/os-api/model/dto/order"

	"github.com/stretchr/testify/assert"
)

func Test_Service_UpdatePayment_Positive(t *testing.T) {
	var (
		id uint64 = 1
	)
	t.Run("when success paid", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		var (
			req = order.UpdatePaymentRequest{
				ID:            id,
				PaymentStatus: orderEnum.PaymentStatusPaid,
			}
		)

		m.repository.EXPECT().UpdateByID(m.Ctx, req.ID, req.ToMapForUpdate()).Return(nil)

		err := m.service.UpdatePayment(m.Ctx, req)
		assert.Nil(t, err)
	})

	t.Run("when success expired", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		var (
			req = order.UpdatePaymentRequest{
				ID:            id,
				PaymentStatus: orderEnum.PaymentStatusExpired,
			}
		)

		m.repository.EXPECT().UpdateByID(m.Ctx, req.ID, req.ToMapForUpdate()).Return(nil)

		err := m.service.UpdatePayment(m.Ctx, req)
		assert.Nil(t, err)
	})
}

func Test_Service_UpdatePayment_Negative(t *testing.T) {
	var (
		id uint64 = 1
	)
	t.Run("when failed on repository.UpdateByID error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		var (
			req = order.UpdatePaymentRequest{
				ID:            id,
				PaymentStatus: orderEnum.PaymentStatusPaid,
			}
		)

		m.repository.EXPECT().UpdateByID(m.Ctx, req.ID, req.ToMapForUpdate()).Return(m.ErrTesting)

		err := m.service.UpdatePayment(m.Ctx, req)
		assert.Error(t, err)
	})
}
