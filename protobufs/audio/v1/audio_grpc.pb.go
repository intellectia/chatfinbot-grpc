// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: audio.proto

package audiopb

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
	AudioService_CreateRecord_FullMethodName    = "/chatfinbot.audio.v1.AudioService/CreateRecord"
	AudioService_GetRecordList_FullMethodName   = "/chatfinbot.audio.v1.AudioService/GetRecordList"
	AudioService_GetRecordStatus_FullMethodName = "/chatfinbot.audio.v1.AudioService/GetRecordStatus"
	AudioService_Transcription_FullMethodName   = "/chatfinbot.audio.v1.AudioService/Transcription"
	AudioService_GetTransData_FullMethodName    = "/chatfinbot.audio.v1.AudioService/GetTransData"
	AudioService_SummaryDepth_FullMethodName    = "/chatfinbot.audio.v1.AudioService/SummaryDepth"
	AudioService_GetSummaryDepth_FullMethodName = "/chatfinbot.audio.v1.AudioService/GetSummaryDepth"
	AudioService_UpdateSummary_FullMethodName   = "/chatfinbot.audio.v1.AudioService/UpdateSummary"
)

// AudioServiceClient is the client API for AudioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioServiceClient interface {
	CreateRecord(ctx context.Context, in *CreateRecordReq, opts ...grpc.CallOption) (*CreateRecordRsp, error)
	GetRecordList(ctx context.Context, in *GetRecordListReq, opts ...grpc.CallOption) (*GetRecordListRsp, error)
	GetRecordStatus(ctx context.Context, in *GetRecordStatusReq, opts ...grpc.CallOption) (*GetRecordStatusRsp, error)
	Transcription(ctx context.Context, in *TranscriptionReq, opts ...grpc.CallOption) (*TranscriptionRsp, error)
	GetTransData(ctx context.Context, in *GetTransDataReq, opts ...grpc.CallOption) (*GetTransDataRsp, error)
	SummaryDepth(ctx context.Context, in *SummaryDepthReq, opts ...grpc.CallOption) (*SummaryDepthRsp, error)
	GetSummaryDepth(ctx context.Context, in *GetSummaryDepthReq, opts ...grpc.CallOption) (*GetSummaryDepthRsp, error)
	UpdateSummary(ctx context.Context, in *UpdateSummaryReq, opts ...grpc.CallOption) (*UpdateSummaryRsp, error)
}

type audioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioServiceClient(cc grpc.ClientConnInterface) AudioServiceClient {
	return &audioServiceClient{cc}
}

