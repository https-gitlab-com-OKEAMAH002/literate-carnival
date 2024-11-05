// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	scrap "go.skia.org/infra/scrap/go/scrap"
)

// ScrapExchange is an autogenerated mock type for the ScrapExchange type
type ScrapExchange struct {
	mock.Mock
}

// CreateScrap provides a mock function with given fields: ctx, _a1
func (_m *ScrapExchange) CreateScrap(ctx context.Context, _a1 scrap.ScrapBody) (scrap.ScrapID, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateScrap")
	}

	var r0 scrap.ScrapID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.ScrapBody) (scrap.ScrapID, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, scrap.ScrapBody) scrap.ScrapID); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(scrap.ScrapID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, scrap.ScrapBody) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteName provides a mock function with given fields: ctx, t, name
func (_m *ScrapExchange) DeleteName(ctx context.Context, t scrap.Type, name string) error {
	ret := _m.Called(ctx, t, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) error); ok {
		r0 = rf(ctx, t, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteScrap provides a mock function with given fields: ctx, t, hashOrName
func (_m *ScrapExchange) DeleteScrap(ctx context.Context, t scrap.Type, hashOrName string) error {
	ret := _m.Called(ctx, t, hashOrName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScrap")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) error); ok {
		r0 = rf(ctx, t, hashOrName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Expand provides a mock function with given fields: ctx, t, hashOrName, lang, w
func (_m *ScrapExchange) Expand(ctx context.Context, t scrap.Type, hashOrName string, lang scrap.Lang, w io.Writer) error {
	ret := _m.Called(ctx, t, hashOrName, lang, w)

	if len(ret) == 0 {
		panic("no return value specified for Expand")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string, scrap.Lang, io.Writer) error); ok {
		r0 = rf(ctx, t, hashOrName, lang, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetName provides a mock function with given fields: ctx, t, name
func (_m *ScrapExchange) GetName(ctx context.Context, t scrap.Type, name string) (scrap.Name, error) {
	ret := _m.Called(ctx, t, name)

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 scrap.Name
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) (scrap.Name, error)); ok {
		return rf(ctx, t, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) scrap.Name); ok {
		r0 = rf(ctx, t, name)
	} else {
		r0 = ret.Get(0).(scrap.Name)
	}

	if rf, ok := ret.Get(1).(func(context.Context, scrap.Type, string) error); ok {
		r1 = rf(ctx, t, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNames provides a mock function with given fields: ctx, t
func (_m *ScrapExchange) ListNames(ctx context.Context, t scrap.Type) ([]string, error) {
	ret := _m.Called(ctx, t)

	if len(ret) == 0 {
		panic("no return value specified for ListNames")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type) ([]string, error)); ok {
		return rf(ctx, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type) []string); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, scrap.Type) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadScrap provides a mock function with given fields: ctx, t, hashOrName
func (_m *ScrapExchange) LoadScrap(ctx context.Context, t scrap.Type, hashOrName string) (scrap.ScrapBody, error) {
	ret := _m.Called(ctx, t, hashOrName)

	if len(ret) == 0 {
		panic("no return value specified for LoadScrap")
	}

	var r0 scrap.ScrapBody
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) (scrap.ScrapBody, error)); ok {
		return rf(ctx, t, hashOrName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string) scrap.ScrapBody); ok {
		r0 = rf(ctx, t, hashOrName)
	} else {
		r0 = ret.Get(0).(scrap.ScrapBody)
	}

	if rf, ok := ret.Get(1).(func(context.Context, scrap.Type, string) error); ok {
		r1 = rf(ctx, t, hashOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutName provides a mock function with given fields: ctx, t, name, nameBody
func (_m *ScrapExchange) PutName(ctx context.Context, t scrap.Type, name string, nameBody scrap.Name) error {
	ret := _m.Called(ctx, t, name, nameBody)

	if len(ret) == 0 {
		panic("no return value specified for PutName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, scrap.Type, string, scrap.Name) error); ok {
		r0 = rf(ctx, t, name, nameBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewScrapExchange creates a new instance of ScrapExchange. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScrapExchange(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScrapExchange {
	mock := &ScrapExchange{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
