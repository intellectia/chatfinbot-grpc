// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: errcode.proto

package commonpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RetCode int32

const (
	// 通用错误码
	RetCode_SUCCEED                                 RetCode = 0
	RetCode_ERR_PARAMETER_WRONG                     RetCode = 10000 // 请求参数错误
	RetCode_ERR_AUTH_FAILED                         RetCode = 10001 // 鉴权失败
	RetCode_ERR_PERMISSION_DENIED                   RetCode = 10002 // 没有权限
	RetCode_ERR_SYSTEM_BUSY                         RetCode = 10003 // 系统错误
	RetCode_ERR_BACKEND_RPC_FAILED                  RetCode = 10004 // 后端访问错误
	RetCode_ERR_REDIS_FAILED                        RetCode = 10005 // 后端访问REDIS错误
	RetCode_ERR_MYSQL_FAILED                        RetCode = 10006 // 后端访问MySQL错误
	RetCode_ERR_JSON_MARSHAL_FAILED                 RetCode = 10007 // JSON序列化失败
	RetCode_ERR_GET_USER_INFO_FAILED                RetCode = 10008 // 获取用户基本信息失败
	RetCode_ERR_UID_NIL                             RetCode = 10009 // 登录态获取UID为空
	RetCode_ERR_FREQUENCY_LIMIT                     RetCode = 10010 // 频率限制了
	RetCode_ERR_SERVER_FAILED                       RetCode = 10011 // 服务器发生一般性错误
	RetCode_ERR_ES                                  RetCode = 10012 // 操作 ES 失败
	RetCode_ERR_JSON_UNMARSHAL_FAILED               RetCode = 10013 // JSON反序列化失败
	RetCode_ERR_HTTP_REQUEST_FAILED                 RetCode = 10014 // http请求失败
	RetCode_ERR_HTTP_GET_RSP_FAILED                 RetCode = 10015 // http解包失败
	RetCode_ERR_PB_MARSHAL_FAILED                   RetCode = 10016 // Protobuf序列化失败
	RetCode_ERR_PB_UNMARSHAL_FAILED                 RetCode = 10017 // Protobuf反序列化失败
	RetCode_ERR_GET_COMM_REQ_FAILED                 RetCode = 10018 // 获取CommReq失败
	RetCode_ERR_NEW_HTTP_CLIENT_WITH_ADDRESS_FAILED RetCode = 10019 // 新建HTTP Client失败
	RetCode_ERR_HTTP_CLIENT_DO_FAILED               RetCode = 10020 // HTTPClient.Do 失败
	RetCode_ERR_HTTP_CLIENT_GET_HTTP_RSP_FAILED     RetCode = 10021 // HTTPClient GetHTTPRsp失败
	RetCode_ERR_KAFKA_FAILED                        RetCode = 10022 // kafka失败
	RetCode_ERR_USER_NOT_LOGIN                      RetCode = 10023 // 用户未登录
	RetCode_ERR_SESSION_INVALID                     RetCode = 10024 // session无效
	// 视频错误码
	RetCode_ERR_BVIDEO_FAILED RetCode = 12000 // 获取B站视频信息失败失败
	// 媒体类型超额错误
	RetCode_ERR_AUDIO_FILE_OVERLIMIT   RetCode = 11001 // 音频文件超额
	RetCode_ERR_AUDIO_LIVE_OVERLIMIT   RetCode = 11002 // 录音超额
	RetCode_ERR_VIDEO_OVERLIMIT        RetCode = 11003 // 视频超额
	RetCode_ERR_DEEP_SUMMARY_OVERLIMIT RetCode = 11004 // 深度总结超额
	RetCode_ERR_CHATDOC_OVERLIMIT      RetCode = 11005 // chatdoc超额
	RetCode_ERR_WRITING_OVERLIMIT      RetCode = 11006 // 写作超额
	// 文件上传错误
	RetCode_ERR_FILE_UPLOAD_FORMAT_LIMITED RetCode = 13001 // 文件上传文件限制
	RetCode_ERR_FILE_NUMBER_LIMITED        RetCode = 13002 // 文件数量限制
	RetCode_ERR_FILE_PAGE_SIZE_LIMITED     RetCode = 13003 // 文件页数限制
	RetCode_ERR_AUDIO_DURATION_LIMITED     RetCode = 13004 // Audio时长限制
)

