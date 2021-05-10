// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: optionsserver/pb/options.proto

package optionsserver

import (
	_ "github.com/danielvladco/go-proto-gql/pb"
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string    `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`          // must be required
	Foo     *Foo2     `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`                // must be required
	Float   []float32 `protobuf:"fixed32,3,rep,packed,name=float,proto3" json:"float,omitempty"`   // must be required because its greater than 0
	String2 string    `protobuf:"bytes,4,opt,name=string2,proto3" json:"string2,omitempty"`        // simple
	Foo2    *Foo2     `protobuf:"bytes,5,opt,name=foo2,proto3" json:"foo2,omitempty"`              // simple
	Float2  []float32 `protobuf:"fixed32,6,rep,packed,name=float2,proto3" json:"float2,omitempty"` // simple
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionsserver_pb_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_optionsserver_pb_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_optionsserver_pb_options_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Data) GetFoo() *Foo2 {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *Data) GetFloat() []float32 {
	if x != nil {
		return x.Float
	}
	return nil
}

func (x *Data) GetString2() string {
	if x != nil {
		return x.String2
	}
	return ""
}

func (x *Data) GetFoo2() *Foo2 {
	if x != nil {
		return x.Foo2
	}
	return nil
}

func (x *Data) GetFloat2() []float32 {
	if x != nil {
		return x.Float2
	}
	return nil
}

type Foo2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param1 string `protobuf:"bytes,1,opt,name=param1,proto3" json:"param1,omitempty"`
}

func (x *Foo2) Reset() {
	*x = Foo2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionsserver_pb_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo2) ProtoMessage() {}

func (x *Foo2) ProtoReflect() protoreflect.Message {
	mi := &file_optionsserver_pb_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo2.ProtoReflect.Descriptor instead.
func (*Foo2) Descriptor() ([]byte, []int) {
	return file_optionsserver_pb_options_proto_rawDescGZIP(), []int{1}
}

func (x *Foo2) GetParam1() string {
	if x != nil {
		return x.Param1
	}
	return ""
}

var File_optionsserver_pb_options_proto protoreflect.FileDescriptor

var file_optionsserver_pb_options_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a,
	0x10, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xce, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xb2, 0xe0, 0x1f, 0x02,
	0x08, 0x01, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x03, 0x66, 0x6f,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6f, 0x32, 0x42, 0x06, 0xb2, 0xe0,
	0x1f, 0x02, 0x08, 0x01, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x1c, 0x0a, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x42, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01,
	0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x32, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x46, 0x6f, 0x6f, 0x32, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x32, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x32, 0x22, 0x1e, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x31, 0x32, 0xaf, 0x06, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x07, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x07, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x32, 0x12, 0x13,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x31, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0,
	0x1f, 0x02, 0x08, 0x02, 0x12, 0x3a, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72, 0x79, 0x32, 0x12, 0x13,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x02,
	0x12, 0x35, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01,
	0x12, 0x37, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x31, 0x12, 0x13, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x11, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x31, 0x12, 0x13,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x02,
	0x28, 0x01, 0x12, 0x47, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x32, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x11, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x33,
	0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02,
	0x08, 0x02, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62,
	0x32, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f,
	0x02, 0x08, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f,
	0x02, 0x10, 0x01, 0x12, 0x3f, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0d, 0xb2, 0xe0, 0x1f, 0x09, 0x1a, 0x07, 0x6e, 0x65, 0x77,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0x75, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33,
	0x0a, 0x07, 0x4e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x1a, 0x00, 0x32, 0xed, 0x01, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x12,
	0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x32, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a,
	0x07, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x30, 0x01, 0x1a, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x02, 0x42, 0x4d, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c,
	0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x71, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_optionsserver_pb_options_proto_rawDescOnce sync.Once
	file_optionsserver_pb_options_proto_rawDescData = file_optionsserver_pb_options_proto_rawDesc
)

func file_optionsserver_pb_options_proto_rawDescGZIP() []byte {
	file_optionsserver_pb_options_proto_rawDescOnce.Do(func() {
		file_optionsserver_pb_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_optionsserver_pb_options_proto_rawDescData)
	})
	return file_optionsserver_pb_options_proto_rawDescData
}

