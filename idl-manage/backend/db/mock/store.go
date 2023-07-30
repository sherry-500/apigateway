// Code generated by MockGen. DO NOT EDIT.
// Source: idlManage/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	db "idlManage/db/sqlc"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateIdlmap mocks base method.
func (m *MockStore) CreateIdlmap(arg0 context.Context, arg1 db.CreateIdlmapParams) (db.Idlmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdlmap", arg0, arg1)
	ret0, _ := ret[0].(db.Idlmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdlmap indicates an expected call of CreateIdlmap.
func (mr *MockStoreMockRecorder) CreateIdlmap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdlmap", reflect.TypeOf((*MockStore)(nil).CreateIdlmap), arg0, arg1)
}

// DeleteIdlmap mocks base method.
func (m *MockStore) DeleteIdlmap(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIdlmap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIdlmap indicates an expected call of DeleteIdlmap.
func (mr *MockStoreMockRecorder) DeleteIdlmap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIdlmap", reflect.TypeOf((*MockStore)(nil).DeleteIdlmap), arg0, arg1)
}

// GetIdlmap mocks base method.
func (m *MockStore) GetIdlmap(arg0 context.Context, arg1 int64) (db.Idlmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdlmap", arg0, arg1)
	ret0, _ := ret[0].(db.Idlmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdlmap indicates an expected call of GetIdlmap.
func (mr *MockStoreMockRecorder) GetIdlmap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdlmap", reflect.TypeOf((*MockStore)(nil).GetIdlmap), arg0, arg1)
}

// GetIdlmapByName mocks base method.
func (m *MockStore) GetIdlmapByName(arg0 context.Context, arg1 string) (db.Idlmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdlmapByName", arg0, arg1)
	ret0, _ := ret[0].(db.Idlmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdlmapByName indicates an expected call of GetIdlmapByName.
func (mr *MockStoreMockRecorder) GetIdlmapByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdlmapByName", reflect.TypeOf((*MockStore)(nil).GetIdlmapByName), arg0, arg1)
}

// ListIdlmap mocks base method.
func (m *MockStore) ListIdlmap(arg0 context.Context, arg1 db.ListIdlmapParams) ([]db.Idlmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIdlmap", arg0, arg1)
	ret0, _ := ret[0].([]db.Idlmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdlmap indicates an expected call of ListIdlmap.
func (mr *MockStoreMockRecorder) ListIdlmap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdlmap", reflect.TypeOf((*MockStore)(nil).ListIdlmap), arg0, arg1)
}

// UpdateIdlmap mocks base method.
func (m *MockStore) UpdateIdlmap(arg0 context.Context, arg1 db.UpdateIdlmapParams) (db.Idlmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIdlmap", arg0, arg1)
	ret0, _ := ret[0].(db.Idlmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIdlmap indicates an expected call of UpdateIdlmap.
func (mr *MockStoreMockRecorder) UpdateIdlmap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdlmap", reflect.TypeOf((*MockStore)(nil).UpdateIdlmap), arg0, arg1)
}