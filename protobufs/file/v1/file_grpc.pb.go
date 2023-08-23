// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: file.proto

package filepb

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*Empty, error)
	GetFileList(ctx context.Context, in *GetFileListRequest, opts ...grpc.CallOption) (*GetFileListResponse, error)
	GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...grpc.CallOption) (*GetFileInfoResponse, error)
	GetPublicS3DownloadURL(ctx context.Context, in *GetPublicS3DownloadURLRequest, opts ...grpc.CallOption) (*GetS3DownloadURLResponse, error)
	GetPrivateS3DownloadURL(ctx context.Context, in *GetPrivateS3DownloadURLRequest, opts ...grpc.CallOption) (*GetS3DownloadURLResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileList(ctx context.Context, in *GetFileListRequest, opts ...grpc.CallOption) (*GetFileListResponse, error) {
	out := new(GetFileListResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/GetFileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...grpc.CallOption) (*GetFileInfoResponse, error) {
	out := new(GetFileInfoResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/GetFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetPublicS3DownloadURL(ctx context.Context, in *GetPublicS3DownloadURLRequest, opts ...grpc.CallOption) (*GetS3DownloadURLResponse, error) {
	out := new(GetS3DownloadURLResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/GetPublicS3DownloadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetPrivateS3DownloadURL(ctx context.Context, in *GetPrivateS3DownloadURLRequest, opts ...grpc.CallOption) (*GetS3DownloadURLResponse, error) {
	out := new(GetS3DownloadURLResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.file.v1.FileService/GetPrivateS3DownloadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadFile(context.Context, *UploadRequest) (*UploadResponse, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*Empty, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*Empty, error)
	GetFileList(context.Context, *GetFileListRequest) (*GetFileListResponse, error)
	GetFileInfo(context.Context, *GetFileInfoRequest) (*GetFileInfoResponse, error)
	GetPublicS3DownloadURL(context.Context, *GetPublicS3DownloadURLRequest) (*GetS3DownloadURLResponse, error)
	GetPrivateS3DownloadURL(context.Context, *GetPrivateS3DownloadURLRequest) (*GetS3DownloadURLResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadFile(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileServiceServer) UpdateFile(context.Context, *UpdateFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileServiceServer) GetFileList(context.Context, *GetFileListRequest) (*GetFileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileList not implemented")
}
func (UnimplementedFileServiceServer) GetFileInfo(context.Context, *GetFileInfoRequest) (*GetFileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (UnimplementedFileServiceServer) GetPublicS3DownloadURL(context.Context, *GetPublicS3DownloadURLRequest) (*GetS3DownloadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicS3DownloadURL not implemented")
}
func (UnimplementedFileServiceServer) GetPrivateS3DownloadURL(context.Context, *GetPrivateS3DownloadURLRequest) (*GetS3DownloadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateS3DownloadURL not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadFile(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/GetFileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileList(ctx, req.(*GetFileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileInfo(ctx, req.(*GetFileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetPublicS3DownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicS3DownloadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetPublicS3DownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/GetPublicS3DownloadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetPublicS3DownloadURL(ctx, req.(*GetPublicS3DownloadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetPrivateS3DownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateS3DownloadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetPrivateS3DownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.file.v1.FileService/GetPrivateS3DownloadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetPrivateS3DownloadURL(ctx, req.(*GetPrivateS3DownloadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.file.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _FileService_UploadFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileService_UpdateFile_Handler,
		},
		{
			MethodName: "GetFileList",
			Handler:    _FileService_GetFileList_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _FileService_GetFileInfo_Handler,
		},
		{
			MethodName: "GetPublicS3DownloadURL",
			Handler:    _FileService_GetPublicS3DownloadURL_Handler,
		},
		{
			MethodName: "GetPrivateS3DownloadURL",
			Handler:    _FileService_GetPrivateS3DownloadURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
