// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"

	types "go.skia.org/infra/npm-audit-mirror/go/types"
)

// NpmDB is an autogenerated mock type for the NpmDB type
type NpmDB struct {
	mock.Mock
}

// GetFromDB provides a mock function with given fields: ctx, key
func (_m *NpmDB) GetFromDB(ctx context.Context, key string) (*types.NpmAuditData, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetFromDB")
	}

	var r0 *types.NpmAuditData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*types.NpmAuditData, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NpmAuditData); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NpmAuditData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutInDB provides a mock function with given fields: ctx, key, issueId, created
func (_m *NpmDB) PutInDB(ctx context.Context, key string, issueId int64, created time.Time) error {
	ret := _m.Called(ctx, key, issueId, created)

	if len(ret) == 0 {
		panic("no return value specified for PutInDB")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, time.Time) error); ok {
		r0 = rf(ctx, key, issueId, created)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewNpmDB creates a new instance of NpmDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNpmDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *NpmDB {
	mock := &NpmDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}