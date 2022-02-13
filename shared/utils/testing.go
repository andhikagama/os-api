package utils

import (
	"errors"
	"testing"

	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/resource"

	"github.com/golang/mock/gomock"
)

type BaseMock struct {
	Ctrl       *gomock.Controller
	Ctx        *Context
	Resource   shared.Resource
	ErrTesting error
}

func SetupBaseMock(t *testing.T) BaseMock {
	ctrl := gomock.NewController(t)
	r := shared.Resource{
		Logger: resource.NewLogger(),
		Echo:   resource.NewEcho(),
	}

	return BaseMock{
		Ctrl:       ctrl,
		Ctx:        &Context{},
		Resource:   r,
		ErrTesting: errors.New("any error for testing"),
	}
}