var file_optionsserver_pb_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_optionsserver_pb_options_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: optionsserver.Data
	(*Foo2)(nil), // 1: optionsserver.Foo2
}
var file_optionsserver_pb_options_proto_depIdxs = []int32{
	1,  // 0: optionsserver.Data.foo:type_name -> optionsserver.Foo2
	1,  // 1: optionsserver.Data.foo2:type_name -> optionsserver.Foo2
	0,  // 2: optionsserver.Service.Mutate1:input_type -> optionsserver.Data
	0,  // 3: optionsserver.Service.Mutate2:input_type -> optionsserver.Data
	0,  // 4: optionsserver.Service.Query1:input_type -> optionsserver.Data
	0,  // 5: optionsserver.Service.Query2:input_type -> optionsserver.Data
	0,  // 6: optionsserver.Service.Publish:input_type -> optionsserver.Data
	0,  // 7: optionsserver.Service.Subscribe:input_type -> optionsserver.Data
	0,  // 8: optionsserver.Service.PubSub1:input_type -> optionsserver.Data
	0,  // 9: optionsserver.Service.InvalidSubscribe1:input_type -> optionsserver.Data
	0,  // 10: optionsserver.Service.InvalidSubscribe2:input_type -> optionsserver.Data
	0,  // 11: optionsserver.Service.InvalidSubscribe3:input_type -> optionsserver.Data
	0,  // 12: optionsserver.Service.PubSub2:input_type -> optionsserver.Data
	0,  // 13: optionsserver.Service.Ignore:input_type -> optionsserver.Data
	0,  // 14: optionsserver.Service.Name:input_type -> optionsserver.Data
	0,  // 15: optionsserver.Test.Name:input_type -> optionsserver.Data
	0,  // 16: optionsserver.Test.NewName:input_type -> optionsserver.Data
	0,  // 17: optionsserver.Query.Query1:input_type -> optionsserver.Data
	0,  // 18: optionsserver.Query.Query2:input_type -> optionsserver.Data
	0,  // 19: optionsserver.Query.Mutate1:input_type -> optionsserver.Data
	0,  // 20: optionsserver.Query.Subscribe:input_type -> optionsserver.Data
	0,  // 21: optionsserver.Service.Mutate1:output_type -> optionsserver.Data
	0,  // 22: optionsserver.Service.Mutate2:output_type -> optionsserver.Data
	0,  // 23: optionsserver.Service.Query1:output_type -> optionsserver.Data
	0,  // 24: optionsserver.Service.Query2:output_type -> optionsserver.Data
	0,  // 25: optionsserver.Service.Publish:output_type -> optionsserver.Data
	0,  // 26: optionsserver.Service.Subscribe:output_type -> optionsserver.Data
	0,  // 27: optionsserver.Service.PubSub1:output_type -> optionsserver.Data
	0,  // 28: optionsserver.Service.InvalidSubscribe1:output_type -> optionsserver.Data
	0,  // 29: optionsserver.Service.InvalidSubscribe2:output_type -> optionsserver.Data
	0,  // 30: optionsserver.Service.InvalidSubscribe3:output_type -> optionsserver.Data
	0,  // 31: optionsserver.Service.PubSub2:output_type -> optionsserver.Data
	0,  // 32: optionsserver.Service.Ignore:output_type -> optionsserver.Data
	0,  // 33: optionsserver.Service.Name:output_type -> optionsserver.Data
	0,  // 34: optionsserver.Test.Name:output_type -> optionsserver.Data
	0,  // 35: optionsserver.Test.NewName:output_type -> optionsserver.Data
	0,  // 36: optionsserver.Query.Query1:output_type -> optionsserver.Data
	0,  // 37: optionsserver.Query.Query2:output_type -> optionsserver.Data
	0,  // 38: optionsserver.Query.Mutate1:output_type -> optionsserver.Data
	0,  // 39: optionsserver.Query.Subscribe:output_type -> optionsserver.Data
	21, // [21:40] is the sub-list for method output_type
	2,  // [2:21] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_optionsserver_pb_options_proto_init() }
func file_optionsserver_pb_options_proto_init() {
	if File_optionsserver_pb_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_optionsserver_pb_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_optionsserver_pb_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo2); i {
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
			RawDescriptor: file_optionsserver_pb_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_optionsserver_pb_options_proto_goTypes,
		DependencyIndexes: file_optionsserver_pb_options_proto_depIdxs,
		MessageInfos:      file_optionsserver_pb_options_proto_msgTypes,
	}.Build()
	File_optionsserver_pb_options_proto = out.File
	file_optionsserver_pb_options_proto_rawDesc = nil
	file_optionsserver_pb_options_proto_goTypes = nil
	file_optionsserver_pb_options_proto_depIdxs = nil
}
