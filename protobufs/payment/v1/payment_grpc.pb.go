// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: payment.proto

package paymentpb

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
	PaymentsService_CreateOrder_FullMethodName             = "/chatfinbot.payment.v1.PaymentsService/CreateOrder"
	PaymentsService_ListOrders_FullMethodName              = "/chatfinbot.payment.v1.PaymentsService/ListOrders"
	PaymentsService_GetOrderStatus_FullMethodName          = "/chatfinbot.payment.v1.PaymentsService/GetOrderStatus"
	PaymentsService_GetPlanDetail_FullMethodName           = "/chatfinbot.payment.v1.PaymentsService/GetPlanDetail"
	PaymentsService_GetUserUsagePermissions_FullMethodName = "/chatfinbot.payment.v1.PaymentsService/GetUserUsagePermissions"
)

// PaymentsServiceClient is the client API for PaymentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error)
	GetPlanDetail(ctx context.Context, in *GetPlanDetailRequest, opts ...grpc.CallOption) (*GetPlanDetailResponse, error)
	GetUserUsagePermissions(ctx context.Context, in *GetUserUsagePermissionsRequest, opts ...grpc.CallOption) (*GetUserUsagePermissionsResponse, error)
}

type paymentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsServiceClient(cc grpc.ClientConnInterface) PaymentsServiceClient {
	return &paymentsServiceClient{cc}
}

func (c *paymentsServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, PaymentsService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, PaymentsService_ListOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsServiceClient) GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error) {
	out := new(GetOrderStatusResponse)
	err := c.cc.Invoke(ctx, PaymentsService_GetOrderStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsServiceClient) GetPlanDetail(ctx context.Context, in *GetPlanDetailRequest, opts ...grpc.CallOption) (*GetPlanDetailResponse, error) {
	out := new(GetPlanDetailResponse)
	err := c.cc.Invoke(ctx, PaymentsService_GetPlanDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsServiceClient) GetUserUsagePermissions(ctx context.Context, in *GetUserUsagePermissionsRequest, opts ...grpc.CallOption) (*GetUserUsagePermissionsResponse, error) {
	out := new(GetUserUsagePermissionsResponse)
	err := c.cc.Invoke(ctx, PaymentsService_GetUserUsagePermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServiceServer is the server API for PaymentsService service.
// All implementations must embed UnimplementedPaymentsServiceServer
// for forward compatibility
type PaymentsServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error)
	GetPlanDetail(context.Context, *GetPlanDetailRequest) (*GetPlanDetailResponse, error)
	GetUserUsagePermissions(context.Context, *GetUserUsagePermissionsRequest) (*GetUserUsagePermissionsResponse, error)
	mustEmbedUnimplementedPaymentsServiceServer()
}

// UnimplementedPaymentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsServiceServer struct {
}

func (UnimplementedPaymentsServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedPaymentsServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedPaymentsServiceServer) GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedPaymentsServiceServer) GetPlanDetail(context.Context, *GetPlanDetailRequest) (*GetPlanDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanDetail not implemented")
}
func (UnimplementedPaymentsServiceServer) GetUserUsagePermissions(context.Context, *GetUserUsagePermissionsRequest) (*GetUserUsagePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserUsagePermissions not implemented")
}
func (UnimplementedPaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {}

// UnsafePaymentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsServiceServer will
// result in compilation errors.
type UnsafePaymentsServiceServer interface {
	mustEmbedUnimplementedPaymentsServiceServer()
}

func RegisterPaymentsServiceServer(s grpc.ServiceRegistrar, srv PaymentsServiceServer) {
	s.RegisterService(&PaymentsService_ServiceDesc, srv)
}

func _PaymentsService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentsService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentsService_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_GetOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).GetOrderStatus(ctx, req.(*GetOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentsService_GetPlanDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).GetPlanDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_GetPlanDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).GetPlanDetail(ctx, req.(*GetPlanDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentsService_GetUserUsagePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserUsagePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).GetUserUsagePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_GetUserUsagePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).GetUserUsagePermissions(ctx, req.(*GetUserUsagePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentsService_ServiceDesc is the grpc.ServiceDesc for PaymentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.payment.v1.PaymentsService",
	HandlerType: (*PaymentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _PaymentsService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _PaymentsService_ListOrders_Handler,
		},
		{
			MethodName: "GetOrderStatus",
			Handler:    _PaymentsService_GetOrderStatus_Handler,
		},
		{
			MethodName: "GetPlanDetail",
			Handler:    _PaymentsService_GetPlanDetail_Handler,
		},
		{
			MethodName: "GetUserUsagePermissions",
			Handler:    _PaymentsService_GetUserUsagePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
