// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/openapi.proto

package rts

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// JSON API Error Response
//
// The standard Ory JSON API error format.
type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ErrorResponse_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResponse) GetError() *ErrorResponse_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ErrorResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Debug information
	//
	// Debug information is often not exposed to protect against leaking sensitive information.
	Debug string `protobuf:"bytes,2,opt,name=debug,proto3" json:"debug,omitempty"`
	// Further error details
	//
	// Further details about the error.
	Details map[string]string `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The error ID
	//
	// The error ID is useful when trying to identify various errors in application logic.
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// The error message
	//
	// The error's message (required).
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	// The error reason
	//
	// Reason holds a human-readable reason for the error.
	Reason string `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason,omitempty"`
	// The request ID
	//
	// The request ID is often exposed internally in order to trace
	// errors across service architectures. This is often a UUID.
	Request string `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
	// The status description
	//
	// Status holds the human-readable HTTP status code.
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ErrorResponse_Error) Reset() {
	*x = ErrorResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse_Error) ProtoMessage() {}

func (x *ErrorResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse_Error.ProtoReflect.Descriptor instead.
func (*ErrorResponse_Error) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ErrorResponse_Error) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorResponse_Error) GetDebug() string {
	if x != nil {
		return x.Debug
	}
	return ""
}

func (x *ErrorResponse_Error) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ErrorResponse_Error) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ErrorResponse_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorResponse_Error) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ErrorResponse_Error) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *ErrorResponse_Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_ory_keto_relation_tuples_v1alpha2_openapi_proto protoreflect.FileDescriptor

var file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x04, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xbf, 0x03, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x08, 0x92, 0x41, 0x05, 0x4a, 0x03, 0x34, 0x30, 0x34, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x5d, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x6f, 0x72, 0x79, 0x2e,
	0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0x92, 0x41, 0x2e, 0x4a, 0x2c, 0x22,
	0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x22, 0xe0, 0x41, 0x02, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x45, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0x92, 0x41, 0x28, 0x4a, 0x26, 0x22, 0x64, 0x37, 0x65, 0x66, 0x35, 0x34, 0x62, 0x31,
	0x2d, 0x65, 0x63, 0x31, 0x35, 0x2d, 0x34, 0x36, 0x65, 0x36, 0x2d, 0x62, 0x63, 0x63, 0x62, 0x2d,
	0x35, 0x32, 0x34, 0x62, 0x38, 0x32, 0x63, 0x30, 0x33, 0x35, 0x65, 0x36, 0x22, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x4e, 0x6f,
	0x74, 0x20, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x1a, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xd1, 0x04, 0x0a,
	0x24, 0x73, 0x68, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x3b, 0x72, 0x74, 0x73, 0xaa, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x5c,
	0x4b, 0x65, 0x74, 0x6f, 0x5c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x92, 0x41, 0x90, 0x03,
	0x12, 0x95, 0x02, 0x0a, 0x08, 0x4f, 0x52, 0x59, 0x20, 0x4b, 0x65, 0x74, 0x6f, 0x12, 0x9b, 0x01,
	0x4f, 0x72, 0x79, 0x20, 0x4b, 0x65, 0x74, 0x6f, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x20, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x65, 0x73, 0x74, 0x2d,
	0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x20, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x73, 0x20, 0x28, 0x52, 0x42, 0x41, 0x43, 0x2c, 0x20, 0x41, 0x42, 0x41, 0x43, 0x2c, 0x20, 0x41,
	0x43, 0x4c, 0x2c, 0x20, 0x41, 0x57, 0x53, 0x20, 0x49, 0x41, 0x4d, 0x20, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2c, 0x20, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x20, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x2c, 0x20, 0x2e, 0x2e, 0x2e, 0x29, 0x20, 0x76, 0x69, 0x61,
	0x20, 0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x22, 0x24, 0x0a, 0x03, 0x4f,
	0x52, 0x59, 0x12, 0x12, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x6f, 0x72, 0x79, 0x2e, 0x73, 0x68, 0x1a, 0x09, 0x68, 0x69, 0x40, 0x6f, 0x72, 0x79, 0x2e, 0x73,
	0x68, 0x2a, 0x3d, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x32, 0x2e, 0x30, 0x12,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f,
	0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x32, 0x06, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x22, 0x01, 0x2f, 0x2a, 0x02, 0x01, 0x02, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x52, 0x4b, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x44, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x32, 0x1a, 0x30,
	0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescData = file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDesc
)

func file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescData)
	})
	return file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ory_keto_relation_tuples_v1alpha2_openapi_proto_goTypes = []interface{}{
	(*ErrorResponse)(nil),       // 0: ory.keto.relation_tuples.v1alpha2.ErrorResponse
	(*ErrorResponse_Error)(nil), // 1: ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error
	nil,                         // 2: ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.DetailsEntry
}
var file_ory_keto_relation_tuples_v1alpha2_openapi_proto_depIdxs = []int32{
	1, // 0: ory.keto.relation_tuples.v1alpha2.ErrorResponse.error:type_name -> ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error
	2, // 1: ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.details:type_name -> ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.DetailsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_openapi_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_openapi_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_openapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse_Error); i {
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
			RawDescriptor: file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_openapi_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_openapi_proto_depIdxs,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_openapi_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_openapi_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_rawDesc = nil
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_openapi_proto_depIdxs = nil
}