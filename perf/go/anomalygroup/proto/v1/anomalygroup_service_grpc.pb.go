// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: anomalygroup_service.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AnomalyGroupService_CreateNewAnomalyGroup_FullMethodName = "/anomalygroup.v1.AnomalyGroupService/CreateNewAnomalyGroup"
	AnomalyGroupService_LoadAnomalyGroupByID_FullMethodName  = "/anomalygroup.v1.AnomalyGroupService/LoadAnomalyGroupByID"
	AnomalyGroupService_UpdateAnomalyGroup_FullMethodName    = "/anomalygroup.v1.AnomalyGroupService/UpdateAnomalyGroup"
	AnomalyGroupService_FindExistingGroups_FullMethodName    = "/anomalygroup.v1.AnomalyGroupService/FindExistingGroups"
	AnomalyGroupService_FindTopAnomalies_FullMethodName      = "/anomalygroup.v1.AnomalyGroupService/FindTopAnomalies"
)

// AnomalyGroupServiceClient is the client API for AnomalyGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnomalyGroupServiceClient interface {
	// Create a new anomaly group based on a set of criterias.
	// Avoid binding it to a specific regression.
	CreateNewAnomalyGroup(ctx context.Context, in *CreateNewAnomalyGroupRequest, opts ...grpc.CallOption) (*CreateNewAnomalyGroupResponse, error)
	// Read info for an anomaly group.
	LoadAnomalyGroupByID(ctx context.Context, in *LoadAnomalyGroupByIDRequest, opts ...grpc.CallOption) (*LoadAnomalyGroupByIDResponse, error)
	// Update a given anomaly group.
	UpdateAnomalyGroup(ctx context.Context, in *UpdateAnomalyGroupRequest, opts ...grpc.CallOption) (*UpdateAnomalyGroupResponse, error)
	// Find matching anomaly groups based on the criterias.
	// (e.g., from a newly found anomaly).
	FindExistingGroups(ctx context.Context, in *FindExistingGroupsRequest, opts ...grpc.CallOption) (*FindExistingGroupsResponse, error)
	FindTopAnomalies(ctx context.Context, in *FindTopAnomaliesRequest, opts ...grpc.CallOption) (*FindTopAnomaliesResponse, error)
}

type anomalyGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnomalyGroupServiceClient(cc grpc.ClientConnInterface) AnomalyGroupServiceClient {
	return &anomalyGroupServiceClient{cc}
}

