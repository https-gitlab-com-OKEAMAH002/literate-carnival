// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	protov1 "go.skia.org/infra/perf/go/subscription/proto/v1"

	v1 "go.skia.org/infra/perf/go/culprit/proto/v1"
)

// Formatter is an autogenerated mock type for the Formatter type
type Formatter struct {
	mock.Mock
}

// GetSubjectAndBody provides a mock function with given fields: ctx, culprit, subscription
func (_m *Formatter) GetSubjectAndBody(ctx context.Context, culprit *v1.Culprit, subscription *protov1.Subscription) (string, string, error) {
	ret := _m.Called(ctx, culprit, subscription)

	if len(ret) == 0 {
		panic("no return value specified for GetSubjectAndBody")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Culprit, *protov1.Subscription) (string, string, error)); ok {
		return rf(ctx, culprit, subscription)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Culprit, *protov1.Subscription) string); ok {
		r0 = rf(ctx, culprit, subscription)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.Culprit, *protov1.Subscription) string); ok {
		r1 = rf(ctx, culprit, subscription)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *v1.Culprit, *protov1.Subscription) error); ok {
		r2 = rf(ctx, culprit, subscription)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewFormatter creates a new instance of Formatter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFormatter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Formatter {
	mock := &Formatter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
