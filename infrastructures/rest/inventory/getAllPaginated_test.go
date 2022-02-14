package inventory

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andhikagama/os-api/config/consts"

	"github.com/andhikagama/os-api/model/dto/inventory"
	"github.com/andhikagama/os-api/shared/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_Controller_GetAllPaginated_Positive(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		var (
			req          = httptest.NewRequest(http.MethodGet, "/inventories", nil)
			res          = httptest.NewRecorder()
			echoCtx      = echo.New().NewContext(req, res)
			ctx          = utils.NewContext(echoCtx, nil)
			paginatedReq = inventory.PaginatedRequest{}
			paginatedRes = inventory.PaginatedResponse{}
			wantStatus   = http.StatusOK
			want         = m.BaseMock.MockDataResponse(paginatedRes)
		)

		paginatedReq.SetDefault()
		m.handler.EXPECT().GetAllPaginated(ctx, paginatedReq).Return(paginatedRes, nil)

		err := m.controller.GetAllPaginated(echoCtx)

		assert.Nil(t, err)
		assert.Equal(t, wantStatus, res.Code)
		assert.Equal(t, want, res.Body.String())
	})
}

func Test_Controller_GetAllPaginated_Negative(t *testing.T) {
	t.Run("when failed on Handler.Inventory.GetAllPaginated error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		var (
			req          = httptest.NewRequest(http.MethodGet, "/inventories", nil)
			res          = httptest.NewRecorder()
			echoCtx      = echo.New().NewContext(req, res)
			ctx          = utils.NewContext(echoCtx, nil)
			paginatedReq = inventory.PaginatedRequest{}
			paginatedRes = inventory.PaginatedResponse{}
			wantStatus   = http.StatusInternalServerError
			want         = m.BaseMock.MockErrorResponse(consts.ErrInternalServerError.Error(), wantStatus)
		)

		paginatedReq.SetDefault()
		m.handler.EXPECT().GetAllPaginated(ctx, paginatedReq).Return(paginatedRes, m.ErrTesting)

		err := m.controller.GetAllPaginated(echoCtx)

		assert.Nil(t, err)
		assert.Equal(t, wantStatus, res.Code)
		assert.Equal(t, want, res.Body.String())
	})
}
