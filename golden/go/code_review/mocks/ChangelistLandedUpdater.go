// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	vcsinfo "go.skia.org/infra/go/vcsinfo"
)

// ChangelistLandedUpdater is an autogenerated mock type for the ChangelistLandedUpdater type
type ChangelistLandedUpdater struct {
	mock.Mock
}

// UpdateChangelistsAsLanded provides a mock function with given fields: ctx, commits
func (_m *ChangelistLandedUpdater) UpdateChangelistsAsLanded(ctx context.Context, commits []*vcsinfo.LongCommit) error {
	ret := _m.Called(ctx, commits)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChangelistsAsLanded")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*vcsinfo.LongCommit) error); ok {
		r0 = rf(ctx, commits)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChangelistLandedUpdater creates a new instance of ChangelistLandedUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChangelistLandedUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChangelistLandedUpdater {
	mock := &ChangelistLandedUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}