// Code generated by protoc-gen-ankr. DO NOT EDIT.
// source: proto/resource.proto

package rbac

import (
	context "context"
	json "encoding/json"
	fmt "fmt"
	math "math"

	logger "github.com/Ankr-network/dccn-tools/logger"
	zap "github.com/Ankr-network/dccn-tools/zap"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
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

type Resource struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Source               string   `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	Memo                 string   `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	Status               int64    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Resource) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Resource) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Resource) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ResourceAddRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Memo                 string   `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAddRequest) Reset()         { *m = ResourceAddRequest{} }
func (m *ResourceAddRequest) String() string { return proto.CompactTextString(m) }
func (*ResourceAddRequest) ProtoMessage()    {}
func (*ResourceAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{1}
}

func (m *ResourceAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAddRequest.Unmarshal(m, b)
}
func (m *ResourceAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAddRequest.Marshal(b, m, deterministic)
}
func (m *ResourceAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAddRequest.Merge(m, src)
}
func (m *ResourceAddRequest) XXX_Size() int {
	return xxx_messageInfo_ResourceAddRequest.Size(m)
}
func (m *ResourceAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAddRequest proto.InternalMessageInfo

func (m *ResourceAddRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResourceAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceAddRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ResourceAddRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ResourceAddRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

type ResourceAddResponse struct {
	Rsp                  *Response `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceAddResponse) Reset()         { *m = ResourceAddResponse{} }
func (m *ResourceAddResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceAddResponse) ProtoMessage()    {}
func (*ResourceAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{2}
}

func (m *ResourceAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAddResponse.Unmarshal(m, b)
}
func (m *ResourceAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAddResponse.Marshal(b, m, deterministic)
}
func (m *ResourceAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAddResponse.Merge(m, src)
}
func (m *ResourceAddResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceAddResponse.Size(m)
}
func (m *ResourceAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAddResponse proto.InternalMessageInfo

func (m *ResourceAddResponse) GetRsp() *Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

type ResourceDelRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceDelRequest) Reset()         { *m = ResourceDelRequest{} }
func (m *ResourceDelRequest) String() string { return proto.CompactTextString(m) }
func (*ResourceDelRequest) ProtoMessage()    {}
func (*ResourceDelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{3}
}

