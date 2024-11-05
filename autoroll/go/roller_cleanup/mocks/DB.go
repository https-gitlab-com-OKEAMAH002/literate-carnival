// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	roller_cleanup "go.skia.org/infra/autoroll/go/roller_cleanup"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// History provides a mock function with given fields: ctx, rollerID, limit
func (_m *DB) History(ctx context.Context, rollerID string, limit int) ([]*roller_cleanup.CleanupRequest, error) {
	ret := _m.Called(ctx, rollerID, limit)

	if len(ret) == 0 {
		panic("no return value specified for History")
	}

	var r0 []*roller_cleanup.CleanupRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) ([]*roller_cleanup.CleanupRequest, error)); ok {
		return rf(ctx, rollerID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []*roller_cleanup.CleanupRequest); ok {
		r0 = rf(ctx, rollerID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*roller_cleanup.CleanupRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, rollerID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestCleanup provides a mock function with given fields: ctx, req
func (_m *DB) RequestCleanup(ctx context.Context, req *roller_cleanup.CleanupRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for RequestCleanup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *roller_cleanup.CleanupRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}