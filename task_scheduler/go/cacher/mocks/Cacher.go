// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	specs "go.skia.org/infra/task_scheduler/go/specs"

	types "go.skia.org/infra/task_scheduler/go/types"
)

// Cacher is an autogenerated mock type for the Cacher type
type Cacher struct {
	mock.Mock
}

// GetOrCacheRepoState provides a mock function with given fields: ctx, rs
func (_m *Cacher) GetOrCacheRepoState(ctx context.Context, rs types.RepoState) (*specs.TasksCfg, error) {
	ret := _m.Called(ctx, rs)

	if len(ret) == 0 {
		panic("no return value specified for GetOrCacheRepoState")
	}

	var r0 *specs.TasksCfg
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RepoState) (*specs.TasksCfg, error)); ok {
		return rf(ctx, rs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RepoState) *specs.TasksCfg); ok {
		r0 = rf(ctx, rs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*specs.TasksCfg)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RepoState) error); ok {
		r1 = rf(ctx, rs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCacher creates a new instance of Cacher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCacher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cacher {
	mock := &Cacher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
