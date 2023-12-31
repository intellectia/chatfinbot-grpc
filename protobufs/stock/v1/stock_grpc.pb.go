// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: stock.proto

package stockpb

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
	StockService_Query_FullMethodName             = "/chatfinbot.stock.v1.StockService/Query"
	StockService_Indicators_FullMethodName        = "/chatfinbot.stock.v1.StockService/Indicators"
	StockService_Watchlists_FullMethodName        = "/chatfinbot.stock.v1.StockService/Watchlists"
	StockService_CreateWatchlist_FullMethodName   = "/chatfinbot.stock.v1.StockService/CreateWatchlist"
	StockService_DelWatchlist_FullMethodName      = "/chatfinbot.stock.v1.StockService/DelWatchlist"
	StockService_AddToWatchlist_FullMethodName    = "/chatfinbot.stock.v1.StockService/AddToWatchlist"
	StockService_DelFromWatchlist_FullMethodName  = "/chatfinbot.stock.v1.StockService/DelFromWatchlist"
	StockService_MoveFromWatchlist_FullMethodName = "/chatfinbot.stock.v1.StockService/MoveFromWatchlist"
	StockService_SortFromWatchlist_FullMethodName = "/chatfinbot.stock.v1.StockService/SortFromWatchlist"
	StockService_Realtime_FullMethodName          = "/chatfinbot.stock.v1.StockService/Realtime"
	StockService_GetEventTimeline_FullMethodName  = "/chatfinbot.stock.v1.StockService/GetEventTimeline"
	StockService_GetEventList_FullMethodName      = "/chatfinbot.stock.v1.StockService/GetEventList"
	StockService_StockHistory_FullMethodName      = "/chatfinbot.stock.v1.StockService/StockHistory"
	StockService_StockAccess_FullMethodName       = "/chatfinbot.stock.v1.StockService/StockAccess"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	// 查询数据
	Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error)
	// 技术指标
	Indicators(ctx context.Context, in *IndicatorsReq, opts ...grpc.CallOption) (*IndicatorsRsp, error)
	// 监视列表
	Watchlists(ctx context.Context, in *WatchlistsReq, opts ...grpc.CallOption) (*WatchlistsRsp, error)
	// 创建
	CreateWatchlist(ctx context.Context, in *CreateWatchlistReq, opts ...grpc.CallOption) (*CreateWatchlistRsp, error)
	// 删除
	DelWatchlist(ctx context.Context, in *DelWatchlistReq, opts ...grpc.CallOption) (*DelWatchlistRsp, error)
	// 添加
	AddToWatchlist(ctx context.Context, in *AddToWatchlistReq, opts ...grpc.CallOption) (*AddToWatchlistRsp, error)
	// 删除
	DelFromWatchlist(ctx context.Context, in *DelFromWatchlistReq, opts ...grpc.CallOption) (*DelFromWatchlistRsp, error)
	// 移动
	MoveFromWatchlist(ctx context.Context, in *MoveFromWatchlistReq, opts ...grpc.CallOption) (*MoveFromWatchlistRsp, error)
	// 排序
	SortFromWatchlist(ctx context.Context, in *SortFromWatchlistReq, opts ...grpc.CallOption) (*SortFromWatchlistRsp, error)
	// 实时数据
	Realtime(ctx context.Context, in *RealtimeReq, opts ...grpc.CallOption) (*RealtimeRsp, error)
	// 股票事件
	GetEventTimeline(ctx context.Context, in *GetEventTimelineReq, opts ...grpc.CallOption) (*GetEventTimelineRsp, error)
	// 股票事件
	GetEventList(ctx context.Context, in *GetEventListReq, opts ...grpc.CallOption) (*GetEventListRsp, error)
	// 股票历史
	StockHistory(ctx context.Context, in *StockHistoryReq, opts ...grpc.CallOption) (*StockHistoryRsp, error)
	// 股票历史
	StockAccess(ctx context.Context, in *StockAccessReq, opts ...grpc.CallOption) (*StockAccessRsp, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRsp, error) {
	out := new(QueryRsp)
	err := c.cc.Invoke(ctx, StockService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) Indicators(ctx context.Context, in *IndicatorsReq, opts ...grpc.CallOption) (*IndicatorsRsp, error) {
	out := new(IndicatorsRsp)
	err := c.cc.Invoke(ctx, StockService_Indicators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) Watchlists(ctx context.Context, in *WatchlistsReq, opts ...grpc.CallOption) (*WatchlistsRsp, error) {
	out := new(WatchlistsRsp)
	err := c.cc.Invoke(ctx, StockService_Watchlists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CreateWatchlist(ctx context.Context, in *CreateWatchlistReq, opts ...grpc.CallOption) (*CreateWatchlistRsp, error) {
	out := new(CreateWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_CreateWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) DelWatchlist(ctx context.Context, in *DelWatchlistReq, opts ...grpc.CallOption) (*DelWatchlistRsp, error) {
	out := new(DelWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_DelWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) AddToWatchlist(ctx context.Context, in *AddToWatchlistReq, opts ...grpc.CallOption) (*AddToWatchlistRsp, error) {
	out := new(AddToWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_AddToWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) DelFromWatchlist(ctx context.Context, in *DelFromWatchlistReq, opts ...grpc.CallOption) (*DelFromWatchlistRsp, error) {
	out := new(DelFromWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_DelFromWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) MoveFromWatchlist(ctx context.Context, in *MoveFromWatchlistReq, opts ...grpc.CallOption) (*MoveFromWatchlistRsp, error) {
	out := new(MoveFromWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_MoveFromWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) SortFromWatchlist(ctx context.Context, in *SortFromWatchlistReq, opts ...grpc.CallOption) (*SortFromWatchlistRsp, error) {
	out := new(SortFromWatchlistRsp)
	err := c.cc.Invoke(ctx, StockService_SortFromWatchlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) Realtime(ctx context.Context, in *RealtimeReq, opts ...grpc.CallOption) (*RealtimeRsp, error) {
	out := new(RealtimeRsp)
	err := c.cc.Invoke(ctx, StockService_Realtime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) GetEventTimeline(ctx context.Context, in *GetEventTimelineReq, opts ...grpc.CallOption) (*GetEventTimelineRsp, error) {
	out := new(GetEventTimelineRsp)
	err := c.cc.Invoke(ctx, StockService_GetEventTimeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) GetEventList(ctx context.Context, in *GetEventListReq, opts ...grpc.CallOption) (*GetEventListRsp, error) {
	out := new(GetEventListRsp)
	err := c.cc.Invoke(ctx, StockService_GetEventList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) StockHistory(ctx context.Context, in *StockHistoryReq, opts ...grpc.CallOption) (*StockHistoryRsp, error) {
	out := new(StockHistoryRsp)
	err := c.cc.Invoke(ctx, StockService_StockHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) StockAccess(ctx context.Context, in *StockAccessReq, opts ...grpc.CallOption) (*StockAccessRsp, error) {
	out := new(StockAccessRsp)
	err := c.cc.Invoke(ctx, StockService_StockAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	// 查询数据
	Query(context.Context, *QueryReq) (*QueryRsp, error)
	// 技术指标
	Indicators(context.Context, *IndicatorsReq) (*IndicatorsRsp, error)
	// 监视列表
	Watchlists(context.Context, *WatchlistsReq) (*WatchlistsRsp, error)
	// 创建
	CreateWatchlist(context.Context, *CreateWatchlistReq) (*CreateWatchlistRsp, error)
	// 删除
	DelWatchlist(context.Context, *DelWatchlistReq) (*DelWatchlistRsp, error)
	// 添加
	AddToWatchlist(context.Context, *AddToWatchlistReq) (*AddToWatchlistRsp, error)
	// 删除
	DelFromWatchlist(context.Context, *DelFromWatchlistReq) (*DelFromWatchlistRsp, error)
	// 移动
	MoveFromWatchlist(context.Context, *MoveFromWatchlistReq) (*MoveFromWatchlistRsp, error)
	// 排序
	SortFromWatchlist(context.Context, *SortFromWatchlistReq) (*SortFromWatchlistRsp, error)
	// 实时数据
	Realtime(context.Context, *RealtimeReq) (*RealtimeRsp, error)
	// 股票事件
	GetEventTimeline(context.Context, *GetEventTimelineReq) (*GetEventTimelineRsp, error)
	// 股票事件
	GetEventList(context.Context, *GetEventListReq) (*GetEventListRsp, error)
	// 股票历史
	StockHistory(context.Context, *StockHistoryReq) (*StockHistoryRsp, error)
	// 股票历史
	StockAccess(context.Context, *StockAccessReq) (*StockAccessRsp, error)
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) Query(context.Context, *QueryReq) (*QueryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedStockServiceServer) Indicators(context.Context, *IndicatorsReq) (*IndicatorsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Indicators not implemented")
}
func (UnimplementedStockServiceServer) Watchlists(context.Context, *WatchlistsReq) (*WatchlistsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watchlists not implemented")
}
func (UnimplementedStockServiceServer) CreateWatchlist(context.Context, *CreateWatchlistReq) (*CreateWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWatchlist not implemented")
}
func (UnimplementedStockServiceServer) DelWatchlist(context.Context, *DelWatchlistReq) (*DelWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelWatchlist not implemented")
}
func (UnimplementedStockServiceServer) AddToWatchlist(context.Context, *AddToWatchlistReq) (*AddToWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWatchlist not implemented")
}
func (UnimplementedStockServiceServer) DelFromWatchlist(context.Context, *DelFromWatchlistReq) (*DelFromWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFromWatchlist not implemented")
}
func (UnimplementedStockServiceServer) MoveFromWatchlist(context.Context, *MoveFromWatchlistReq) (*MoveFromWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveFromWatchlist not implemented")
}
func (UnimplementedStockServiceServer) SortFromWatchlist(context.Context, *SortFromWatchlistReq) (*SortFromWatchlistRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortFromWatchlist not implemented")
}
func (UnimplementedStockServiceServer) Realtime(context.Context, *RealtimeReq) (*RealtimeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Realtime not implemented")
}
func (UnimplementedStockServiceServer) GetEventTimeline(context.Context, *GetEventTimelineReq) (*GetEventTimelineRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTimeline not implemented")
}
func (UnimplementedStockServiceServer) GetEventList(context.Context, *GetEventListReq) (*GetEventListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventList not implemented")
}
func (UnimplementedStockServiceServer) StockHistory(context.Context, *StockHistoryReq) (*StockHistoryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockHistory not implemented")
}
func (UnimplementedStockServiceServer) StockAccess(context.Context, *StockAccessReq) (*StockAccessRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockAccess not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Query(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_Indicators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndicatorsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Indicators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Indicators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Indicators(ctx, req.(*IndicatorsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_Watchlists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchlistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Watchlists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Watchlists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Watchlists(ctx, req.(*WatchlistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CreateWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CreateWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_CreateWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CreateWatchlist(ctx, req.(*CreateWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_DelWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).DelWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_DelWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).DelWatchlist(ctx, req.(*DelWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_AddToWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).AddToWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_AddToWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).AddToWatchlist(ctx, req.(*AddToWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_DelFromWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFromWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).DelFromWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_DelFromWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).DelFromWatchlist(ctx, req.(*DelFromWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_MoveFromWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveFromWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).MoveFromWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_MoveFromWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).MoveFromWatchlist(ctx, req.(*MoveFromWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_SortFromWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortFromWatchlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).SortFromWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_SortFromWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).SortFromWatchlist(ctx, req.(*SortFromWatchlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_Realtime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RealtimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).Realtime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_Realtime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).Realtime(ctx, req.(*RealtimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_GetEventTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventTimelineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetEventTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_GetEventTimeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetEventTimeline(ctx, req.(*GetEventTimelineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_GetEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_GetEventList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetEventList(ctx, req.(*GetEventListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_StockHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).StockHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_StockHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).StockHistory(ctx, req.(*StockHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_StockAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).StockAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_StockAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).StockAccess(ctx, req.(*StockAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.stock.v1.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _StockService_Query_Handler,
		},
		{
			MethodName: "Indicators",
			Handler:    _StockService_Indicators_Handler,
		},
		{
			MethodName: "Watchlists",
			Handler:    _StockService_Watchlists_Handler,
		},
		{
			MethodName: "CreateWatchlist",
			Handler:    _StockService_CreateWatchlist_Handler,
		},
		{
			MethodName: "DelWatchlist",
			Handler:    _StockService_DelWatchlist_Handler,
		},
		{
			MethodName: "AddToWatchlist",
			Handler:    _StockService_AddToWatchlist_Handler,
		},
		{
			MethodName: "DelFromWatchlist",
			Handler:    _StockService_DelFromWatchlist_Handler,
		},
		{
			MethodName: "MoveFromWatchlist",
			Handler:    _StockService_MoveFromWatchlist_Handler,
		},
		{
			MethodName: "SortFromWatchlist",
			Handler:    _StockService_SortFromWatchlist_Handler,
		},
		{
			MethodName: "Realtime",
			Handler:    _StockService_Realtime_Handler,
		},
		{
			MethodName: "GetEventTimeline",
			Handler:    _StockService_GetEventTimeline_Handler,
		},
		{
			MethodName: "GetEventList",
			Handler:    _StockService_GetEventList_Handler,
		},
		{
			MethodName: "StockHistory",
			Handler:    _StockService_StockHistory_Handler,
		},
		{
			MethodName: "StockAccess",
			Handler:    _StockService_StockAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}
