// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: search.proto

package searchpb

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
	SearchService_GetCompanyInfo_FullMethodName  = "/chatfinbot.search.v1.SearchService/GetCompanyInfo"
	SearchService_GetIndicators_FullMethodName   = "/chatfinbot.search.v1.SearchService/GetIndicators"
	SearchService_GetShareholders_FullMethodName = "/chatfinbot.search.v1.SearchService/GetShareholders"
	SearchService_GetExecutives_FullMethodName   = "/chatfinbot.search.v1.SearchService/GetExecutives"
	SearchService_GetStaffInfo_FullMethodName    = "/chatfinbot.search.v1.SearchService/GetStaffInfo"
	SearchService_GetMBRevenue_FullMethodName    = "/chatfinbot.search.v1.SearchService/GetMBRevenue"
	SearchService_GetFinancial_FullMethodName    = "/chatfinbot.search.v1.SearchService/GetFinancial"
)

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	// 获取公司数据
	GetCompanyInfo(ctx context.Context, in *GetCompanyInfoReq, opts ...grpc.CallOption) (*GetCompanyInfoRsp, error)
	// 获取公司行情指标
	GetIndicators(ctx context.Context, in *GetIndicatorsReq, opts ...grpc.CallOption) (*GetIndicatorsRsp, error)
	// 获取公司股权结构
	GetShareholders(ctx context.Context, in *GetShareholdersReq, opts ...grpc.CallOption) (*GetShareholdersRsp, error)
	// 获取公司管理层信息
	GetExecutives(ctx context.Context, in *GetExecutivesReq, opts ...grpc.CallOption) (*GetExecutivesRsp, error)
	// 获取员工构成
	GetStaffInfo(ctx context.Context, in *GetStaffInfoReq, opts ...grpc.CallOption) (*GetStaffInfoRsp, error)
	// 获取主营营收
	GetMBRevenue(ctx context.Context, in *GetMBRevenueReq, opts ...grpc.CallOption) (*GetMBRevenueRsp, error)
	// 获取财务概况
	GetFinancial(ctx context.Context, in *GetFinancialReq, opts ...grpc.CallOption) (*GetFinancialRsp, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) GetCompanyInfo(ctx context.Context, in *GetCompanyInfoReq, opts ...grpc.CallOption) (*GetCompanyInfoRsp, error) {
	out := new(GetCompanyInfoRsp)
	err := c.cc.Invoke(ctx, SearchService_GetCompanyInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetIndicators(ctx context.Context, in *GetIndicatorsReq, opts ...grpc.CallOption) (*GetIndicatorsRsp, error) {
	out := new(GetIndicatorsRsp)
	err := c.cc.Invoke(ctx, SearchService_GetIndicators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetShareholders(ctx context.Context, in *GetShareholdersReq, opts ...grpc.CallOption) (*GetShareholdersRsp, error) {
	out := new(GetShareholdersRsp)
	err := c.cc.Invoke(ctx, SearchService_GetShareholders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetExecutives(ctx context.Context, in *GetExecutivesReq, opts ...grpc.CallOption) (*GetExecutivesRsp, error) {
	out := new(GetExecutivesRsp)
	err := c.cc.Invoke(ctx, SearchService_GetExecutives_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetStaffInfo(ctx context.Context, in *GetStaffInfoReq, opts ...grpc.CallOption) (*GetStaffInfoRsp, error) {
	out := new(GetStaffInfoRsp)
	err := c.cc.Invoke(ctx, SearchService_GetStaffInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetMBRevenue(ctx context.Context, in *GetMBRevenueReq, opts ...grpc.CallOption) (*GetMBRevenueRsp, error) {
	out := new(GetMBRevenueRsp)
	err := c.cc.Invoke(ctx, SearchService_GetMBRevenue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetFinancial(ctx context.Context, in *GetFinancialReq, opts ...grpc.CallOption) (*GetFinancialRsp, error) {
	out := new(GetFinancialRsp)
	err := c.cc.Invoke(ctx, SearchService_GetFinancial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	// 获取公司数据
	GetCompanyInfo(context.Context, *GetCompanyInfoReq) (*GetCompanyInfoRsp, error)
	// 获取公司行情指标
	GetIndicators(context.Context, *GetIndicatorsReq) (*GetIndicatorsRsp, error)
	// 获取公司股权结构
	GetShareholders(context.Context, *GetShareholdersReq) (*GetShareholdersRsp, error)
	// 获取公司管理层信息
	GetExecutives(context.Context, *GetExecutivesReq) (*GetExecutivesRsp, error)
	// 获取员工构成
	GetStaffInfo(context.Context, *GetStaffInfoReq) (*GetStaffInfoRsp, error)
	// 获取主营营收
	GetMBRevenue(context.Context, *GetMBRevenueReq) (*GetMBRevenueRsp, error)
	// 获取财务概况
	GetFinancial(context.Context, *GetFinancialReq) (*GetFinancialRsp, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) GetCompanyInfo(context.Context, *GetCompanyInfoReq) (*GetCompanyInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyInfo not implemented")
}
func (UnimplementedSearchServiceServer) GetIndicators(context.Context, *GetIndicatorsReq) (*GetIndicatorsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndicators not implemented")
}
func (UnimplementedSearchServiceServer) GetShareholders(context.Context, *GetShareholdersReq) (*GetShareholdersRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShareholders not implemented")
}
func (UnimplementedSearchServiceServer) GetExecutives(context.Context, *GetExecutivesReq) (*GetExecutivesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecutives not implemented")
}
func (UnimplementedSearchServiceServer) GetStaffInfo(context.Context, *GetStaffInfoReq) (*GetStaffInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffInfo not implemented")
}
func (UnimplementedSearchServiceServer) GetMBRevenue(context.Context, *GetMBRevenueReq) (*GetMBRevenueRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMBRevenue not implemented")
}
func (UnimplementedSearchServiceServer) GetFinancial(context.Context, *GetFinancialReq) (*GetFinancialRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancial not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_GetCompanyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetCompanyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetCompanyInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetCompanyInfo(ctx, req.(*GetCompanyInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetIndicators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndicatorsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetIndicators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetIndicators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetIndicators(ctx, req.(*GetIndicatorsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetShareholders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareholdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetShareholders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetShareholders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetShareholders(ctx, req.(*GetShareholdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetExecutives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExecutivesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetExecutives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetExecutives_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetExecutives(ctx, req.(*GetExecutivesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetStaffInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetStaffInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetStaffInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetStaffInfo(ctx, req.(*GetStaffInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetMBRevenue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMBRevenueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetMBRevenue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetMBRevenue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetMBRevenue(ctx, req.(*GetMBRevenueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetFinancial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetFinancial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_GetFinancial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetFinancial(ctx, req.(*GetFinancialReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.search.v1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyInfo",
			Handler:    _SearchService_GetCompanyInfo_Handler,
		},
		{
			MethodName: "GetIndicators",
			Handler:    _SearchService_GetIndicators_Handler,
		},
		{
			MethodName: "GetShareholders",
			Handler:    _SearchService_GetShareholders_Handler,
		},
		{
			MethodName: "GetExecutives",
			Handler:    _SearchService_GetExecutives_Handler,
		},
		{
			MethodName: "GetStaffInfo",
			Handler:    _SearchService_GetStaffInfo_Handler,
		},
		{
			MethodName: "GetMBRevenue",
			Handler:    _SearchService_GetMBRevenue_Handler,
		},
		{
			MethodName: "GetFinancial",
			Handler:    _SearchService_GetFinancial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}