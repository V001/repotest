// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	model "repotest/model"

	gomock "github.com/golang/mock/gomock"
)

// MockIBookRepository is a mock of IBookRepository interface.
type MockIBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBookRepositoryMockRecorder
}

// MockIBookRepositoryMockRecorder is the mock recorder for MockIBookRepository.
type MockIBookRepositoryMockRecorder struct {
	mock *MockIBookRepository
}

// NewMockIBookRepository creates a new mock instance.
func NewMockIBookRepository(ctrl *gomock.Controller) *MockIBookRepository {
	mock := &MockIBookRepository{ctrl: ctrl}
	mock.recorder = &MockIBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBookRepository) EXPECT() *MockIBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIBookRepository) Create() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create")
}

// Create indicates an expected call of Create.
func (mr *MockIBookRepositoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBookRepository)(nil).Create))
}

// Delete mocks base method.
func (m *MockIBookRepository) Delete() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete")
}

// Delete indicates an expected call of Delete.
func (mr *MockIBookRepositoryMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIBookRepository)(nil).Delete))
}

// Get mocks base method.
func (m *MockIBookRepository) Get() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get")
}

// Get indicates an expected call of Get.
func (mr *MockIBookRepositoryMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIBookRepository)(nil).Get))
}

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockIUserRepository) Auth(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockIUserRepositoryMockRecorder) Auth(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockIUserRepository)(nil).Auth), ctx, user)
}

// CheckIsPhoneExist mocks base method.
func (m *MockIUserRepository) CheckIsPhoneExist(ctx context.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsPhoneExist", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsPhoneExist indicates an expected call of CheckIsPhoneExist.
func (mr *MockIUserRepositoryMockRecorder) CheckIsPhoneExist(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsPhoneExist", reflect.TypeOf((*MockIUserRepository)(nil).CheckIsPhoneExist), ctx, username)
}

// Create mocks base method.
func (m *MockIUserRepository) Create(ctx context.Context, user model.User) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), ctx, user)
}

// Delete mocks base method.
func (m *MockIUserRepository) Delete(ctx context.Context, ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepositoryMockRecorder) Delete(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepository)(nil).Delete), ctx, ID)
}

// GetAll mocks base method.
func (m *MockIUserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUserRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUserRepository)(nil).GetAll), ctx)
}

// GetByUsername mocks base method.
func (m *MockIUserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIUserRepositoryMockRecorder) GetByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIUserRepository)(nil).GetByUsername), ctx, username)
}

// IsVerified mocks base method.
func (m *MockIUserRepository) IsVerified(ctx context.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVerified", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsVerified indicates an expected call of IsVerified.
func (mr *MockIUserRepositoryMockRecorder) IsVerified(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVerified", reflect.TypeOf((*MockIUserRepository)(nil).IsVerified), ctx, username)
}

// Update mocks base method.
func (m *MockIUserRepository) Update(ctx context.Context, user model.User, someField string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user, someField)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepositoryMockRecorder) Update(ctx, user, someField interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), ctx, user, someField)
}

// Verify mocks base method.
func (m *MockIUserRepository) Verify(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockIUserRepositoryMockRecorder) Verify(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockIUserRepository)(nil).Verify), ctx, username)
}