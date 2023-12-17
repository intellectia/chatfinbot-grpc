// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: subscription/v1/subscription.proto

package subscriptionpb

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

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	GetOrderInfo(ctx context.Context, in *GetOrderInfoRequest, opts ...grpc.CallOption) (*GetOrderInfoResponse, error)
	GetPlanDetail(ctx context.Context, in *GetPlanDetailRequest, opts ...grpc.CallOption) (*GetPlanDetailResponse, error)
	GetUserUsage(ctx context.Context, in *GetUserUsageRequest, opts ...grpc.CallOption) (*GetUserUsageResponse, error)
	ConsumeUserPackageUsage(ctx context.Context, in *ConsumeUserPackageUsageRequest, opts ...grpc.CallOption) (*ConsumeUserPackageUsageResponse, error)
	CreateUsages(ctx context.Context, in *CreateUsagesRequest, opts ...grpc.CallOption) (*CreateUsagesResponse, error)
	GetPlans(ctx context.Context, in *GetPlansRequest, opts ...grpc.CallOption) (*GetPlansResponse, error)
	CancelSubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...grpc.CallOption) (*CancelSubscriptionResponse, error)
	TriggerPaypalSubscription(ctx context.Context, in *TriggerPaypalSubscriptionRequest, opts ...grpc.CallOption) (*TriggerPaypalSubscriptionResponse, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetOrderInfo(ctx context.Context, in *GetOrderInfoRequest, opts ...grpc.CallOption) (*GetOrderInfoResponse, error) {
	out := new(GetOrderInfoResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/GetOrderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetPlanDetail(ctx context.Context, in *GetPlanDetailRequest, opts ...grpc.CallOption) (*GetPlanDetailResponse, error) {
	out := new(GetPlanDetailResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/GetPlanDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetUserUsage(ctx context.Context, in *GetUserUsageRequest, opts ...grpc.CallOption) (*GetUserUsageResponse, error) {
	out := new(GetUserUsageResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/GetUserUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ConsumeUserPackageUsage(ctx context.Context, in *ConsumeUserPackageUsageRequest, opts ...grpc.CallOption) (*ConsumeUserPackageUsageResponse, error) {
	out := new(ConsumeUserPackageUsageResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/ConsumeUserPackageUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CreateUsages(ctx context.Context, in *CreateUsagesRequest, opts ...grpc.CallOption) (*CreateUsagesResponse, error) {
	out := new(CreateUsagesResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/CreateUsages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetPlans(ctx context.Context, in *GetPlansRequest, opts ...grpc.CallOption) (*GetPlansResponse, error) {
	out := new(GetPlansResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/GetPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CancelSubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...grpc.CallOption) (*CancelSubscriptionResponse, error) {
	out := new(CancelSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/CancelSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) TriggerPaypalSubscription(ctx context.Context, in *TriggerPaypalSubscriptionRequest, opts ...grpc.CallOption) (*TriggerPaypalSubscriptionResponse, error) {
	out := new(TriggerPaypalSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.subscription.v1.SubscriptionService/TriggerPaypalSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	GetOrderInfo(context.Context, *GetOrderInfoRequest) (*GetOrderInfoResponse, error)
	GetPlanDetail(context.Context, *GetPlanDetailRequest) (*GetPlanDetailResponse, error)
	GetUserUsage(context.Context, *GetUserUsageRequest) (*GetUserUsageResponse, error)
	ConsumeUserPackageUsage(context.Context, *ConsumeUserPackageUsageRequest) (*ConsumeUserPackageUsageResponse, error)
	CreateUsages(context.Context, *CreateUsagesRequest) (*CreateUsagesResponse, error)
	GetPlans(context.Context, *GetPlansRequest) (*GetPlansResponse, error)
	CancelSubscription(context.Context, *CancelSubscriptionRequest) (*CancelSubscriptionResponse, error)
	TriggerPaypalSubscription(context.Context, *TriggerPaypalSubscriptionRequest) (*TriggerPaypalSubscriptionResponse, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetOrderInfo(context.Context, *GetOrderInfoRequest) (*GetOrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInfo not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetPlanDetail(context.Context, *GetPlanDetailRequest) (*GetPlanDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanDetail not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetUserUsage(context.Context, *GetUserUsageRequest) (*GetUserUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserUsage not implemented")
}
func (UnimplementedSubscriptionServiceServer) ConsumeUserPackageUsage(context.Context, *ConsumeUserPackageUsageRequest) (*ConsumeUserPackageUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeUserPackageUsage not implemented")
}
func (UnimplementedSubscriptionServiceServer) CreateUsages(context.Context, *CreateUsagesRequest) (*CreateUsagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUsages not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetPlans(context.Context, *GetPlansRequest) (*GetPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlans not implemented")
}
func (UnimplementedSubscriptionServiceServer) CancelSubscription(context.Context, *CancelSubscriptionRequest) (*CancelSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) TriggerPaypalSubscription(context.Context, *TriggerPaypalSubscriptionRequest) (*TriggerPaypalSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerPaypalSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/GetOrderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetOrderInfo(ctx, req.(*GetOrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetPlanDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetPlanDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/GetPlanDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetPlanDetail(ctx, req.(*GetPlanDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetUserUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetUserUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/GetUserUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetUserUsage(ctx, req.(*GetUserUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ConsumeUserPackageUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeUserPackageUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ConsumeUserPackageUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/ConsumeUserPackageUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ConsumeUserPackageUsage(ctx, req.(*ConsumeUserPackageUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CreateUsages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUsagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateUsages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/CreateUsages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateUsages(ctx, req.(*CreateUsagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/GetPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetPlans(ctx, req.(*GetPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CancelSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CancelSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/CancelSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CancelSubscription(ctx, req.(*CancelSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_TriggerPaypalSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerPaypalSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).TriggerPaypalSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.subscription.v1.SubscriptionService/TriggerPaypalSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).TriggerPaypalSubscription(ctx, req.(*TriggerPaypalSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.subscription.v1.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _SubscriptionService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _SubscriptionService_ListOrders_Handler,
		},
		{
			MethodName: "GetOrderInfo",
			Handler:    _SubscriptionService_GetOrderInfo_Handler,
		},
		{
			MethodName: "GetPlanDetail",
			Handler:    _SubscriptionService_GetPlanDetail_Handler,
		},
		{
			MethodName: "GetUserUsage",
			Handler:    _SubscriptionService_GetUserUsage_Handler,
		},
		{
			MethodName: "ConsumeUserPackageUsage",
			Handler:    _SubscriptionService_ConsumeUserPackageUsage_Handler,
		},
		{
			MethodName: "CreateUsages",
			Handler:    _SubscriptionService_CreateUsages_Handler,
		},
		{
			MethodName: "GetPlans",
			Handler:    _SubscriptionService_GetPlans_Handler,
		},
		{
			MethodName: "CancelSubscription",
			Handler:    _SubscriptionService_CancelSubscription_Handler,
		},
		{
			MethodName: "TriggerPaypalSubscription",
			Handler:    _SubscriptionService_TriggerPaypalSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription/v1/subscription.proto",
}
