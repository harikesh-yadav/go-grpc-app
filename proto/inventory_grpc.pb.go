// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/inventory.proto

package go_grpc_app

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetItemByName(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Items, error)
	GetItemById(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Item, error)
	AddItem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetItemByName(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Items, error) {
	out := new(Items)
	err := c.cc.Invoke(ctx, "/grpc_service.InventoryService/getItemByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetItemById(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/grpc_service.InventoryService/getItemById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) AddItem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/grpc_service.InventoryService/addItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	GetItemByName(context.Context, *wrapperspb.StringValue) (*Items, error)
	GetItemById(context.Context, *wrapperspb.StringValue) (*Item, error)
	AddItem(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) GetItemByName(context.Context, *wrapperspb.StringValue) (*Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemByName not implemented")
}
func (UnimplementedInventoryServiceServer) GetItemById(context.Context, *wrapperspb.StringValue) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedInventoryServiceServer) AddItem(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_GetItemByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetItemByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.InventoryService/getItemByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetItemByName(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.InventoryService/getItemById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetItemById(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.InventoryService/addItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).AddItem(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_service.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getItemByName",
			Handler:    _InventoryService_GetItemByName_Handler,
		},
		{
			MethodName: "getItemById",
			Handler:    _InventoryService_GetItemById_Handler,
		},
		{
			MethodName: "addItem",
			Handler:    _InventoryService_AddItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/inventory.proto",
}
