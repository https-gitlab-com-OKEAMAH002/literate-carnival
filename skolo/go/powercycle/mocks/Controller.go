// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	powercycle "go.skia.org/infra/skolo/go/powercycle"

	time "time"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeviceIDs provides a mock function with given fields:
func (_m *Controller) DeviceIDs() []powercycle.DeviceID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeviceIDs")
	}

	var r0 []powercycle.DeviceID
	if rf, ok := ret.Get(0).(func() []powercycle.DeviceID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]powercycle.DeviceID)
		}
	}

	return r0
}

// PowerCycle provides a mock function with given fields: ctx, id, delayOverride
func (_m *Controller) PowerCycle(ctx context.Context, id powercycle.DeviceID, delayOverride time.Duration) error {
	ret := _m.Called(ctx, id, delayOverride)

	if len(ret) == 0 {
		panic("no return value specified for PowerCycle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, powercycle.DeviceID, time.Duration) error); ok {
		r0 = rf(ctx, id, delayOverride)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
