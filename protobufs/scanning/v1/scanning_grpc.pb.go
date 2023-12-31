// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: scanning.proto

package scanningpb

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

// ScanningServiceClient is the client API for ScanningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScanningServiceClient interface {
	ScanningText(ctx context.Context, in *ScanningTextRequest, opts ...grpc.CallOption) (*ScanningTextResponse, error)
}

type scanningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScanningServiceClient(cc grpc.ClientConnInterface) ScanningServiceClient {
	return &scanningServiceClient{cc}
}

func (c *scanningServiceClient) ScanningText(ctx context.Context, in *ScanningTextRequest, opts ...grpc.CallOption) (*ScanningTextResponse, error) {
	out := new(ScanningTextResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.scanning.v1.ScanningService/ScanningText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanningServiceServer is the server API for ScanningService service.
// All implementations must embed UnimplementedScanningServiceServer
// for forward compatibility
type ScanningServiceServer interface {
	ScanningText(context.Context, *ScanningTextRequest) (*ScanningTextResponse, error)
	mustEmbedUnimplementedScanningServiceServer()
}

// UnimplementedScanningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScanningServiceServer struct {
}

func (UnimplementedScanningServiceServer) ScanningText(context.Context, *ScanningTextRequest) (*ScanningTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanningText not implemented")
}
func (UnimplementedScanningServiceServer) mustEmbedUnimplementedScanningServiceServer() {}

// UnsafeScanningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScanningServiceServer will
// result in compilation errors.
type UnsafeScanningServiceServer interface {
	mustEmbedUnimplementedScanningServiceServer()
}

func RegisterScanningServiceServer(s grpc.ServiceRegistrar, srv ScanningServiceServer) {
	s.RegisterService(&ScanningService_ServiceDesc, srv)
}

func _ScanningService_ScanningText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanningTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanningServiceServer).ScanningText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.scanning.v1.ScanningService/ScanningText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanningServiceServer).ScanningText(ctx, req.(*ScanningTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScanningService_ServiceDesc is the grpc.ServiceDesc for ScanningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScanningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.scanning.v1.ScanningService",
	HandlerType: (*ScanningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanningText",
			Handler:    _ScanningService_ScanningText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanning.proto",
}
