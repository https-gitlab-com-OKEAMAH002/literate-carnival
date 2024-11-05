// Code generated by mockery v0.0.0-dev. (with some manual edits)
//
// Note: Please be careful autogenerating this file again, as mockery doesn't do a perfect job
// of creating a mock grpc server - it doesn't add `v1.UnimplementedAnomalyGroupServiceServer` as
// part of `AnomalyGroupServiceServer`. In case you want to regenerate it, re-add it to
// AnomalyGroupServiceServer definition.
package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "go.skia.org/infra/perf/go/anomalygroup/proto/v1"
)

// AnomalyGroupServiceServer is an autogenerated mock type for the AnomalyGroupServiceServer type
type AnomalyGroupServiceServer struct {
	mock.Mock
	// -----------Line below was added manually-----------
	v1.UnimplementedAnomalyGroupServiceServer
}

// CreateNewAnomalyGroup provides a mock function with given fields: _a0, _a1
func (_m *AnomalyGroupServiceServer) CreateNewAnomalyGroup(_a0 context.Context, _a1 *v1.CreateNewAnomalyGroupRequest) (*v1.CreateNewAnomalyGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewAnomalyGroup")
	}

	var r0 *v1.CreateNewAnomalyGroupResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateNewAnomalyGroupRequest) (*v1.CreateNewAnomalyGroupResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateNewAnomalyGroupRequest) *v1.CreateNewAnomalyGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.CreateNewAnomalyGroupResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateNewAnomalyGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExistingGroups provides a mock function with given fields: _a0, _a1
func (_m *AnomalyGroupServiceServer) FindExistingGroups(_a0 context.Context, _a1 *v1.FindExistingGroupsRequest) (*v1.FindExistingGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindExistingGroups")
	}

	var r0 *v1.FindExistingGroupsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.FindExistingGroupsRequest) (*v1.FindExistingGroupsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.FindExistingGroupsRequest) *v1.FindExistingGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.FindExistingGroupsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.FindExistingGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTopAnomalies provides a mock function with given fields: _a0, _a1
func (_m *AnomalyGroupServiceServer) FindTopAnomalies(_a0 context.Context, _a1 *v1.FindTopAnomaliesRequest) (*v1.FindTopAnomaliesResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindTopAnomalies")
	}

	var r0 *v1.FindTopAnomaliesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.FindTopAnomaliesRequest) (*v1.FindTopAnomaliesResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.FindTopAnomaliesRequest) *v1.FindTopAnomaliesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.FindTopAnomaliesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.FindTopAnomaliesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadAnomalyGroupByID provides a mock function with given fields: _a0, _a1
func (_m *AnomalyGroupServiceServer) LoadAnomalyGroupByID(_a0 context.Context, _a1 *v1.LoadAnomalyGroupByIDRequest) (*v1.LoadAnomalyGroupByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LoadAnomalyGroupByID")
	}

	var r0 *v1.LoadAnomalyGroupByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LoadAnomalyGroupByIDRequest) (*v1.LoadAnomalyGroupByIDResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LoadAnomalyGroupByIDRequest) *v1.LoadAnomalyGroupByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.LoadAnomalyGroupByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.LoadAnomalyGroupByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAnomalyGroup provides a mock function with given fields: _a0, _a1
func (_m *AnomalyGroupServiceServer) UpdateAnomalyGroup(_a0 context.Context, _a1 *v1.UpdateAnomalyGroupRequest) (*v1.UpdateAnomalyGroupResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAnomalyGroup")
	}

	var r0 *v1.UpdateAnomalyGroupResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.UpdateAnomalyGroupRequest) (*v1.UpdateAnomalyGroupResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.UpdateAnomalyGroupRequest) *v1.UpdateAnomalyGroupResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.UpdateAnomalyGroupResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.UpdateAnomalyGroupRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedAnomalyGroupServiceServer provides a mock function with given fields:
func (_m *AnomalyGroupServiceServer) mustEmbedUnimplementedAnomalyGroupServiceServer() {
	_m.Called()
}

// NewAnomalyGroupServiceServer creates a new instance of AnomalyGroupServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAnomalyGroupServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *AnomalyGroupServiceServer {
	mock := &AnomalyGroupServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
