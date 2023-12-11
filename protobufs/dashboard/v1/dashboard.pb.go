// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: dashboard/v1/dashboard.proto

package dashboardpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for CreateCustomActivity.
type CreateCustomActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AegisId2 string `protobuf:"bytes,2,opt,name=aegis_id2,json=aegisId2,proto3" json:"aegis_id2,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Extra    []byte `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"` // Assuming json.RawMessage can be represented as bytes.
}

func (x *CreateCustomActivityRequest) Reset() {
	*x = CreateCustomActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomActivityRequest) ProtoMessage() {}

func (x *CreateCustomActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomActivityRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomActivityRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCustomActivityRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCustomActivityRequest) GetAegisId2() string {
	if x != nil {
		return x.AegisId2
	}
	return ""
}

func (x *CreateCustomActivityRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateCustomActivityRequest) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

// Response message for CustomActivity.
type CustomActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AegisId   string `protobuf:"bytes,3,opt,name=aegis_id,json=aegisId,proto3" json:"aegis_id,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Extra     []byte `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Using int64 for Unix timestamp in milliseconds
}

func (x *CustomActivity) Reset() {
	*x = CustomActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomActivity) ProtoMessage() {}

func (x *CustomActivity) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomActivity.ProtoReflect.Descriptor instead.
func (*CustomActivity) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *CustomActivity) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomActivity) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CustomActivity) GetAegisId() string {
	if x != nil {
		return x.AegisId
	}
	return ""
}

func (x *CustomActivity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CustomActivity) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *CustomActivity) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Request message for GetAllCustomActivities.
type GetAllCustomActivitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int64  `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // Using int64 for Unix timestamp in milliseconds
	EndTime   int64  `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       // Using int64 for Unix timestamp in milliseconds
	UserType  string `protobuf:"bytes,3,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
}

func (x *GetAllCustomActivitiesRequest) Reset() {
	*x = GetAllCustomActivitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCustomActivitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCustomActivitiesRequest) ProtoMessage() {}

func (x *GetAllCustomActivitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCustomActivitiesRequest.ProtoReflect.Descriptor instead.
func (*GetAllCustomActivitiesRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCustomActivitiesRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetAllCustomActivitiesRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GetAllCustomActivitiesRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

// Response message for multiple CustomActivities.
type GetAllCustomActivitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  int32             `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` // 错误码
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`  // msg
	List []*CustomActivity `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetAllCustomActivitiesResponse) Reset() {
	*x = GetAllCustomActivitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCustomActivitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCustomActivitiesResponse) ProtoMessage() {}

func (x *GetAllCustomActivitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCustomActivitiesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCustomActivitiesResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCustomActivitiesResponse) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *GetAllCustomActivitiesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAllCustomActivitiesResponse) GetList() []*CustomActivity {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateCustomActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  int32           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` // Error code
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`  // Message
	Data *CustomActivity `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateCustomActivityResponse) Reset() {
	*x = CreateCustomActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomActivityResponse) ProtoMessage() {}

func (x *CreateCustomActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomActivityResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomActivityResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCustomActivityResponse) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *CreateCustomActivityResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateCustomActivityResponse) GetData() *CustomActivity {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_dashboard_v1_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_v1_dashboard_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x65, 0x67, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x65, 0x67, 0x69, 0x73, 0x49, 0x64, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x65, 0x67, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x65, 0x67, 0x69, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x76, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x81, 0x01, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x7f, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xd6, 0x02, 0x0a, 0x10, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x34, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62,
	0x6f, 0x74, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x12, 0xa1, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dashboard_v1_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_v1_dashboard_proto_rawDescData = file_dashboard_v1_dashboard_proto_rawDesc
)

func file_dashboard_v1_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_v1_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_v1_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_v1_dashboard_proto_rawDescData)
	})
	return file_dashboard_v1_dashboard_proto_rawDescData
}

var file_dashboard_v1_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dashboard_v1_dashboard_proto_goTypes = []interface{}{
	(*CreateCustomActivityRequest)(nil),    // 0: chatfinbot.dashboard.v1.CreateCustomActivityRequest
	(*CustomActivity)(nil),                 // 1: chatfinbot.dashboard.v1.CustomActivity
	(*GetAllCustomActivitiesRequest)(nil),  // 2: chatfinbot.dashboard.v1.GetAllCustomActivitiesRequest
	(*GetAllCustomActivitiesResponse)(nil), // 3: chatfinbot.dashboard.v1.GetAllCustomActivitiesResponse
	(*CreateCustomActivityResponse)(nil),   // 4: chatfinbot.dashboard.v1.CreateCustomActivityResponse
}
var file_dashboard_v1_dashboard_proto_depIdxs = []int32{
	1, // 0: chatfinbot.dashboard.v1.GetAllCustomActivitiesResponse.list:type_name -> chatfinbot.dashboard.v1.CustomActivity
	1, // 1: chatfinbot.dashboard.v1.CreateCustomActivityResponse.data:type_name -> chatfinbot.dashboard.v1.CustomActivity
	0, // 2: chatfinbot.dashboard.v1.DashboardService.CreateCustomActivity:input_type -> chatfinbot.dashboard.v1.CreateCustomActivityRequest
	2, // 3: chatfinbot.dashboard.v1.DashboardService.GetAllCustomActivities:input_type -> chatfinbot.dashboard.v1.GetAllCustomActivitiesRequest
	4, // 4: chatfinbot.dashboard.v1.DashboardService.CreateCustomActivity:output_type -> chatfinbot.dashboard.v1.CreateCustomActivityResponse
	3, // 5: chatfinbot.dashboard.v1.DashboardService.GetAllCustomActivities:output_type -> chatfinbot.dashboard.v1.GetAllCustomActivitiesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dashboard_v1_dashboard_proto_init() }
func file_dashboard_v1_dashboard_proto_init() {
	if File_dashboard_v1_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dashboard_v1_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomActivityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_v1_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomActivity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_v1_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCustomActivitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_v1_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCustomActivitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_v1_dashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomActivityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dashboard_v1_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_v1_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_v1_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_v1_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_v1_dashboard_proto = out.File
	file_dashboard_v1_dashboard_proto_rawDesc = nil
	file_dashboard_v1_dashboard_proto_goTypes = nil
	file_dashboard_v1_dashboard_proto_depIdxs = nil
}
