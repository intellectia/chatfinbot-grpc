// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: news.proto

package newspb

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
	NewsService_GetNewsList_FullMethodName      = "/chatfinbot.news.v1.NewsService/GetNewsList"
	NewsService_GetNewsInfo_FullMethodName      = "/chatfinbot.news.v1.NewsService/GetNewsInfo"
	NewsService_GetCacheInfo_FullMethodName     = "/chatfinbot.news.v1.NewsService/GetCacheInfo"
	NewsService_GetSentimentList_FullMethodName = "/chatfinbot.news.v1.NewsService/GetSentimentList"
)

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsServiceClient interface {
	// Get a list of news articles
	GetNewsList(ctx context.Context, in *GetNewsListReq, opts ...grpc.CallOption) (*GetNewsListRsp, error)
	// Get detailed information about a specific news article
	GetNewsInfo(ctx context.Context, in *GetNewsInfoReq, opts ...grpc.CallOption) (*GetNewsInfoRsp, error)
	// Get detailed information about a specific news article
	GetCacheInfo(ctx context.Context, in *GetCacheInfoReq, opts ...grpc.CallOption) (*GetCacheInfoRsp, error)
	// Get a list of sentiment
	GetSentimentList(ctx context.Context, in *GetSentimentListReq, opts ...grpc.CallOption) (*GetSentimentListRsp, error)
}

type newsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsServiceClient(cc grpc.ClientConnInterface) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) GetNewsList(ctx context.Context, in *GetNewsListReq, opts ...grpc.CallOption) (*GetNewsListRsp, error) {
	out := new(GetNewsListRsp)
	err := c.cc.Invoke(ctx, NewsService_GetNewsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetNewsInfo(ctx context.Context, in *GetNewsInfoReq, opts ...grpc.CallOption) (*GetNewsInfoRsp, error) {
	out := new(GetNewsInfoRsp)
	err := c.cc.Invoke(ctx, NewsService_GetNewsInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetCacheInfo(ctx context.Context, in *GetCacheInfoReq, opts ...grpc.CallOption) (*GetCacheInfoRsp, error) {
	out := new(GetCacheInfoRsp)
	err := c.cc.Invoke(ctx, NewsService_GetCacheInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetSentimentList(ctx context.Context, in *GetSentimentListReq, opts ...grpc.CallOption) (*GetSentimentListRsp, error) {
	out := new(GetSentimentListRsp)
	err := c.cc.Invoke(ctx, NewsService_GetSentimentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
// All implementations must embed UnimplementedNewsServiceServer
// for forward compatibility
type NewsServiceServer interface {
	// Get a list of news articles
	GetNewsList(context.Context, *GetNewsListReq) (*GetNewsListRsp, error)
	// Get detailed information about a specific news article
	GetNewsInfo(context.Context, *GetNewsInfoReq) (*GetNewsInfoRsp, error)
	// Get detailed information about a specific news article
	GetCacheInfo(context.Context, *GetCacheInfoReq) (*GetCacheInfoRsp, error)
	// Get a list of sentiment
	GetSentimentList(context.Context, *GetSentimentListReq) (*GetSentimentListRsp, error)
	mustEmbedUnimplementedNewsServiceServer()
}

// UnimplementedNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (UnimplementedNewsServiceServer) GetNewsList(context.Context, *GetNewsListReq) (*GetNewsListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsList not implemented")
}
func (UnimplementedNewsServiceServer) GetNewsInfo(context.Context, *GetNewsInfoReq) (*GetNewsInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsInfo not implemented")
}
func (UnimplementedNewsServiceServer) GetCacheInfo(context.Context, *GetCacheInfoReq) (*GetCacheInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheInfo not implemented")
}
func (UnimplementedNewsServiceServer) GetSentimentList(context.Context, *GetSentimentListReq) (*GetSentimentListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSentimentList not implemented")
}
func (UnimplementedNewsServiceServer) mustEmbedUnimplementedNewsServiceServer() {}

// UnsafeNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServiceServer will
// result in compilation errors.
type UnsafeNewsServiceServer interface {
	mustEmbedUnimplementedNewsServiceServer()
}

func RegisterNewsServiceServer(s grpc.ServiceRegistrar, srv NewsServiceServer) {
	s.RegisterService(&NewsService_ServiceDesc, srv)
}

func _NewsService_GetNewsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNewsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsService_GetNewsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNewsList(ctx, req.(*GetNewsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetNewsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNewsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsService_GetNewsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNewsInfo(ctx, req.(*GetNewsInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetCacheInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetCacheInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsService_GetCacheInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetCacheInfo(ctx, req.(*GetCacheInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetSentimentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSentimentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetSentimentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsService_GetSentimentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetSentimentList(ctx, req.(*GetSentimentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsService_ServiceDesc is the grpc.ServiceDesc for NewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.news.v1.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewsList",
			Handler:    _NewsService_GetNewsList_Handler,
		},
		{
			MethodName: "GetNewsInfo",
			Handler:    _NewsService_GetNewsInfo_Handler,
		},
		{
			MethodName: "GetCacheInfo",
			Handler:    _NewsService_GetCacheInfo_Handler,
		},
		{
			MethodName: "GetSentimentList",
			Handler:    _NewsService_GetSentimentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}
