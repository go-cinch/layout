// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: helloworld/v1/helloword.proto

package v1

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

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloworldClient interface {
	CreateGreeter(ctx context.Context, in *CreateGreeterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetGreeter(ctx context.Context, in *GetGreeterRequest, opts ...grpc.CallOption) (*GetGreeterReply, error)
	FindGreeter(ctx context.Context, in *FindGreeterRequest, opts ...grpc.CallOption) (*FindGreeterReply, error)
	UpdateGreeter(ctx context.Context, in *UpdateGreeterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteGreeter(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type helloworldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldClient(cc grpc.ClientConnInterface) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) CreateGreeter(ctx context.Context, in *CreateGreeterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Helloworld/CreateGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) GetGreeter(ctx context.Context, in *GetGreeterRequest, opts ...grpc.CallOption) (*GetGreeterReply, error) {
	out := new(GetGreeterReply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Helloworld/GetGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) FindGreeter(ctx context.Context, in *FindGreeterRequest, opts ...grpc.CallOption) (*FindGreeterReply, error) {
	out := new(FindGreeterReply)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Helloworld/FindGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) UpdateGreeter(ctx context.Context, in *UpdateGreeterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Helloworld/UpdateGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) DeleteGreeter(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/helloworld.v1.Helloworld/DeleteGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServer is the server API for Helloworld service.
// All implementations must embed UnimplementedHelloworldServer
// for forward compatibility
type HelloworldServer interface {
	CreateGreeter(context.Context, *CreateGreeterRequest) (*emptypb.Empty, error)
	GetGreeter(context.Context, *GetGreeterRequest) (*GetGreeterReply, error)
	FindGreeter(context.Context, *FindGreeterRequest) (*FindGreeterReply, error)
	UpdateGreeter(context.Context, *UpdateGreeterRequest) (*emptypb.Empty, error)
	DeleteGreeter(context.Context, *IdsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHelloworldServer()
}

// UnimplementedHelloworldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (UnimplementedHelloworldServer) CreateGreeter(context.Context, *CreateGreeterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGreeter not implemented")
}
func (UnimplementedHelloworldServer) GetGreeter(context.Context, *GetGreeterRequest) (*GetGreeterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeter not implemented")
}
func (UnimplementedHelloworldServer) FindGreeter(context.Context, *FindGreeterRequest) (*FindGreeterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGreeter not implemented")
}
func (UnimplementedHelloworldServer) UpdateGreeter(context.Context, *UpdateGreeterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGreeter not implemented")
}
func (UnimplementedHelloworldServer) DeleteGreeter(context.Context, *IdsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGreeter not implemented")
}
func (UnimplementedHelloworldServer) mustEmbedUnimplementedHelloworldServer() {}

// UnsafeHelloworldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloworldServer will
// result in compilation errors.
type UnsafeHelloworldServer interface {
	mustEmbedUnimplementedHelloworldServer()
}

func RegisterHelloworldServer(s grpc.ServiceRegistrar, srv HelloworldServer) {
	s.RegisterService(&Helloworld_ServiceDesc, srv)
}

func _Helloworld_CreateGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).CreateGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Helloworld/CreateGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).CreateGreeter(ctx, req.(*CreateGreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_GetGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).GetGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Helloworld/GetGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).GetGreeter(ctx, req.(*GetGreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_FindGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindGreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).FindGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Helloworld/FindGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).FindGreeter(ctx, req.(*FindGreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_UpdateGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).UpdateGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Helloworld/UpdateGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).UpdateGreeter(ctx, req.(*UpdateGreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_DeleteGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).DeleteGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.v1.Helloworld/DeleteGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).DeleteGreeter(ctx, req.(*IdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Helloworld_ServiceDesc is the grpc.ServiceDesc for Helloworld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Helloworld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGreeter",
			Handler:    _Helloworld_CreateGreeter_Handler,
		},
		{
			MethodName: "GetGreeter",
			Handler:    _Helloworld_GetGreeter_Handler,
		},
		{
			MethodName: "FindGreeter",
			Handler:    _Helloworld_FindGreeter_Handler,
		},
		{
			MethodName: "UpdateGreeter",
			Handler:    _Helloworld_UpdateGreeter_Handler,
		},
		{
			MethodName: "DeleteGreeter",
			Handler:    _Helloworld_DeleteGreeter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/helloword.proto",
}
