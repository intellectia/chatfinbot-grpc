// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: user.proto

package userpb

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	LoginWithEmail(ctx context.Context, in *LoginWithEmailRequest, opts ...grpc.CallOption) (*LoginWithEmailResponse, error)
	LoginWithPhone(ctx context.Context, in *LoginWithPhoneRequest, opts ...grpc.CallOption) (*LoginWithPhoneResponse, error)
	LoginWithUserName(ctx context.Context, in *LoginWithUserNameRequest, opts ...grpc.CallOption) (*LoginWithUserNameResponse, error)
	SendSMSCode(ctx context.Context, in *SendSMSCodeRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	SetInitPassword(ctx context.Context, in *SetInitPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	InitUserExtraInfo(ctx context.Context, in *InitUserExtraInfoRequest, opts ...grpc.CallOption) (*InitUserExtraInfoResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginWithEmail(ctx context.Context, in *LoginWithEmailRequest, opts ...grpc.CallOption) (*LoginWithEmailResponse, error) {
	out := new(LoginWithEmailResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/LoginWithEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginWithPhone(ctx context.Context, in *LoginWithPhoneRequest, opts ...grpc.CallOption) (*LoginWithPhoneResponse, error) {
	out := new(LoginWithPhoneResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/LoginWithPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginWithUserName(ctx context.Context, in *LoginWithUserNameRequest, opts ...grpc.CallOption) (*LoginWithUserNameResponse, error) {
	out := new(LoginWithUserNameResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/LoginWithUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendSMSCode(ctx context.Context, in *SendSMSCodeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/SendSMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetInitPassword(ctx context.Context, in *SetInitPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/SetInitPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InitUserExtraInfo(ctx context.Context, in *InitUserExtraInfoRequest, opts ...grpc.CallOption) (*InitUserExtraInfoResponse, error) {
	out := new(InitUserExtraInfoResponse)
	err := c.cc.Invoke(ctx, "/chatfinbot.user.v1.UserService/InitUserExtraInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *UserRequest) (*RegisterResponse, error)
	LoginWithEmail(context.Context, *LoginWithEmailRequest) (*LoginWithEmailResponse, error)
	LoginWithPhone(context.Context, *LoginWithPhoneRequest) (*LoginWithPhoneResponse, error)
	LoginWithUserName(context.Context, *LoginWithUserNameRequest) (*LoginWithUserNameResponse, error)
	SendSMSCode(context.Context, *SendSMSCodeRequest) (*Empty, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	SetInitPassword(context.Context, *SetInitPasswordRequest) (*Empty, error)
	InitUserExtraInfo(context.Context, *InitUserExtraInfoRequest) (*InitUserExtraInfoResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *UserRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) LoginWithEmail(context.Context, *LoginWithEmailRequest) (*LoginWithEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithEmail not implemented")
}
func (UnimplementedUserServiceServer) LoginWithPhone(context.Context, *LoginWithPhoneRequest) (*LoginWithPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithPhone not implemented")
}
func (UnimplementedUserServiceServer) LoginWithUserName(context.Context, *LoginWithUserNameRequest) (*LoginWithUserNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithUserName not implemented")
}
func (UnimplementedUserServiceServer) SendSMSCode(context.Context, *SendSMSCodeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMSCode not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) SetInitPassword(context.Context, *SetInitPasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInitPassword not implemented")
}
func (UnimplementedUserServiceServer) InitUserExtraInfo(context.Context, *InitUserExtraInfoRequest) (*InitUserExtraInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitUserExtraInfo not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/LoginWithEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginWithEmail(ctx, req.(*LoginWithEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginWithPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginWithPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/LoginWithPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginWithPhone(ctx, req.(*LoginWithPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginWithUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginWithUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/LoginWithUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginWithUserName(ctx, req.(*LoginWithUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/SendSMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendSMSCode(ctx, req.(*SendSMSCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetInitPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInitPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetInitPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/SetInitPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetInitPassword(ctx, req.(*SetInitPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InitUserExtraInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitUserExtraInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InitUserExtraInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatfinbot.user.v1.UserService/InitUserExtraInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InitUserExtraInfo(ctx, req.(*InitUserExtraInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatfinbot.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "LoginWithEmail",
			Handler:    _UserService_LoginWithEmail_Handler,
		},
		{
			MethodName: "LoginWithPhone",
			Handler:    _UserService_LoginWithPhone_Handler,
		},
		{
			MethodName: "LoginWithUserName",
			Handler:    _UserService_LoginWithUserName_Handler,
		},
		{
			MethodName: "SendSMSCode",
			Handler:    _UserService_SendSMSCode_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "SetInitPassword",
			Handler:    _UserService_SetInitPassword_Handler,
		},
		{
			MethodName: "InitUserExtraInfo",
			Handler:    _UserService_InitUserExtraInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
