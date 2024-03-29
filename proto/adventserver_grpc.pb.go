// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AdventServerServiceClient is the client API for AdventServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdventServerServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type adventServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdventServerServiceClient(cc grpc.ClientConnInterface) AdventServerServiceClient {
	return &adventServerServiceClient{cc}
}

func (c *adventServerServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, "/adventserver.AdventServerService/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdventServerServiceServer is the server API for AdventServerService service.
// All implementations should embed UnimplementedAdventServerServiceServer
// for forward compatibility
type AdventServerServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

// UnimplementedAdventServerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdventServerServiceServer struct {
}

func (UnimplementedAdventServerServiceServer) Solve(context.Context, *SolveRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

// UnsafeAdventServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdventServerServiceServer will
// result in compilation errors.
type UnsafeAdventServerServiceServer interface {
	mustEmbedUnimplementedAdventServerServiceServer()
}

func RegisterAdventServerServiceServer(s grpc.ServiceRegistrar, srv AdventServerServiceServer) {
	s.RegisterService(&_AdventServerService_serviceDesc, srv)
}

func _AdventServerService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdventServerServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adventserver.AdventServerService/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdventServerServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdventServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adventserver.AdventServerService",
	HandlerType: (*AdventServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _AdventServerService_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adventserver.proto",
}
