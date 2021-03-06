// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/serverless/triggers/v1/predicate.proto

package triggers

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Predicate struct {
	// Types that are valid to be assigned to Predicate:
	//	*Predicate_AndPredicate
	//	*Predicate_FieldValuePredicate
	Predicate            isPredicate_Predicate `protobuf_oneof:"predicate"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{0}
}

func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predicate.Unmarshal(m, b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
}
func (m *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(m, src)
}
func (m *Predicate) XXX_Size() int {
	return xxx_messageInfo_Predicate.Size(m)
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

type isPredicate_Predicate interface {
	isPredicate_Predicate()
}

type Predicate_AndPredicate struct {
	AndPredicate *AndPredicate `protobuf:"bytes,2,opt,name=and_predicate,json=andPredicate,proto3,oneof"`
}

type Predicate_FieldValuePredicate struct {
	FieldValuePredicate *FieldValuePredicate `protobuf:"bytes,4,opt,name=field_value_predicate,json=fieldValuePredicate,proto3,oneof"`
}

func (*Predicate_AndPredicate) isPredicate_Predicate() {}

func (*Predicate_FieldValuePredicate) isPredicate_Predicate() {}

func (m *Predicate) GetPredicate() isPredicate_Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *Predicate) GetAndPredicate() *AndPredicate {
	if x, ok := m.GetPredicate().(*Predicate_AndPredicate); ok {
		return x.AndPredicate
	}
	return nil
}

func (m *Predicate) GetFieldValuePredicate() *FieldValuePredicate {
	if x, ok := m.GetPredicate().(*Predicate_FieldValuePredicate); ok {
		return x.FieldValuePredicate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Predicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Predicate_AndPredicate)(nil),
		(*Predicate_FieldValuePredicate)(nil),
	}
}

type AndPredicate struct {
	Predicate            []*Predicate `protobuf:"bytes,1,rep,name=predicate,proto3" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AndPredicate) Reset()         { *m = AndPredicate{} }
func (m *AndPredicate) String() string { return proto.CompactTextString(m) }
func (*AndPredicate) ProtoMessage()    {}
func (*AndPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{1}
}

func (m *AndPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndPredicate.Unmarshal(m, b)
}
func (m *AndPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndPredicate.Marshal(b, m, deterministic)
}
func (m *AndPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndPredicate.Merge(m, src)
}
func (m *AndPredicate) XXX_Size() int {
	return xxx_messageInfo_AndPredicate.Size(m)
}
func (m *AndPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_AndPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_AndPredicate proto.InternalMessageInfo

func (m *AndPredicate) GetPredicate() []*Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

type FieldValuePredicate struct {
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*FieldValuePredicate_Exact
	//	*FieldValuePredicate_Prefix
	//	*FieldValuePredicate_Suffix
	Value                isFieldValuePredicate_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FieldValuePredicate) Reset()         { *m = FieldValuePredicate{} }
func (m *FieldValuePredicate) String() string { return proto.CompactTextString(m) }
func (*FieldValuePredicate) ProtoMessage()    {}
func (*FieldValuePredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{2}
}

func (m *FieldValuePredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValuePredicate.Unmarshal(m, b)
}
func (m *FieldValuePredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValuePredicate.Marshal(b, m, deterministic)
}
func (m *FieldValuePredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValuePredicate.Merge(m, src)
}
func (m *FieldValuePredicate) XXX_Size() int {
	return xxx_messageInfo_FieldValuePredicate.Size(m)
}
func (m *FieldValuePredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValuePredicate.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValuePredicate proto.InternalMessageInfo

func (m *FieldValuePredicate) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

type isFieldValuePredicate_Value interface {
	isFieldValuePredicate_Value()
}

type FieldValuePredicate_Exact struct {
	Exact string `protobuf:"bytes,3,opt,name=exact,proto3,oneof"`
}

type FieldValuePredicate_Prefix struct {
	Prefix string `protobuf:"bytes,8,opt,name=prefix,proto3,oneof"`
}

type FieldValuePredicate_Suffix struct {
	Suffix string `protobuf:"bytes,9,opt,name=suffix,proto3,oneof"`
}

func (*FieldValuePredicate_Exact) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Prefix) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Suffix) isFieldValuePredicate_Value() {}

func (m *FieldValuePredicate) GetValue() isFieldValuePredicate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FieldValuePredicate) GetExact() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *FieldValuePredicate) GetPrefix() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *FieldValuePredicate) GetSuffix() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Suffix); ok {
		return x.Suffix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldValuePredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldValuePredicate_Exact)(nil),
		(*FieldValuePredicate_Prefix)(nil),
		(*FieldValuePredicate_Suffix)(nil),
	}
}

func init() {
	proto.RegisterType((*Predicate)(nil), "yandex.cloud.serverless.triggers.v1.Predicate")
	proto.RegisterType((*AndPredicate)(nil), "yandex.cloud.serverless.triggers.v1.AndPredicate")
	proto.RegisterType((*FieldValuePredicate)(nil), "yandex.cloud.serverless.triggers.v1.FieldValuePredicate")
}

