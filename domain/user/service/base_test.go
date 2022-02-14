package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_user "github.com/andhikagama/os-api/mocks/user"
	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/utils"
)

type mock struct {
	utils.BaseMock
	repository *mock_user.MockRepository
	service    Service
}

func TestNew(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := mock_user.NewMockRepository(ctrl)

		resource := shared.Resource{}
		svc, err := New(repository, resource)

		assert.Nil(t, err)
		assert.Equal(t, &service{
			repository: repository,
			resource:   resource,
		}, svc)
	})
}

func setupMock(t *testing.T) mock {
	m := mock{}
	m.BaseMock = utils.SetupBaseMock(t)
	m.repository = mock_user.NewMockRepository(m.Ctrl)

	svc, _ := New(m.repository, m.Resource)
	m.service = svc

	return m
}
