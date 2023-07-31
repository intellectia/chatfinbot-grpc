// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: file.proto

package filepb

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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileContent []byte `protobuf:"bytes,1,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
	Filename    string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

func (x *UploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashValue string `protobuf:"bytes,1,opt,name=hash_value,json=hashValue,proto3" json:"hash_value,omitempty"`
	S3Path    string `protobuf:"bytes,2,opt,name=s3_path,json=s3Path,proto3" json:"s3_path,omitempty"`
	FileName  string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetHashValue() string {
	if x != nil {
		return x.HashValue
	}
	return ""
}

func (x *UploadResponse) GetS3Path() string {
	if x != nil {
		return x.S3Path
	}
	return ""
}

func (x *UploadResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68,
	0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x61, 0x73, 0x68, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x33, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x33, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x62, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f,
	0x74, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69,
	0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x3b, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_file_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),  // 0: chatfinbot.file.v1.UploadRequest
	(*UploadResponse)(nil), // 1: chatfinbot.file.v1.UploadResponse
}
var file_file_proto_depIdxs = []int32{
	0, // 0: chatfinbot.file.v1.FileService.UploadFile:input_type -> chatfinbot.file.v1.UploadRequest
	1, // 1: chatfinbot.file.v1.FileService.UploadFile:output_type -> chatfinbot.file.v1.UploadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