func init() {
	proto.RegisterFile("yandex/cloud/serverless/triggers/v1/predicate.proto", fileDescriptor_7702d63b0613d55f)
}

var fileDescriptor_7702d63b0613d55f = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xf6, 0x0f, 0x66, 0xdb, 0x1e, 0xdc, 0xa2, 0x04, 0x41, 0x28, 0xed, 0xa5, 0x97,
	0xee, 0x92, 0xf6, 0x22, 0x78, 0x32, 0x07, 0xe9, 0xc1, 0x43, 0xc9, 0x41, 0x44, 0x84, 0xb2, 0xcd,
	0x6e, 0xd2, 0x85, 0x98, 0x84, 0xcd, 0x26, 0xd4, 0x07, 0xf1, 0x7d, 0xf4, 0xe4, 0xab, 0xf4, 0x31,
	0x24, 0xd9, 0x26, 0x8d, 0xe8, 0x21, 0xb7, 0xcc, 0x7c, 0xf9, 0x7e, 0xdf, 0xcc, 0x32, 0x70, 0xf9,
	0x4e, 0x43, 0xc6, 0xf7, 0xc4, 0x0d, 0xa2, 0x94, 0x91, 0x84, 0xcb, 0x8c, 0xcb, 0x80, 0x27, 0x09,
	0x51, 0x52, 0xf8, 0x3e, 0x97, 0x09, 0xc9, 0x2c, 0x12, 0x4b, 0xce, 0x84, 0x4b, 0x15, 0xc7, 0xb1,
	0x8c, 0x54, 0x84, 0xa6, 0xda, 0x84, 0x0b, 0x13, 0x3e, 0x99, 0x70, 0x69, 0xc2, 0x99, 0x75, 0x7d,
	0xf3, 0x8b, 0x9c, 0xd1, 0x40, 0x30, 0xaa, 0x44, 0x14, 0x6a, 0xc6, 0xe4, 0x00, 0xa0, 0xb1, 0x2e,
	0xb9, 0xe8, 0x19, 0x0e, 0x69, 0xc8, 0x36, 0x55, 0x90, 0x79, 0x36, 0x06, 0xb3, 0xfe, 0xc2, 0xc2,
	0x0d, 0x92, 0xf0, 0x7d, 0xc8, 0x2a, 0xd2, 0xaa, 0xe5, 0x0c, 0x68, 0xad, 0x46, 0x21, 0xbc, 0xf4,
	0x04, 0x0f, 0xd8, 0x26, 0xa3, 0x41, 0xca, 0x6b, 0x09, 0x9d, 0x22, 0xe1, 0xb6, 0x51, 0xc2, 0x43,
	0x4e, 0x78, 0xca, 0x01, 0xf5, 0xa0, 0x91, 0xf7, 0xb7, 0x6d, 0x5f, 0x40, 0xa3, 0xca, 0x40, 0x9d,
	0xcf, 0x2f, 0x0b, 0x4c, 0x5e, 0xe1, 0xa0, 0x3e, 0x22, 0x7a, 0xac, 0xfd, 0x62, 0x82, 0x71, 0x7b,
	0xd6, 0x5f, 0xe0, 0x46, 0x63, 0x54, 0x08, 0xe7, 0x04, 0x98, 0x7c, 0x00, 0x38, 0xfa, 0x67, 0x3e,
	0x34, 0x85, 0x50, 0x2f, 0x1e, 0x53, 0xb5, 0x33, 0xc1, 0x18, 0xcc, 0x0c, 0xbb, 0x73, 0xf8, 0xb6,
	0x80, 0x63, 0x14, 0xfd, 0x35, 0x55, 0x3b, 0x74, 0x05, 0xbb, 0x7c, 0x4f, 0x5d, 0x65, 0xb6, 0x73,
	0x7d, 0xd5, 0x72, 0x74, 0x89, 0x4c, 0xd8, 0x8b, 0x25, 0xf7, 0xc4, 0xde, 0x3c, 0x3f, 0x0a, 0xc7,
	0x3a, 0x57, 0x92, 0xd4, 0xcb, 0x15, 0xa3, 0x54, 0x74, 0x6d, 0x0f, 0x61, 0xb7, 0x78, 0x63, 0xbd,
	0xb5, 0xed, 0xbc, 0xac, 0x7d, 0xa1, 0x76, 0xe9, 0x16, 0xbb, 0xd1, 0x1b, 0xd1, 0xeb, 0xcd, 0xf5,
	0x31, 0xf8, 0xd1, 0xdc, 0xe7, 0x61, 0x71, 0x07, 0xa4, 0xc1, 0xfd, 0xdd, 0x95, 0xdf, 0xdb, 0x5e,
	0xe1, 0x59, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x4c, 0xb1, 0x19, 0xb6, 0x02, 0x00, 0x00,
}