func (c *anomalyGroupServiceClient) CreateNewAnomalyGroup(ctx context.Context, in *CreateNewAnomalyGroupRequest, opts ...grpc.CallOption) (*CreateNewAnomalyGroupResponse, error) {
	out := new(CreateNewAnomalyGroupResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_CreateNewAnomalyGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) LoadAnomalyGroupByID(ctx context.Context, in *LoadAnomalyGroupByIDRequest, opts ...grpc.CallOption) (*LoadAnomalyGroupByIDResponse, error) {
	out := new(LoadAnomalyGroupByIDResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_LoadAnomalyGroupByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) UpdateAnomalyGroup(ctx context.Context, in *UpdateAnomalyGroupRequest, opts ...grpc.CallOption) (*UpdateAnomalyGroupResponse, error) {
	out := new(UpdateAnomalyGroupResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_UpdateAnomalyGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) FindExistingGroups(ctx context.Context, in *FindExistingGroupsRequest, opts ...grpc.CallOption) (*FindExistingGroupsResponse, error) {
	out := new(FindExistingGroupsResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_FindExistingGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyGroupServiceClient) FindTopAnomalies(ctx context.Context, in *FindTopAnomaliesRequest, opts ...grpc.CallOption) (*FindTopAnomaliesResponse, error) {
	out := new(FindTopAnomaliesResponse)
	err := c.cc.Invoke(ctx, AnomalyGroupService_FindTopAnomalies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnomalyGroupServiceServer is the server API for AnomalyGroupService service.
// All implementations must embed UnimplementedAnomalyGroupServiceServer
// for forward compatibility
type AnomalyGroupServiceServer interface {
	// Create a new anomaly group based on a set of criterias.
	// Avoid binding it to a specific regression.
	CreateNewAnomalyGroup(context.Context, *CreateNewAnomalyGroupRequest) (*CreateNewAnomalyGroupResponse, error)
	// Read info for an anomaly group.
	LoadAnomalyGroupByID(context.Context, *LoadAnomalyGroupByIDRequest) (*LoadAnomalyGroupByIDResponse, error)
	// Update a given anomaly group.
	UpdateAnomalyGroup(context.Context, *UpdateAnomalyGroupRequest) (*UpdateAnomalyGroupResponse, error)
	// Find matching anomaly groups based on the criterias.
	// (e.g., from a newly found anomaly).
	FindExistingGroups(context.Context, *FindExistingGroupsRequest) (*FindExistingGroupsResponse, error)
	FindTopAnomalies(context.Context, *FindTopAnomaliesRequest) (*FindTopAnomaliesResponse, error)
	mustEmbedUnimplementedAnomalyGroupServiceServer()
}

// UnimplementedAnomalyGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnomalyGroupServiceServer struct {
}

func (UnimplementedAnomalyGroupServiceServer) CreateNewAnomalyGroup(context.Context, *CreateNewAnomalyGroupRequest) (*CreateNewAnomalyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewAnomalyGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) LoadAnomalyGroupByID(context.Context, *LoadAnomalyGroupByIDRequest) (*LoadAnomalyGroupByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadAnomalyGroupByID not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) UpdateAnomalyGroup(context.Context, *UpdateAnomalyGroupRequest) (*UpdateAnomalyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnomalyGroup not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) FindExistingGroups(context.Context, *FindExistingGroupsRequest) (*FindExistingGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindExistingGroups not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) FindTopAnomalies(context.Context, *FindTopAnomaliesRequest) (*FindTopAnomaliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTopAnomalies not implemented")
}
func (UnimplementedAnomalyGroupServiceServer) mustEmbedUnimplementedAnomalyGroupServiceServer() {}

// UnsafeAnomalyGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnomalyGroupServiceServer will
// result in compilation errors.
type UnsafeAnomalyGroupServiceServer interface {
	mustEmbedUnimplementedAnomalyGroupServiceServer()
}

func RegisterAnomalyGroupServiceServer(s grpc.ServiceRegistrar, srv AnomalyGroupServiceServer) {
	s.RegisterService(&AnomalyGroupService_ServiceDesc, srv)
}

func _AnomalyGroupService_CreateNewAnomalyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewAnomalyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).CreateNewAnomalyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_CreateNewAnomalyGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).CreateNewAnomalyGroup(ctx, req.(*CreateNewAnomalyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_LoadAnomalyGroupByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadAnomalyGroupByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).LoadAnomalyGroupByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_LoadAnomalyGroupByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).LoadAnomalyGroupByID(ctx, req.(*LoadAnomalyGroupByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_UpdateAnomalyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnomalyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).UpdateAnomalyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_UpdateAnomalyGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).UpdateAnomalyGroup(ctx, req.(*UpdateAnomalyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_FindExistingGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindExistingGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).FindExistingGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_FindExistingGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).FindExistingGroups(ctx, req.(*FindExistingGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnomalyGroupService_FindTopAnomalies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTopAnomaliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyGroupServiceServer).FindTopAnomalies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnomalyGroupService_FindTopAnomalies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyGroupServiceServer).FindTopAnomalies(ctx, req.(*FindTopAnomaliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnomalyGroupService_ServiceDesc is the grpc.ServiceDesc for AnomalyGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnomalyGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anomalygroup.v1.AnomalyGroupService",
	HandlerType: (*AnomalyGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewAnomalyGroup",
			Handler:    _AnomalyGroupService_CreateNewAnomalyGroup_Handler,
		},
		{
			MethodName: "LoadAnomalyGroupByID",
			Handler:    _AnomalyGroupService_LoadAnomalyGroupByID_Handler,
		},
		{
			MethodName: "UpdateAnomalyGroup",
			Handler:    _AnomalyGroupService_UpdateAnomalyGroup_Handler,
		},
		{
			MethodName: "FindExistingGroups",
			Handler:    _AnomalyGroupService_FindExistingGroups_Handler,
		},
		{
			MethodName: "FindTopAnomalies",
			Handler:    _AnomalyGroupService_FindTopAnomalies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anomalygroup_service.proto",
}
