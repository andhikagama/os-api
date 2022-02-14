// Code generated by MockGen. DO NOT EDIT.
// Source: domain/inventory/service/base.go

// Package mock_inventory is a generated GoMock package.
package mock_inventory

import (
	reflect "reflect"

	dao "github.com/andhikagama/os-api/model/dao"
	inventory "github.com/andhikagama/os-api/model/dto/inventory"
	utils "github.com/andhikagama/os-api/shared/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DecreaseAvailableQty mocks base method.
func (m *MockService) DecreaseAvailableQty(ctx *utils.Context, id uint64, qty float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseAvailableQty", ctx, id, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecreaseAvailableQty indicates an expected call of DecreaseAvailableQty.
func (mr *MockServiceMockRecorder) DecreaseAvailableQty(ctx, id, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseAvailableQty", reflect.TypeOf((*MockService)(nil).DecreaseAvailableQty), ctx, id, qty)
}

// GetAllPaginated mocks base method.
func (m *MockService) GetAllPaginated(ctx *utils.Context, paginatedRequest inventory.PaginatedRequest) (inventory.PaginatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaginated", ctx, paginatedRequest)
	ret0, _ := ret[0].(inventory.PaginatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPaginated indicates an expected call of GetAllPaginated.
func (mr *MockServiceMockRecorder) GetAllPaginated(ctx, paginatedRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaginated", reflect.TypeOf((*MockService)(nil).GetAllPaginated), ctx, paginatedRequest)
}

// GetByID mocks base method.
func (m *MockService) GetByID(ctx *utils.Context, id uint64) (dao.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(dao.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockServiceMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockService)(nil).GetByID), ctx, id)
}

// IncreaseAvailableQty mocks base method.
func (m *MockService) IncreaseAvailableQty(ctx *utils.Context, id uint64, qty float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseAvailableQty", ctx, id, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseAvailableQty indicates an expected call of IncreaseAvailableQty.
func (mr *MockServiceMockRecorder) IncreaseAvailableQty(ctx, id, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseAvailableQty", reflect.TypeOf((*MockService)(nil).IncreaseAvailableQty), ctx, id, qty)
}
