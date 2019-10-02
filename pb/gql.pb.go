// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gql.proto

package gql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_DEFAULT  Type = 0
	Type_MUTATION Type = 1
	Type_QUERY    Type = 2
)

var Type_name = map[int32]string{
	0: "DEFAULT",
	1: "MUTATION",
	2: "QUERY",
}

var Type_value = map[string]int32{
	"DEFAULT":  0,
	"MUTATION": 1,
	"QUERY":    2,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f41526d4337e9701, []int{0}
}

type Field struct {
	Required             *bool    `protobuf:"varint,1,opt,name=required" json:"required,omitempty"`
	Params               *string  `protobuf:"bytes,2,opt,name=params" json:"params,omitempty"`
	Dirs                 *string  `protobuf:"bytes,3,opt,name=dirs" json:"dirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41526d4337e9701, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetRequired() bool {
	if m != nil && m.Required != nil {
		return *m.Required
	}
	return false
}

func (m *Field) GetParams() string {
	if m != nil && m.Params != nil {
		return *m.Params
	}
	return ""
}

func (m *Field) GetDirs() string {
	if m != nil && m.Dirs != nil {
		return *m.Dirs
	}
	return ""
}

var E_RpcType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Type)(nil),
	Field:         65030,
	Name:          "gql.rpc_type",
	Tag:           "varint,65030,opt,name=rpc_type,enum=gql.Type",
	Filename:      "gql.proto",
}

var E_SvcType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Type)(nil),
	Field:         65030,
	Name:          "gql.svc_type",
	Tag:           "varint,65030,opt,name=svc_type,enum=gql.Type",
	Filename:      "gql.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Field)(nil),
	Field:         65030,
	Name:          "gql.field",
	Tag:           "bytes,65030,opt,name=field",
	Filename:      "gql.proto",
}

func init() {
	proto.RegisterEnum("gql.Type", Type_name, Type_value)
	proto.RegisterType((*Field)(nil), "gql.Field")
	proto.RegisterExtension(E_RpcType)
	proto.RegisterExtension(E_SvcType)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("gql.proto", fileDescriptor_f41526d4337e9701) }

var fileDescriptor_f41526d4337e9701 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x7e, 0xb8, 0xf6, 0x4d, 0x64, 0xe4, 0x20, 0x65, 0xa0, 0x0e, 0x4f, 0x43, 0x5d,
	0x0b, 0x3b, 0xce, 0xd3, 0xc4, 0x4d, 0x04, 0xe7, 0xb0, 0x76, 0x07, 0xbd, 0x48, 0xd7, 0x64, 0x59,
	0x20, 0x5b, 0xd2, 0x24, 0x2d, 0xf8, 0x0f, 0xf8, 0x57, 0x7b, 0x90, 0xa6, 0x53, 0x90, 0x81, 0xb7,
	0xf7, 0x5e, 0xf2, 0x3e, 0xdf, 0xef, 0xfb, 0x82, 0x47, 0x33, 0x1e, 0x48, 0x25, 0x8c, 0x40, 0x75,
	0x9a, 0xf1, 0x6e, 0x8f, 0x0a, 0x41, 0x39, 0x09, 0xed, 0x68, 0x99, 0xaf, 0x42, 0x4c, 0x74, 0xaa,
	0x98, 0x34, 0x42, 0x55, 0xdf, 0x2e, 0xe6, 0xd0, 0x9c, 0x32, 0xc2, 0x31, 0xea, 0x82, 0xab, 0x48,
	0x96, 0x33, 0x45, 0xb0, 0xef, 0xf4, 0x9c, 0xbe, 0x1b, 0xfd, 0xf6, 0xe8, 0x04, 0x0e, 0x65, 0xa2,
	0x92, 0x8d, 0xf6, 0x6b, 0x3d, 0xa7, 0xef, 0x45, 0xbb, 0x0e, 0x21, 0x68, 0x60, 0xa6, 0xb4, 0x5f,
	0xb7, 0x53, 0x5b, 0x5f, 0x5e, 0x43, 0x23, 0xfe, 0x90, 0x04, 0xb5, 0xa1, 0x75, 0x37, 0x99, 0x8e,
	0x17, 0x8f, 0x71, 0xe7, 0x00, 0x1d, 0x81, 0x3b, 0x5b, 0xc4, 0xe3, 0xf8, 0x61, 0xfe, 0xd4, 0x71,
	0x90, 0x07, 0xcd, 0xe7, 0xc5, 0x24, 0x7a, 0xed, 0xd4, 0x46, 0x53, 0x70, 0x95, 0x4c, 0xdf, 0x4d,
	0xb9, 0x71, 0x16, 0x54, 0x6e, 0x83, 0x1f, 0xb7, 0xc1, 0x8c, 0x98, 0xb5, 0xc0, 0x73, 0x69, 0x98,
	0xd8, 0x6a, 0xff, 0xf3, 0xab, 0xd4, 0x39, 0x1e, 0x7a, 0x41, 0x79, 0x65, 0x29, 0x12, 0xb5, 0x94,
	0x4c, 0xcb, 0x62, 0x74, 0x0f, 0xae, 0x2e, 0x76, 0x9c, 0xf3, 0x3d, 0xce, 0x0b, 0x51, 0x05, 0x4b,
	0xc9, 0x7f, 0x20, 0x5d, 0x54, 0xa0, 0x31, 0x34, 0x57, 0x36, 0x8f, 0xd3, 0x3d, 0x8a, 0xcd, 0xe9,
	0x2f, 0xa3, 0x3d, 0x04, 0xcb, 0xb0, 0x4f, 0x51, 0xb5, 0x79, 0x3b, 0x78, 0xbb, 0xa2, 0xcc, 0xac,
	0xf3, 0x65, 0x90, 0x8a, 0x4d, 0x88, 0x93, 0x2d, 0x23, 0xbc, 0xe0, 0x09, 0x4e, 0x45, 0x48, 0xc5,
	0xc0, 0xe2, 0x06, 0x34, 0xe3, 0xa1, 0x5c, 0xde, 0xd0, 0x8c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0x6f, 0xfa, 0xa1, 0xb4, 0x01, 0x00, 0x00,
}