// Enum value maps for RetCode.
var (
	RetCode_name = map[int32]string{
		0:     "SUCCEED",
		10000: "ERR_PARAMETER_WRONG",
		10001: "ERR_AUTH_FAILED",
		10002: "ERR_PERMISSION_DENIED",
		10003: "ERR_SYSTEM_BUSY",
		10004: "ERR_BACKEND_RPC_FAILED",
		10005: "ERR_REDIS_FAILED",
		10006: "ERR_MYSQL_FAILED",
		10007: "ERR_JSON_MARSHAL_FAILED",
		10008: "ERR_GET_USER_INFO_FAILED",
		10009: "ERR_UID_NIL",
		10010: "ERR_FREQUENCY_LIMIT",
		10011: "ERR_SERVER_FAILED",
		10012: "ERR_ES",
		10013: "ERR_JSON_UNMARSHAL_FAILED",
		10014: "ERR_HTTP_REQUEST_FAILED",
		10015: "ERR_HTTP_GET_RSP_FAILED",
		10016: "ERR_PB_MARSHAL_FAILED",
		10017: "ERR_PB_UNMARSHAL_FAILED",
		10018: "ERR_GET_COMM_REQ_FAILED",
		10019: "ERR_NEW_HTTP_CLIENT_WITH_ADDRESS_FAILED",
		10020: "ERR_HTTP_CLIENT_DO_FAILED",
		10021: "ERR_HTTP_CLIENT_GET_HTTP_RSP_FAILED",
		10022: "ERR_KAFKA_FAILED",
		10023: "ERR_USER_NOT_LOGIN",
		10024: "ERR_SESSION_INVALID",
		12000: "ERR_BVIDEO_FAILED",
		11001: "ERR_AUDIO_FILE_OVERLIMIT",
		11002: "ERR_AUDIO_LIVE_OVERLIMIT",
		11003: "ERR_VIDEO_OVERLIMIT",
		11004: "ERR_DEEP_SUMMARY_OVERLIMIT",
		11005: "ERR_CHATDOC_OVERLIMIT",
		11006: "ERR_WRITING_OVERLIMIT",
		13001: "ERR_FILE_UPLOAD_FORMAT_LIMITED",
		13002: "ERR_FILE_NUMBER_LIMITED",
		13003: "ERR_FILE_PAGE_SIZE_LIMITED",
		13004: "ERR_AUDIO_DURATION_LIMITED",
	}
	RetCode_value = map[string]int32{
		"SUCCEED":                                 0,
		"ERR_PARAMETER_WRONG":                     10000,
		"ERR_AUTH_FAILED":                         10001,
		"ERR_PERMISSION_DENIED":                   10002,
		"ERR_SYSTEM_BUSY":                         10003,
		"ERR_BACKEND_RPC_FAILED":                  10004,
		"ERR_REDIS_FAILED":                        10005,
		"ERR_MYSQL_FAILED":                        10006,
		"ERR_JSON_MARSHAL_FAILED":                 10007,
		"ERR_GET_USER_INFO_FAILED":                10008,
		"ERR_UID_NIL":                             10009,
		"ERR_FREQUENCY_LIMIT":                     10010,
		"ERR_SERVER_FAILED":                       10011,
		"ERR_ES":                                  10012,
		"ERR_JSON_UNMARSHAL_FAILED":               10013,
		"ERR_HTTP_REQUEST_FAILED":                 10014,
		"ERR_HTTP_GET_RSP_FAILED":                 10015,
		"ERR_PB_MARSHAL_FAILED":                   10016,
		"ERR_PB_UNMARSHAL_FAILED":                 10017,
		"ERR_GET_COMM_REQ_FAILED":                 10018,
		"ERR_NEW_HTTP_CLIENT_WITH_ADDRESS_FAILED": 10019,
		"ERR_HTTP_CLIENT_DO_FAILED":               10020,
		"ERR_HTTP_CLIENT_GET_HTTP_RSP_FAILED":     10021,
		"ERR_KAFKA_FAILED":                        10022,
		"ERR_USER_NOT_LOGIN":                      10023,
		"ERR_SESSION_INVALID":                     10024,
		"ERR_BVIDEO_FAILED":                       12000,
		"ERR_AUDIO_FILE_OVERLIMIT":                11001,
		"ERR_AUDIO_LIVE_OVERLIMIT":                11002,
		"ERR_VIDEO_OVERLIMIT":                     11003,
		"ERR_DEEP_SUMMARY_OVERLIMIT":              11004,
		"ERR_CHATDOC_OVERLIMIT":                   11005,
		"ERR_WRITING_OVERLIMIT":                   11006,
		"ERR_FILE_UPLOAD_FORMAT_LIMITED":          13001,
		"ERR_FILE_NUMBER_LIMITED":                 13002,
		"ERR_FILE_PAGE_SIZE_LIMITED":              13003,
		"ERR_AUDIO_DURATION_LIMITED":              13004,
	}
)

func (x RetCode) Enum() *RetCode {
	p := new(RetCode)
	*p = x
	return p
}

func (x RetCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errcode_proto_enumTypes[0].Descriptor()
}

func (RetCode) Type() protoreflect.EnumType {
	return &file_errcode_proto_enumTypes[0]
}