func (c *audioServiceClient) CreateRecord(ctx context.Context, in *CreateRecordReq, opts ...grpc.CallOption) (*CreateRecordRsp, error) {
	out := new(CreateRecordRsp)
	err := c.cc.Invoke(ctx, AudioService_CreateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetRecordList(ctx context.Context, in *GetRecordListReq, opts ...grpc.CallOption) (*GetRecordListRsp, error) {
	out := new(GetRecordListRsp)
	err := c.cc.Invoke(ctx, AudioService_GetRecordList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetRecordStatus(ctx context.Context, in *GetRecordStatusReq, opts ...grpc.CallOption) (*GetRecordStatusRsp, error) {
	out := new(GetRecordStatusRsp)
	err := c.cc.Invoke(ctx, AudioService_GetRecordStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) Transcription(ctx context.Context, in *TranscriptionReq, opts ...grpc.CallOption) (*TranscriptionRsp, error) {
	out := new(TranscriptionRsp)
	err := c.cc.Invoke(ctx, AudioService_Transcription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetTransData(ctx context.Context, in *GetTransDataReq, opts ...grpc.CallOption) (*GetTransDataRsp, error) {
	out := new(GetTransDataRsp)
	err := c.cc.Invoke(ctx, AudioService_GetTransData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) SummaryDepth(ctx context.Context, in *SummaryDepthReq, opts ...grpc.CallOption) (*SummaryDepthRsp, error) {
	out := new(SummaryDepthRsp)
	err := c.cc.Invoke(ctx, AudioService_SummaryDepth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetSummaryDepth(ctx context.Context, in *GetSummaryDepthReq, opts ...grpc.CallOption) (*GetSummaryDepthRsp, error) {
	out := new(GetSummaryDepthRsp)
	err := c.cc.Invoke(ctx, AudioService_GetSummaryDepth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) UpdateSummary(ctx context.Context, in *UpdateSummaryReq, opts ...grpc.CallOption) (*UpdateSummaryRsp, error) {
	out := new(UpdateSummaryRsp)
	err := c.cc.Invoke(ctx, AudioService_UpdateSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioServiceServer is the server API for AudioService service.
// All implementations must embed UnimplementedAudioServiceServer
// for forward compatibility
type AudioServiceServer interface {
	CreateRecord(context.Context, *CreateRecordReq) (*CreateRecordRsp, error)
	GetRecordList(context.Context, *GetRecordListReq) (*GetRecordListRsp, error)
	GetRecordStatus(context.Context, *GetRecordStatusReq) (*GetRecordStatusRsp, error)
	Transcription(context.Context, *TranscriptionReq) (*TranscriptionRsp, error)
	GetTransData(context.Context, *GetTransDataReq) (*GetTransDataRsp, error)
	SummaryDepth(context.Context, *SummaryDepthReq) (*SummaryDepthRsp, error)
	GetSummaryDepth(context.Context, *GetSummaryDepthReq) (*GetSummaryDepthRsp, error)
	UpdateSummary(context.Context, *UpdateSummaryReq) (*UpdateSummaryRsp, error)
	mustEmbedUnimplementedAudioServiceServer()
}

// UnimplementedAudioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAudioServiceServer struct {
}

func (UnimplementedAudioServiceServer) CreateRecord(context.Context, *CreateRecordReq) (*CreateRecordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedAudioServiceServer) GetRecordList(context.Context, *GetRecordListReq) (*GetRecordListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecordList not implemented")
}
func (UnimplementedAudioServiceServer) GetRecordStatus(context.Context, *GetRecordStatusReq) (*GetRecordStatusRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecordStatus not implemented")
}
func (UnimplementedAudioServiceServer) Transcription(context.Context, *TranscriptionReq) (*TranscriptionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transcription not implemented")
}
func (UnimplementedAudioServiceServer) GetTransData(context.Context, *GetTransDataReq) (*GetTransDataRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransData not implemented")
}
func (UnimplementedAudioServiceServer) SummaryDepth(context.Context, *SummaryDepthReq) (*SummaryDepthRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryDepth not implemented")
}
func (UnimplementedAudioServiceServer) GetSummaryDepth(context.Context, *GetSummaryDepthReq) (*GetSummaryDepthRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummaryDepth not implemented")
}
func (UnimplementedAudioServiceServer) UpdateSummary(context.Context, *UpdateSummaryReq) (*UpdateSummaryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSummary not implemented")
}
func (UnimplementedAudioServiceServer) mustEmbedUnimplementedAudioServiceServer() {}

// UnsafeAudioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServiceServer will
// result in compilation errors.
type UnsafeAudioServiceServer interface {
	mustEmbedUnimplementedAudioServiceServer()
}

func RegisterAudioServiceServer(s grpc.ServiceRegistrar, srv AudioServiceServer) {
	s.RegisterService(&AudioService_ServiceDesc, srv)
}

func _AudioService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_CreateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).CreateRecord(ctx, req.(*CreateRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetRecordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetRecordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetRecordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetRecordList(ctx, req.(*GetRecordListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetRecordStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetRecordStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetRecordStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetRecordStatus(ctx, req.(*GetRecordStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_Transcription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).Transcription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_Transcription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).Transcription(ctx, req.(*TranscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetTransData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetTransData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetTransData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetTransData(ctx, req.(*GetTransDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_SummaryDepth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryDepthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).SummaryDepth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_SummaryDepth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).SummaryDepth(ctx, req.(*SummaryDepthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetSummaryDepth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummaryDepthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetSummaryDepth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetSummaryDepth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetSummaryDepth(ctx, req.(*GetSummaryDepthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_UpdateSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSummaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).UpdateSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_UpdateSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).UpdateSummary(ctx, req.(*UpdateSummaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioService_ServiceDesc is the grpc.ServiceDesc for AudioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.audio.v1.AudioService",
	HandlerType: (*AudioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _AudioService_CreateRecord_Handler,
		},
		{
			MethodName: "GetRecordList",
			Handler:    _AudioService_GetRecordList_Handler,
		},
		{
			MethodName: "GetRecordStatus",
			Handler:    _AudioService_GetRecordStatus_Handler,
		},
		{
			MethodName: "Transcription",
			Handler:    _AudioService_Transcription_Handler,
		},
		{
			MethodName: "GetTransData",
			Handler:    _AudioService_GetTransData_Handler,
		},
		{
			MethodName: "SummaryDepth",
			Handler:    _AudioService_SummaryDepth_Handler,
		},
		{
			MethodName: "GetSummaryDepth",
			Handler:    _AudioService_GetSummaryDepth_Handler,
		},
		{
			MethodName: "UpdateSummary",
			Handler:    _AudioService_UpdateSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audio.proto",
}
