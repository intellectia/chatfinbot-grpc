// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: community.proto

package communitypb

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
	CommunityService_CreatePost_FullMethodName    = "/chatfinbot.community.v1.CommunityService/CreatePost"
	CommunityService_DeletePost_FullMethodName    = "/chatfinbot.community.v1.CommunityService/DeletePost"
	CommunityService_GetStockPosts_FullMethodName = "/chatfinbot.community.v1.CommunityService/GetStockPosts"
	CommunityService_GetMyPosts_FullMethodName    = "/chatfinbot.community.v1.CommunityService/GetMyPosts"
	CommunityService_AddComment_FullMethodName    = "/chatfinbot.community.v1.CommunityService/AddComment"
	CommunityService_DeleteComment_FullMethodName = "/chatfinbot.community.v1.CommunityService/DeleteComment"
	CommunityService_GetMyComment_FullMethodName  = "/chatfinbot.community.v1.CommunityService/GetMyComment"
)

// CommunityServiceClient is the client API for CommunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunityServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRsp, error)
	DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRsp, error)
	GetStockPosts(ctx context.Context, in *GetStockPostsReq, opts ...grpc.CallOption) (*GetStockPostsRsp, error)
	GetMyPosts(ctx context.Context, in *GetMyPostsReq, opts ...grpc.CallOption) (*GetMyPostsRsp, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRsp, error)
	DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentRsp, error)
	GetMyComment(ctx context.Context, in *GetMyCommentsReq, opts ...grpc.CallOption) (*GetMyCommentsRsp, error)
}

type communityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityServiceClient(cc grpc.ClientConnInterface) CommunityServiceClient {
	return &communityServiceClient{cc}
}

func (c *communityServiceClient) CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRsp, error) {
	out := new(CreatePostRsp)
	err := c.cc.Invoke(ctx, CommunityService_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRsp, error) {
	out := new(DeletePostRsp)
	err := c.cc.Invoke(ctx, CommunityService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetStockPosts(ctx context.Context, in *GetStockPostsReq, opts ...grpc.CallOption) (*GetStockPostsRsp, error) {
	out := new(GetStockPostsRsp)
	err := c.cc.Invoke(ctx, CommunityService_GetStockPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetMyPosts(ctx context.Context, in *GetMyPostsReq, opts ...grpc.CallOption) (*GetMyPostsRsp, error) {
	out := new(GetMyPostsRsp)
	err := c.cc.Invoke(ctx, CommunityService_GetMyPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRsp, error) {
	out := new(AddCommentRsp)
	err := c.cc.Invoke(ctx, CommunityService_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentRsp, error) {
	out := new(DeleteCommentRsp)
	err := c.cc.Invoke(ctx, CommunityService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetMyComment(ctx context.Context, in *GetMyCommentsReq, opts ...grpc.CallOption) (*GetMyCommentsRsp, error) {
	out := new(GetMyCommentsRsp)
	err := c.cc.Invoke(ctx, CommunityService_GetMyComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServiceServer is the server API for CommunityService service.
// All implementations must embed UnimplementedCommunityServiceServer
// for forward compatibility
type CommunityServiceServer interface {
	CreatePost(context.Context, *CreatePostReq) (*CreatePostRsp, error)
	DeletePost(context.Context, *DeletePostReq) (*DeletePostRsp, error)
	GetStockPosts(context.Context, *GetStockPostsReq) (*GetStockPostsRsp, error)
	GetMyPosts(context.Context, *GetMyPostsReq) (*GetMyPostsRsp, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentRsp, error)
	DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentRsp, error)
	GetMyComment(context.Context, *GetMyCommentsReq) (*GetMyCommentsRsp, error)
	mustEmbedUnimplementedCommunityServiceServer()
}

// UnimplementedCommunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunityServiceServer struct {
}

func (UnimplementedCommunityServiceServer) CreatePost(context.Context, *CreatePostReq) (*CreatePostRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedCommunityServiceServer) DeletePost(context.Context, *DeletePostReq) (*DeletePostRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedCommunityServiceServer) GetStockPosts(context.Context, *GetStockPostsReq) (*GetStockPostsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockPosts not implemented")
}
func (UnimplementedCommunityServiceServer) GetMyPosts(context.Context, *GetMyPostsReq) (*GetMyPostsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyPosts not implemented")
}
func (UnimplementedCommunityServiceServer) AddComment(context.Context, *AddCommentReq) (*AddCommentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommunityServiceServer) DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommunityServiceServer) GetMyComment(context.Context, *GetMyCommentsReq) (*GetMyCommentsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyComment not implemented")
}
func (UnimplementedCommunityServiceServer) mustEmbedUnimplementedCommunityServiceServer() {}

// UnsafeCommunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunityServiceServer will
// result in compilation errors.
type UnsafeCommunityServiceServer interface {
	mustEmbedUnimplementedCommunityServiceServer()
}

func RegisterCommunityServiceServer(s grpc.ServiceRegistrar, srv CommunityServiceServer) {
	s.RegisterService(&CommunityService_ServiceDesc, srv)
}

func _CommunityService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreatePost(ctx, req.(*CreatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeletePost(ctx, req.(*DeletePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetStockPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockPostsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetStockPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetStockPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetStockPosts(ctx, req.(*GetStockPostsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetMyPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyPostsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetMyPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetMyPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetMyPosts(ctx, req.(*GetMyPostsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeleteComment(ctx, req.(*DeleteCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetMyComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCommentsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetMyComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetMyComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetMyComment(ctx, req.(*GetMyCommentsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunityService_ServiceDesc is the grpc.ServiceDesc for CommunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.community.v1.CommunityService",
	HandlerType: (*CommunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _CommunityService_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _CommunityService_DeletePost_Handler,
		},
		{
			MethodName: "GetStockPosts",
			Handler:    _CommunityService_GetStockPosts_Handler,
		},
		{
			MethodName: "GetMyPosts",
			Handler:    _CommunityService_GetMyPosts_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _CommunityService_AddComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommunityService_DeleteComment_Handler,
		},
		{
			MethodName: "GetMyComment",
			Handler:    _CommunityService_GetMyComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "community.proto",
}