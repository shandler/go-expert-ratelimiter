// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source repository.go -destination mock/repository_mock.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/shandler/go-expert-ratelimiter/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockcommonRepo is a mock of commonRepo interface.
type MockcommonRepo struct {
	ctrl     *gomock.Controller
	recorder *MockcommonRepoMockRecorder
}

// MockcommonRepoMockRecorder is the mock recorder for MockcommonRepo.
type MockcommonRepoMockRecorder struct {
	mock *MockcommonRepo
}

// NewMockcommonRepo creates a new mock instance.
func NewMockcommonRepo(ctrl *gomock.Controller) *MockcommonRepo {
	mock := &MockcommonRepo{ctrl: ctrl}
	mock.recorder = &MockcommonRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcommonRepo) EXPECT() *MockcommonRepoMockRecorder {
	return m.recorder
}

// GetBlockedDuration mocks base method.
func (m *MockcommonRepo) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockcommonRepoMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockcommonRepo)(nil).GetBlockedDuration), ctx, key)
}

// GetRequest mocks base method.
func (m *MockcommonRepo) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", ctx, key)
	ret0, _ := ret[0].(*entity.RateLimiter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockcommonRepoMockRecorder) GetRequest(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockcommonRepo)(nil).GetRequest), ctx, key)
}

// SaveBlockedDuration mocks base method.
func (m *MockcommonRepo) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockcommonRepoMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockcommonRepo)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockcommonRepo) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockcommonRepoMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockcommonRepo)(nil).UpsertRequest), ctx, key, rl)
}

// MockAPIKeyRepository is a mock of APIKeyRepository interface.
type MockAPIKeyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyRepositoryMockRecorder
}

// MockAPIKeyRepositoryMockRecorder is the mock recorder for MockAPIKeyRepository.
type MockAPIKeyRepositoryMockRecorder struct {
	mock *MockAPIKeyRepository
}

// NewMockAPIKeyRepository creates a new mock instance.
func NewMockAPIKeyRepository(ctrl *gomock.Controller) *MockAPIKeyRepository {
	mock := &MockAPIKeyRepository{ctrl: ctrl}
	mock.recorder = &MockAPIKeyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyRepository) EXPECT() *MockAPIKeyRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAPIKeyRepository) Get(ctx context.Context, value string) (*entity.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, value)
	ret0, _ := ret[0].(*entity.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAPIKeyRepositoryMockRecorder) Get(ctx, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAPIKeyRepository)(nil).Get), ctx, value)
}

// GetBlockedDuration mocks base method.
func (m *MockAPIKeyRepository) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockAPIKeyRepositoryMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockAPIKeyRepository)(nil).GetBlockedDuration), ctx, key)
}

// GetRequest mocks base method.
func (m *MockAPIKeyRepository) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", ctx, key)
	ret0, _ := ret[0].(*entity.RateLimiter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockAPIKeyRepositoryMockRecorder) GetRequest(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockAPIKeyRepository)(nil).GetRequest), ctx, key)
}

// Save mocks base method.
func (m *MockAPIKeyRepository) Save(ctx context.Context, key *entity.APIKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockAPIKeyRepositoryMockRecorder) Save(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAPIKeyRepository)(nil).Save), ctx, key)
}

// SaveBlockedDuration mocks base method.
func (m *MockAPIKeyRepository) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockAPIKeyRepositoryMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockAPIKeyRepository)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockAPIKeyRepository) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockAPIKeyRepositoryMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockAPIKeyRepository)(nil).UpsertRequest), ctx, key, rl)
}

// MockIPRepository is a mock of IPRepository interface.
type MockIPRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPRepositoryMockRecorder
}

// MockIPRepositoryMockRecorder is the mock recorder for MockIPRepository.
type MockIPRepositoryMockRecorder struct {
	mock *MockIPRepository
}

// NewMockIPRepository creates a new mock instance.
func NewMockIPRepository(ctrl *gomock.Controller) *MockIPRepository {
	mock := &MockIPRepository{ctrl: ctrl}
	mock.recorder = &MockIPRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPRepository) EXPECT() *MockIPRepositoryMockRecorder {
	return m.recorder
}

// GetBlockedDuration mocks base method.
func (m *MockIPRepository) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockIPRepositoryMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockIPRepository)(nil).GetBlockedDuration), ctx, key)
}

// GetRequest mocks base method.
func (m *MockIPRepository) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", ctx, key)
	ret0, _ := ret[0].(*entity.RateLimiter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockIPRepositoryMockRecorder) GetRequest(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockIPRepository)(nil).GetRequest), ctx, key)
}

// SaveBlockedDuration mocks base method.
func (m *MockIPRepository) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockIPRepositoryMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockIPRepository)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockIPRepository) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockIPRepositoryMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockIPRepository)(nil).UpsertRequest), ctx, key, rl)
}