func (x RetCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetCode.Descriptor instead.
func (RetCode) EnumDescriptor() ([]byte, []int) {
	return file_errcode_proto_rawDescGZIP(), []int{0}
}

var File_errcode_proto protoreflect.FileDescriptor

var file_errcode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x95, 0x08, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f,
	0x57, 0x52, 0x4f, 0x4e, 0x47, 0x10, 0x90, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x91, 0x4e, 0x12, 0x1a,
	0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x52,
	0x52, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x93, 0x4e,
	0x12, 0x1b, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f,
	0x52, 0x50, 0x43, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x94, 0x4e, 0x12, 0x15, 0x0a,
	0x10, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x95, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x59, 0x53, 0x51,
	0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x96, 0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x45,
	0x52, 0x52, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x52, 0x53, 0x48, 0x41, 0x4c, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x52, 0x52,
	0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x98, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x45, 0x52, 0x52, 0x5f,
	0x55, 0x49, 0x44, 0x5f, 0x4e, 0x49, 0x4c, 0x10, 0x99, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x52,
	0x52, 0x5f, 0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x9a, 0x4e, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x9b, 0x4e, 0x12, 0x0b, 0x0a, 0x06,
	0x45, 0x52, 0x52, 0x5f, 0x45, 0x53, 0x10, 0x9c, 0x4e, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x52, 0x52,
	0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x52, 0x53, 0x48, 0x41, 0x4c, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x9d, 0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x45, 0x52, 0x52,
	0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x9e, 0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x48,
	0x54, 0x54, 0x50, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x53, 0x50, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x9f, 0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x42, 0x5f,
	0x4d, 0x41, 0x52, 0x53, 0x48, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa0,
	0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x42, 0x5f, 0x55, 0x4e, 0x4d, 0x41,
	0x52, 0x53, 0x48, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa1, 0x4e, 0x12,
	0x1c, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x5f,
	0x52, 0x45, 0x51, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa2, 0x4e, 0x12, 0x2c, 0x0a,
	0x27, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53,
	0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa3, 0x4e, 0x12, 0x1e, 0x0a, 0x19, 0x45,
	0x52, 0x52, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44,
	0x4f, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa4, 0x4e, 0x12, 0x28, 0x0a, 0x23, 0x45,
	0x52, 0x52, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x47,
	0x45, 0x54, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52, 0x53, 0x50, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0xa5, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x4b, 0x41, 0x46,
	0x4b, 0x41, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa6, 0x4e, 0x12, 0x17, 0x0a, 0x12,
	0x45, 0x52, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x10, 0xa7, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xa8, 0x4e, 0x12,
	0x16, 0x0a, 0x11, 0x45, 0x52, 0x52, 0x5f, 0x42, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0xe0, 0x5d, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x5f, 0x41,
	0x55, 0x44, 0x49, 0x4f, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x10, 0xf9, 0x55, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4d,
	0x49, 0x54, 0x10, 0xfa, 0x55, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x56, 0x49, 0x44,
	0x45, 0x4f, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xfb, 0x55, 0x12,
	0x1f, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x45, 0x45, 0x50, 0x5f, 0x53, 0x55, 0x4d, 0x4d,
	0x41, 0x52, 0x59, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xfc, 0x55,
	0x12, 0x1a, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x44, 0x4f, 0x43, 0x5f,
	0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xfd, 0x55, 0x12, 0x1a, 0x0a, 0x15,
	0x45, 0x52, 0x52, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x56, 0x45, 0x52,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xfe, 0x55, 0x12, 0x23, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xc9, 0x65, 0x12, 0x1c, 0x0a,
	0x17, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xca, 0x65, 0x12, 0x1f, 0x0a, 0x1a, 0x45,
	0x52, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x49, 0x5a,
	0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xcb, 0x65, 0x12, 0x1f, 0x0a, 0x1a,
	0x45, 0x52, 0x52, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xcc, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_errcode_proto_rawDescOnce sync.Once
	file_errcode_proto_rawDescData = file_errcode_proto_rawDesc
)

func file_errcode_proto_rawDescGZIP() []byte {
	file_errcode_proto_rawDescOnce.Do(func() {
		file_errcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errcode_proto_rawDescData)
	})
	return file_errcode_proto_rawDescData
}

var file_errcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errcode_proto_goTypes = []interface{}{
	(RetCode)(0), // 0: chatfinbot.common.v1.RetCode
}
var file_errcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errcode_proto_init() }
func file_errcode_proto_init() {
	if File_errcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errcode_proto_goTypes,
		DependencyIndexes: file_errcode_proto_depIdxs,
		EnumInfos:         file_errcode_proto_enumTypes,
	}.Build()
	File_errcode_proto = out.File
	file_errcode_proto_rawDesc = nil
	file_errcode_proto_goTypes = nil
	file_errcode_proto_depIdxs = nil
}