func (m *ResourceDelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDelRequest.Unmarshal(m, b)
}
func (m *ResourceDelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDelRequest.Marshal(b, m, deterministic)
}
func (m *ResourceDelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDelRequest.Merge(m, src)
}
func (m *ResourceDelRequest) XXX_Size() int {
	return xxx_messageInfo_ResourceDelRequest.Size(m)
}
func (m *ResourceDelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDelRequest proto.InternalMessageInfo

func (m *ResourceDelRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ResourceDelResponse struct {
	Rsp                  *Response `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceDelResponse) Reset()         { *m = ResourceDelResponse{} }
func (m *ResourceDelResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceDelResponse) ProtoMessage()    {}
func (*ResourceDelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{4}
}

func (m *ResourceDelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceDelResponse.Unmarshal(m, b)
}
func (m *ResourceDelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceDelResponse.Marshal(b, m, deterministic)
}
func (m *ResourceDelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDelResponse.Merge(m, src)
}
func (m *ResourceDelResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceDelResponse.Size(m)
}
func (m *ResourceDelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDelResponse proto.InternalMessageInfo

func (m *ResourceDelResponse) GetRsp() *Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

type ResourceModRequest struct {
	Rc                   *Resource `protobuf:"bytes,1,opt,name=rc,proto3" json:"rc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceModRequest) Reset()         { *m = ResourceModRequest{} }
func (m *ResourceModRequest) String() string { return proto.CompactTextString(m) }
func (*ResourceModRequest) ProtoMessage()    {}
func (*ResourceModRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{5}
}

func (m *ResourceModRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceModRequest.Unmarshal(m, b)
}
func (m *ResourceModRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceModRequest.Marshal(b, m, deterministic)
}
func (m *ResourceModRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceModRequest.Merge(m, src)
}
func (m *ResourceModRequest) XXX_Size() int {
	return xxx_messageInfo_ResourceModRequest.Size(m)
}
func (m *ResourceModRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceModRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceModRequest proto.InternalMessageInfo

func (m *ResourceModRequest) GetRc() *Resource {
	if m != nil {
		return m.Rc
	}
	return nil
}

type ResourceModResponse struct {
	Rsp                  *Response `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceModResponse) Reset()         { *m = ResourceModResponse{} }
func (m *ResourceModResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceModResponse) ProtoMessage()    {}
func (*ResourceModResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{6}
}

func (m *ResourceModResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceModResponse.Unmarshal(m, b)
}
func (m *ResourceModResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceModResponse.Marshal(b, m, deterministic)
}
func (m *ResourceModResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceModResponse.Merge(m, src)
}
func (m *ResourceModResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceModResponse.Size(m)
}
func (m *ResourceModResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceModResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceModResponse proto.InternalMessageInfo

func (m *ResourceModResponse) GetRsp() *Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

type ResourceQryRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceQryRequest) Reset()         { *m = ResourceQryRequest{} }
func (m *ResourceQryRequest) String() string { return proto.CompactTextString(m) }
func (*ResourceQryRequest) ProtoMessage()    {}
func (*ResourceQryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{7}
}

func (m *ResourceQryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceQryRequest.Unmarshal(m, b)
}
func (m *ResourceQryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceQryRequest.Marshal(b, m, deterministic)
}
func (m *ResourceQryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceQryRequest.Merge(m, src)
}
func (m *ResourceQryRequest) XXX_Size() int {
	return xxx_messageInfo_ResourceQryRequest.Size(m)
}
func (m *ResourceQryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceQryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceQryRequest proto.InternalMessageInfo

func (m *ResourceQryRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ResourceQryRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ResourceQryRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ResourceQryResponse struct {
	Rsp                  *Response   `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	Data                 []*Resource `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Total                int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResourceQryResponse) Reset()         { *m = ResourceQryResponse{} }
func (m *ResourceQryResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceQryResponse) ProtoMessage()    {}
func (*ResourceQryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{8}
}

func (m *ResourceQryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceQryResponse.Unmarshal(m, b)
}
func (m *ResourceQryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceQryResponse.Marshal(b, m, deterministic)
}
func (m *ResourceQryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceQryResponse.Merge(m, src)
}
func (m *ResourceQryResponse) XXX_Size() int {
	return xxx_messageInfo_ResourceQryResponse.Size(m)
}
func (m *ResourceQryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceQryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceQryResponse proto.InternalMessageInfo

func (m *ResourceQryResponse) GetRsp() *Response {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func (m *ResourceQryResponse) GetData() []*Resource {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResourceQryResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Resource)(nil), "rbac.Resource")
	proto.RegisterType((*ResourceAddRequest)(nil), "rbac.ResourceAddRequest")
	proto.RegisterType((*ResourceAddResponse)(nil), "rbac.ResourceAddResponse")
	proto.RegisterType((*ResourceDelRequest)(nil), "rbac.ResourceDelRequest")
	proto.RegisterType((*ResourceDelResponse)(nil), "rbac.ResourceDelResponse")
	proto.RegisterType((*ResourceModRequest)(nil), "rbac.ResourceModRequest")
	proto.RegisterType((*ResourceModResponse)(nil), "rbac.ResourceModResponse")
	proto.RegisterType((*ResourceQryRequest)(nil), "rbac.ResourceQryRequest")
	proto.RegisterType((*ResourceQryResponse)(nil), "rbac.ResourceQryResponse")
}

func init() { proto.RegisterFile("proto/resource.proto", fileDescriptor_e41559ab98850371) }

var fileDescriptor_e41559ab98850371 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x4d, 0x26, 0x5d, 0xed, 0x14, 0x8a, 0x8c, 0xb5, 0xc4, 0xb0, 0xd4, 0x30, 0xa7, 0x52,
	0xec, 0x06, 0xab, 0x28, 0x78, 0xab, 0xf4, 0xda, 0x43, 0xe3, 0xc5, 0x43, 0x51, 0x66, 0x93, 0x21,
	0x0d, 0x26, 0x99, 0x74, 0x32, 0xd9, 0x65, 0x15, 0x2f, 0x5e, 0xfc, 0x00, 0x5e, 0xc5, 0x2f, 0xe5,
	0x07, 0x10, 0xc4, 0x0f, 0x22, 0x33, 0x93, 0x64, 0x92, 0xec, 0x8a, 0x7b, 0x9b, 0xf7, 0x4f, 0x9e,
	0xe7, 0x37, 0xf3, 0xbe, 0x04, 0x1e, 0x94, 0x9c, 0x09, 0x16, 0x70, 0x5a, 0xb1, 0x9a, 0x47, 0x74,
	0xa6, 0x42, 0xe4, 0xf0, 0x39, 0x89, 0xbc, 0x17, 0x49, 0x2a, 0x6e, 0xea, 0xf9, 0x2c, 0x62, 0x79,
	0x90, 0x2f, 0x53, 0xf1, 0x81, 0x2d, 0x83, 0x84, 0x9d, 0xaa, 0x96, 0xd3, 0x05, 0xc9, 0xd2, 0x98,
	0x08, 0xc6, 0xab, 0xa0, 0x3b, 0xea, 0xaf, 0xbd, 0x69, 0xc2, 0x58, 0x92, 0xd1, 0x80, 0x94, 0x69,
	0x40, 0x8a, 0x82, 0x09, 0x22, 0x52, 0x56, 0x54, 0x4d, 0xd5, 0x38, 0x96, 0xac, 0xa8, 0x1a, 0x47,
	0xfc, 0xdd, 0x82, 0xf7, 0xc2, 0x06, 0x02, 0xed, 0x43, 0x3b, 0x8d, 0x5d, 0xcb, 0xb7, 0x8e, 0x41,
	0x68, 0xa7, 0x31, 0x42, 0xd0, 0x11, 0xab, 0x92, 0xba, 0xb6, 0x6f, 0x1d, 0xef, 0x86, 0xea, 0x2c,
	0x73, 0x05, 0xc9, 0xa9, 0x0b, 0x74, 0x4e, 0x9e, 0xd1, 0x01, 0xdc, 0x59, 0x90, 0xac, 0xa6, 0xae,
	0xa3, 0x92, 0x3a, 0x40, 0x87, 0x70, 0xa2, 0x75, 0xdd, 0x1d, 0x95, 0x6e, 0x22, 0xa9, 0x90, 0xd3,
	0x9c, 0xb9, 0x13, 0xad, 0x20, 0xcf, 0xaa, 0x57, 0x10, 0x51, 0x57, 0xee, 0x5d, 0xe5, 0xde, 0x44,
	0xf8, 0x87, 0x05, 0x51, 0x8b, 0x77, 0x1e, 0xc7, 0x21, 0xbd, 0xad, 0x69, 0x25, 0x90, 0xd7, 0x80,
	0x49, 0xd4, 0xdd, 0xd7, 0x93, 0xdf, 0xbf, 0x1e, 0xdb, 0x6f, 0xad, 0x06, 0xd0, 0x6b, 0x00, 0xed,
	0x61, 0x4d, 0x81, 0x4e, 0x5b, 0x50, 0x30, 0x28, 0x36, 0xc0, 0x47, 0x1d, 0xb0, 0x33, 0x28, 0x8f,
	0xc1, 0x77, 0x0c, 0x38, 0x7e, 0x09, 0x1f, 0x0c, 0xf8, 0xf4, 0xe3, 0x22, 0x1f, 0x02, 0x5e, 0x95,
	0x8a, 0x6f, 0xef, 0x6c, 0x7f, 0x26, 0xc7, 0x3a, 0x6b, 0x8b, 0xa1, 0x2c, 0xe1, 0x27, 0xe6, 0x62,
	0x17, 0x34, 0x6b, 0x2f, 0x76, 0x68, 0x26, 0xa0, 0xed, 0xef, 0xdf, 0x91, 0x93, 0xe8, 0xdb, 0xa8,
	0xee, 0xad, 0x6d, 0x9e, 0x1b, 0x9b, 0x4b, 0xd6, 0xbd, 0xdf, 0x11, 0xb4, 0x79, 0xb4, 0xf6, 0x99,
	0xea, 0x0a, 0x6d, 0x1e, 0xf5, 0xed, 0xd4, 0x57, 0x5b, 0xdb, 0xdd, 0x18, 0xbb, 0x2b, 0xbe, 0xea,
	0x8d, 0xab, 0x24, 0x09, 0x1d, 0xdd, 0x4b, 0xe5, 0x64, 0xad, 0x4a, 0x3f, 0xea, 0x71, 0xf5, 0x6a,
	0x32, 0xd7, 0x8d, 0x19, 0xac, 0x8f, 0x19, 0xdf, 0x1a, 0x44, 0xe5, 0xb4, 0x2d, 0x22, 0xc2, 0xd0,
	0x89, 0x89, 0x20, 0xae, 0xed, 0x83, 0x0d, 0xb7, 0x57, 0x35, 0xb9, 0xd0, 0x82, 0x09, 0x92, 0x29,
	0x67, 0x10, 0xea, 0xe0, 0xec, 0x2b, 0x80, 0x7b, 0x6d, 0xe3, 0x1b, 0xbe, 0x40, 0xd7, 0x10, 0x9c,
	0xc7, 0x31, 0x72, 0x87, 0x12, 0x66, 0x4d, 0xbd, 0x47, 0x1b, 0x2a, 0x1a, 0x05, 0xfb, 0x5f, 0x7e,
	0xfe, 0xf9, 0x66, 0x7b, 0xf8, 0x61, 0x20, 0x5b, 0xde, 0x2f, 0xe9, 0x3c, 0x58, 0x3c, 0xed, 0x7e,
	0x07, 0xaf, 0xac, 0x13, 0xf4, 0x0e, 0x82, 0x0b, 0x9a, 0x8d, 0xd5, 0xcd, 0xae, 0x8c, 0xd5, 0x7b,
	0x7b, 0x81, 0xb1, 0x52, 0x9f, 0x9e, 0x78, 0x1b, 0xd5, 0x83, 0x4f, 0x69, 0xfc, 0x59, 0xd2, 0x5f,
	0xb2, 0x35, 0x7a, 0xb3, 0x24, 0x63, 0xfd, 0xde, 0x22, 0xb4, 0xf4, 0xde, 0xbf, 0xe9, 0xaf, 0x21,
	0xb8, 0xe2, 0xab, 0xb1, 0xba, 0xd9, 0x89, 0xb1, 0x7a, 0x6f, 0x86, 0xff, 0x7f, 0x9b, 0xf9, 0x44,
	0xfd, 0xbc, 0x9e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xe8, 0xbf, 0xc9, 0x46, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ResourceSrvClient is the client API for ResourceSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceSrvClient interface {
	Add(ctx context.Context, in *ResourceAddRequest, opts ...grpc.CallOption) (*ResourceAddResponse, error)
	Del(ctx context.Context, in *ResourceDelRequest, opts ...grpc.CallOption) (*ResourceDelResponse, error)
	Mod(ctx context.Context, in *ResourceModRequest, opts ...grpc.CallOption) (*ResourceModResponse, error)
	Qry(ctx context.Context, in *ResourceQryRequest, opts ...grpc.CallOption) (*ResourceQryResponse, error)
	Close() error
}

type resourceSrvClient struct {
	cc *grpc.ClientConn
}

// origin client method
func NewResourceSrvClient(cc *grpc.ClientConn) ResourceSrvClient {
	return &resourceSrvClient{cc}
}

// new client method
func NewAnkrResourceSrvClient(addr string) (ResourceSrvClient, error) {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &resourceSrvClient{c}, nil
}

// new client close method
func (c *resourceSrvClient) Close() error {
	return c.cc.Close()
}

func (c *resourceSrvClient) Add(ctx context.Context, in *ResourceAddRequest, opts ...grpc.CallOption) (*ResourceAddResponse, error) {
	out := new(ResourceAddResponse)
	err := c.cc.Invoke(ctx, "/rbac.ResourceSrv/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceSrvClient) Del(ctx context.Context, in *ResourceDelRequest, opts ...grpc.CallOption) (*ResourceDelResponse, error) {
	out := new(ResourceDelResponse)
	err := c.cc.Invoke(ctx, "/rbac.ResourceSrv/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceSrvClient) Mod(ctx context.Context, in *ResourceModRequest, opts ...grpc.CallOption) (*ResourceModResponse, error) {
	out := new(ResourceModResponse)
	err := c.cc.Invoke(ctx, "/rbac.ResourceSrv/Mod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceSrvClient) Qry(ctx context.Context, in *ResourceQryRequest, opts ...grpc.CallOption) (*ResourceQryResponse, error) {
	out := new(ResourceQryResponse)
	err := c.cc.Invoke(ctx, "/rbac.ResourceSrv/Qry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceSrvServer is the server API for ResourceSrv service.
type ResourceSrvServer interface {
	Add(context.Context, *ResourceAddRequest) (*ResourceAddResponse, error)
	Del(context.Context, *ResourceDelRequest) (*ResourceDelResponse, error)
	Mod(context.Context, *ResourceModRequest) (*ResourceModResponse, error)
	Qry(context.Context, *ResourceQryRequest) (*ResourceQryResponse, error)
}

// UnimplementedResourceSrvServer can be embedded to have forward compatible implementations.
type UnimplementedResourceSrvServer struct {
}

func (*UnimplementedResourceSrvServer) Add(ctx context.Context, req *ResourceAddRequest) (*ResourceAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedResourceSrvServer) Del(ctx context.Context, req *ResourceDelRequest) (*ResourceDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedResourceSrvServer) Mod(ctx context.Context, req *ResourceModRequest) (*ResourceModResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mod not implemented")
}
func (*UnimplementedResourceSrvServer) Qry(ctx context.Context, req *ResourceQryRequest) (*ResourceQryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Qry not implemented")
}

func RegisterResourceSrvServer(s *grpc.Server, srv ResourceSrvServer) {
	s.RegisterService(&_ResourceSrv_serviceDesc, srv)
}

func _ResourceSrv_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	//if exist trace id then set new span id, else set the entire id values
	m := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vs := md.Get(logger.TraceID); len(vs) > 0 {
			m[logger.TraceID] = vs[len(vs)-1]
		} else {
			m[logger.TraceID] = l.Generate().String()
		}
		if vs := md.Get(logger.SpanID); len(vs) > 0 {
			m[logger.ParentSpanID] = vs[len(vs)-1]
		} else {
			m[logger.ParentSpanID] = l.Generate().String()
		}
	}
	m[logger.SpanID] = l.Generate().String()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	if interceptor == nil {
		// output request
		reqBody, err := json.Marshal(&in)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "request"))
		} else {
			l.Info(ctx, string(reqBody), zap.String("type", "request"))
		}
		rsp, err := srv.(ResourceSrvServer).Add(ctx, in)
		// output response
		rspBody, err := json.Marshal(&rsp)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "response"))
		} else {
			l.Info(ctx, string(rspBody), zap.String("type", "response"))
		}
		return rsp, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.ResourceSrv/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceSrvServer).Add(ctx, req.(*ResourceAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceSrv_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	//if exist trace id then set new span id, else set the entire id values
	m := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vs := md.Get(logger.TraceID); len(vs) > 0 {
			m[logger.TraceID] = vs[len(vs)-1]
		} else {
			m[logger.TraceID] = l.Generate().String()
		}
		if vs := md.Get(logger.SpanID); len(vs) > 0 {
			m[logger.ParentSpanID] = vs[len(vs)-1]
		} else {
			m[logger.ParentSpanID] = l.Generate().String()
		}
	}
	m[logger.SpanID] = l.Generate().String()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	if interceptor == nil {
		// output request
		reqBody, err := json.Marshal(&in)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "request"))
		} else {
			l.Info(ctx, string(reqBody), zap.String("type", "request"))
		}
		rsp, err := srv.(ResourceSrvServer).Del(ctx, in)
		// output response
		rspBody, err := json.Marshal(&rsp)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "response"))
		} else {
			l.Info(ctx, string(rspBody), zap.String("type", "response"))
		}
		return rsp, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.ResourceSrv/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceSrvServer).Del(ctx, req.(*ResourceDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceSrv_Mod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceModRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	//if exist trace id then set new span id, else set the entire id values
	m := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vs := md.Get(logger.TraceID); len(vs) > 0 {
			m[logger.TraceID] = vs[len(vs)-1]
		} else {
			m[logger.TraceID] = l.Generate().String()
		}
		if vs := md.Get(logger.SpanID); len(vs) > 0 {
			m[logger.ParentSpanID] = vs[len(vs)-1]
		} else {
			m[logger.ParentSpanID] = l.Generate().String()
		}
	}
	m[logger.SpanID] = l.Generate().String()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	if interceptor == nil {
		// output request
		reqBody, err := json.Marshal(&in)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "request"))
		} else {
			l.Info(ctx, string(reqBody), zap.String("type", "request"))
		}
		rsp, err := srv.(ResourceSrvServer).Mod(ctx, in)
		// output response
		rspBody, err := json.Marshal(&rsp)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "response"))
		} else {
			l.Info(ctx, string(rspBody), zap.String("type", "response"))
		}
		return rsp, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.ResourceSrv/Mod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceSrvServer).Mod(ctx, req.(*ResourceModRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceSrv_Qry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceQryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	//if exist trace id then set new span id, else set the entire id values
	m := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vs := md.Get(logger.TraceID); len(vs) > 0 {
			m[logger.TraceID] = vs[len(vs)-1]
		} else {
			m[logger.TraceID] = l.Generate().String()
		}
		if vs := md.Get(logger.SpanID); len(vs) > 0 {
			m[logger.ParentSpanID] = vs[len(vs)-1]
		} else {
			m[logger.ParentSpanID] = l.Generate().String()
		}
	}
	m[logger.SpanID] = l.Generate().String()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	if interceptor == nil {
		// output request
		reqBody, err := json.Marshal(&in)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "request"))
		} else {
			l.Info(ctx, string(reqBody), zap.String("type", "request"))
		}
		rsp, err := srv.(ResourceSrvServer).Qry(ctx, in)
		// output response
		rspBody, err := json.Marshal(&rsp)
		if err != nil {
			l.Error(ctx, err.Error(), zap.String("type", "response"))
		} else {
			l.Info(ctx, string(rspBody), zap.String("type", "response"))
		}
		return rsp, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.ResourceSrv/Qry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceSrvServer).Qry(ctx, req.(*ResourceQryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rbac.ResourceSrv",
	HandlerType: (*ResourceSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ResourceSrv_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _ResourceSrv_Del_Handler,
		},
		{
			MethodName: "Mod",
			Handler:    _ResourceSrv_Mod_Handler,
		},
		{
			MethodName: "Qry",
			Handler:    _ResourceSrv_Qry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/resource.proto",
}
