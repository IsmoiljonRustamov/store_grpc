// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: todo.proto

package store

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	CreateStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*Store, error)
	GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*Store, error)
	UpdateStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) CreateStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/genprot.StoreService/CreateStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/genprot.StoreService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) UpdateStore(ctx context.Context, in *Store, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genprot.StoreService/UpdateStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) DeleteStore(ctx context.Context, in *GetStoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genprot.StoreService/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility
type StoreServiceServer interface {
	CreateStore(context.Context, *Store) (*Store, error)
	GetStore(context.Context, *GetStoreRequest) (*Store, error)
	UpdateStore(context.Context, *Store) (*emptypb.Empty, error)
	DeleteStore(context.Context, *GetStoreRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (UnimplementedStoreServiceServer) CreateStore(context.Context, *Store) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedStoreServiceServer) GetStore(context.Context, *GetStoreRequest) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}
func (UnimplementedStoreServiceServer) UpdateStore(context.Context, *Store) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStore not implemented")
}
func (UnimplementedStoreServiceServer) DeleteStore(context.Context, *GetStoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Store)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genprot.StoreService/CreateStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CreateStore(ctx, req.(*Store))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genprot.StoreService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_UpdateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Store)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).UpdateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genprot.StoreService/UpdateStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).UpdateStore(ctx, req.(*Store))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genprot.StoreService/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).DeleteStore(ctx, req.(*GetStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genprot.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _StoreService_CreateStore_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _StoreService_GetStore_Handler,
		},
		{
			MethodName: "UpdateStore",
			Handler:    _StoreService_UpdateStore_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _StoreService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
