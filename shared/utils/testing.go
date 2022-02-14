package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/andhikagama/os-api/shared"
	"github.com/andhikagama/os-api/shared/resource"

	"github.com/golang/mock/gomock"
)

type (
	BaseMock struct {
		Ctrl       *gomock.Controller
		Ctx        *Context
		Resource   shared.Resource
		ErrTesting error
	}

	ErrorData struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	ErrorResponse struct {
		Error *ErrorData `json:"error,omitempty"`
	}
)

func SetupBaseMock(t *testing.T) BaseMock {
	ctrl := gomock.NewController(t)
	r := shared.Resource{
		Logger:    resource.NewLogger(),
		Echo:      resource.NewEcho(),
		Validator: resource.NewValidator(),
	}

	return BaseMock{
		Ctrl:       ctrl,
		Ctx:        &Context{},
		Resource:   r,
		ErrTesting: errors.New("any error for testing"),
	}
}

func (bm BaseMock) MockDataResponse(expectedContent interface{}) string {
	expectedResponse := struct {
		Data interface{} `json:"data"`
	}{Data: expectedContent}

	expectedByte, _ := json.Marshal(expectedResponse)
	expectedResult := string(expectedByte) + "\n"
	return expectedResult
}

func (bm BaseMock) MockErrorResponse(msg string, statusCode int) string {
	return fmt.Sprintf(`{"error":{"message":"%s","messages":["%s"]}}%s`, msg, msg, "\n")
}
