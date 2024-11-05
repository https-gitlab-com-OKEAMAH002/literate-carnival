// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	alerts "go.skia.org/infra/perf/go/alerts"

	mock "github.com/stretchr/testify/mock"
)

// Transport is an autogenerated mock type for the Transport type
type Transport struct {
	mock.Mock
}

// SendNewRegression provides a mock function with given fields: ctx, alert, body, subject
func (_m *Transport) SendNewRegression(ctx context.Context, alert *alerts.Alert, body string, subject string) (string, error) {
	ret := _m.Called(ctx, alert, body, subject)

	if len(ret) == 0 {
		panic("no return value specified for SendNewRegression")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *alerts.Alert, string, string) (string, error)); ok {
		return rf(ctx, alert, body, subject)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *alerts.Alert, string, string) string); ok {
		r0 = rf(ctx, alert, body, subject)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *alerts.Alert, string, string) error); ok {
		r1 = rf(ctx, alert, body, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRegressionMissing provides a mock function with given fields: ctx, threadingReference, alert, body, subject
func (_m *Transport) SendRegressionMissing(ctx context.Context, threadingReference string, alert *alerts.Alert, body string, subject string) error {
	ret := _m.Called(ctx, threadingReference, alert, body, subject)

	if len(ret) == 0 {
		panic("no return value specified for SendRegressionMissing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *alerts.Alert, string, string) error); ok {
		r0 = rf(ctx, threadingReference, alert, body, subject)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTransport creates a new instance of Transport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransport(t interface {
	mock.TestingT
	Cleanup(func())
}) *Transport {
	mock := &Transport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
