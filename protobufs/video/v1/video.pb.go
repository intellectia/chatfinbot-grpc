// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: video/v1/video.proto

package videopb

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

// Message for a Video Record
type VideoRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                     // Primary key
	DocId       string `protobuf:"bytes,2,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty"`                   // Document ID
	SessionId   string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`       // Session ID
	JobId       string `protobuf:"bytes,4,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`                   // Job ID
	Summary     string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`                            // Summary
	DeepSummary string `protobuf:"bytes,6,opt,name=deep_summary,json=deepSummary,proto3" json:"deep_summary,omitempty"` // Deep Summary
	UserId      string `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // User ID
	Subtitles   string `protobuf:"bytes,8,opt,name=subtitles,proto3" json:"subtitles,omitempty"`                        // Subtitles
	Url         string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`                                    // URL of the video
	Title       string `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`                               // Title of the video
	JobStatus   int32  `protobuf:"varint,11,opt,name=job_status,json=jobStatus,proto3" json:"job_status,omitempty"`     // Job status
	Thumbnail   string `protobuf:"bytes,12,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`                       // Thumbnail URL
}

func (x *VideoRecord) Reset() {
	*x = VideoRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRecord) ProtoMessage() {}

func (x *VideoRecord) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRecord.ProtoReflect.Descriptor instead.
func (*VideoRecord) Descriptor() ([]byte, []int) {
	return file_video_v1_video_proto_rawDescGZIP(), []int{0}
}

func (x *VideoRecord) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoRecord) GetDocId() string {
	if x != nil {
		return x.DocId
	}
	return ""
}

func (x *VideoRecord) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *VideoRecord) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *VideoRecord) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *VideoRecord) GetDeepSummary() string {
	if x != nil {
		return x.DeepSummary
	}
	return ""
}

func (x *VideoRecord) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VideoRecord) GetSubtitles() string {
	if x != nil {
		return x.Subtitles
	}
	return ""
}

func (x *VideoRecord) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoRecord) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoRecord) GetJobStatus() int32 {
	if x != nil {
		return x.JobStatus
	}
	return 0
}

func (x *VideoRecord) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

type GetVideoRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"` // Session ID to fetch the video record
}

func (x *GetVideoRecordRequest) Reset() {
	*x = GetVideoRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRecordRequest) ProtoMessage() {}

func (x *GetVideoRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRecordRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRecordRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_video_proto_rawDescGZIP(), []int{1}
}

func (x *GetVideoRecordRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type UpdateVideoRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string       `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`       // Session ID of the video record to be updated
	VideoRecord *VideoRecord `protobuf:"bytes,2,opt,name=video_record,json=videoRecord,proto3" json:"video_record,omitempty"` // Video record data to be updated
}

func (x *UpdateVideoRecordRequest) Reset() {
	*x = UpdateVideoRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVideoRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVideoRecordRequest) ProtoMessage() {}

func (x *UpdateVideoRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVideoRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateVideoRecordRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_video_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateVideoRecordRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *UpdateVideoRecordRequest) GetVideoRecord() *VideoRecord {
	if x != nil {
		return x.VideoRecord
	}
	return nil
}

type VideoRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoRecord *VideoRecord `protobuf:"bytes,1,opt,name=video_record,json=videoRecord,proto3" json:"video_record,omitempty"` // The requested or updated video record
	Ret         int32        `protobuf:"varint,2,opt,name=ret,proto3" json:"ret,omitempty"`                                   // Status of the request (e.g., 0 for success)
	Msg         string       `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`                                    // Additional message information
}

func (x *VideoRecordResponse) Reset() {
	*x = VideoRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRecordResponse) ProtoMessage() {}

func (x *VideoRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRecordResponse.ProtoReflect.Descriptor instead.
func (*VideoRecordResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_video_proto_rawDescGZIP(), []int{3}
}

func (x *VideoRecordResponse) GetVideoRecord() *VideoRecord {
	if x != nil {
		return x.VideoRecord
	}
	return nil
}

func (x *VideoRecordResponse) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *VideoRecordResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_video_v1_video_proto protoreflect.FileDescriptor

var file_video_v1_video_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62,
	0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02, 0x0a, 0x0b, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x6f, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x63, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x65, 0x70, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x65, 0x70, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x22,
	0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66,
	0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x7e, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f,
	0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xbe, 0x02, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x66, 0x69,
	0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x7b, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x66, 0x69, 0x6e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x0c, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1b, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x7b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_v1_video_proto_rawDescOnce sync.Once
	file_video_v1_video_proto_rawDescData = file_video_v1_video_proto_rawDesc
)

func file_video_v1_video_proto_rawDescGZIP() []byte {
	file_video_v1_video_proto_rawDescOnce.Do(func() {
		file_video_v1_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_v1_video_proto_rawDescData)
	})
	return file_video_v1_video_proto_rawDescData
}

var file_video_v1_video_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_video_v1_video_proto_goTypes = []interface{}{
	(*VideoRecord)(nil),              // 0: chatfinbot.video.v1.VideoRecord
	(*GetVideoRecordRequest)(nil),    // 1: chatfinbot.video.v1.GetVideoRecordRequest
	(*UpdateVideoRecordRequest)(nil), // 2: chatfinbot.video.v1.UpdateVideoRecordRequest
	(*VideoRecordResponse)(nil),      // 3: chatfinbot.video.v1.VideoRecordResponse
}
var file_video_v1_video_proto_depIdxs = []int32{
	0, // 0: chatfinbot.video.v1.UpdateVideoRecordRequest.video_record:type_name -> chatfinbot.video.v1.VideoRecord
	0, // 1: chatfinbot.video.v1.VideoRecordResponse.video_record:type_name -> chatfinbot.video.v1.VideoRecord
	1, // 2: chatfinbot.video.v1.VideoService.GetVideoRecord:input_type -> chatfinbot.video.v1.GetVideoRecordRequest
	2, // 3: chatfinbot.video.v1.VideoService.UpdateVideoRecord:input_type -> chatfinbot.video.v1.UpdateVideoRecordRequest
	3, // 4: chatfinbot.video.v1.VideoService.GetVideoRecord:output_type -> chatfinbot.video.v1.VideoRecordResponse
	3, // 5: chatfinbot.video.v1.VideoService.UpdateVideoRecord:output_type -> chatfinbot.video.v1.VideoRecordResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_video_v1_video_proto_init() }
func file_video_v1_video_proto_init() {
	if File_video_v1_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_v1_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoRecord); i {
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
		file_video_v1_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRecordRequest); i {
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
		file_video_v1_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVideoRecordRequest); i {
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
		file_video_v1_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoRecordResponse); i {
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
			RawDescriptor: file_video_v1_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_v1_video_proto_goTypes,
		DependencyIndexes: file_video_v1_video_proto_depIdxs,
		MessageInfos:      file_video_v1_video_proto_msgTypes,
	}.Build()
	File_video_v1_video_proto = out.File
	file_video_v1_video_proto_rawDesc = nil
	file_video_v1_video_proto_goTypes = nil
	file_video_v1_video_proto_depIdxs = nil
}
