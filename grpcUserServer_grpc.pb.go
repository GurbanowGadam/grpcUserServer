// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpcUserServer.proto

package grpcUserServer

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

// UserServerServiceClient is the client API for UserServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServerServiceClient interface {
	GetProfileData(ctx context.Context, in *GetProfileDataRequest, opts ...grpc.CallOption) (*GetProfileDataResponse, error)
	GetBlockedLists(ctx context.Context, in *GetBlockedListsRequest, opts ...grpc.CallOption) (*GetBlockedListsResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type userServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServerServiceClient(cc grpc.ClientConnInterface) UserServerServiceClient {
	return &userServerServiceClient{cc}
}

func (c *userServerServiceClient) GetProfileData(ctx context.Context, in *GetProfileDataRequest, opts ...grpc.CallOption) (*GetProfileDataResponse, error) {
	out := new(GetProfileDataResponse)
	err := c.cc.Invoke(ctx, "/grpcUserServer.userServerService/GetProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerServiceClient) GetBlockedLists(ctx context.Context, in *GetBlockedListsRequest, opts ...grpc.CallOption) (*GetBlockedListsResponse, error) {
	out := new(GetBlockedListsResponse)
	err := c.cc.Invoke(ctx, "/grpcUserServer.userServerService/GetBlockedLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/grpcUserServer.userServerService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServiceServer is the server API for UserServerService service.
// All implementations must embed UnimplementedUserServerServiceServer
// for forward compatibility
type UserServerServiceServer interface {
	GetProfileData(context.Context, *GetProfileDataRequest) (*GetProfileDataResponse, error)
	GetBlockedLists(context.Context, *GetBlockedListsRequest) (*GetBlockedListsResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	mustEmbedUnimplementedUserServerServiceServer()
}

// UnimplementedUserServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServerServiceServer struct {
}

func (UnimplementedUserServerServiceServer) GetProfileData(context.Context, *GetProfileDataRequest) (*GetProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileData not implemented")
}
func (UnimplementedUserServerServiceServer) GetBlockedLists(context.Context, *GetBlockedListsRequest) (*GetBlockedListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockedLists not implemented")
}
func (UnimplementedUserServerServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServerServiceServer) mustEmbedUnimplementedUserServerServiceServer() {}

// UnsafeUserServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServerServiceServer will
// result in compilation errors.
type UnsafeUserServerServiceServer interface {
	mustEmbedUnimplementedUserServerServiceServer()
}

func RegisterUserServerServiceServer(s grpc.ServiceRegistrar, srv UserServerServiceServer) {
	s.RegisterService(&UserServerService_ServiceDesc, srv)
}

func _UserServerService_GetProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServiceServer).GetProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcUserServer.userServerService/GetProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServiceServer).GetProfileData(ctx, req.(*GetProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServerService_GetBlockedLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockedListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServiceServer).GetBlockedLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcUserServer.userServerService/GetBlockedLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServiceServer).GetBlockedLists(ctx, req.(*GetBlockedListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServerService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcUserServer.userServerService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServerService_ServiceDesc is the grpc.ServiceDesc for UserServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcUserServer.userServerService",
	HandlerType: (*UserServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileData",
			Handler:    _UserServerService_GetProfileData_Handler,
		},
		{
			MethodName: "GetBlockedLists",
			Handler:    _UserServerService_GetBlockedLists_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserServerService_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcUserServer.proto",
}
