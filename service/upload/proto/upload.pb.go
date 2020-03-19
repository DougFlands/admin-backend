// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upload.proto

package go_micro_service_upload

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Resp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b94b655bd2a7e5, []int{0}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Resp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReqEntry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEntry) Reset()         { *m = ReqEntry{} }
func (m *ReqEntry) String() string { return proto.CompactTextString(m) }
func (*ReqEntry) ProtoMessage()    {}
func (*ReqEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b94b655bd2a7e5, []int{1}
}

func (m *ReqEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEntry.Unmarshal(m, b)
}
func (m *ReqEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEntry.Marshal(b, m, deterministic)
}
func (m *ReqEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEntry.Merge(m, src)
}
func (m *ReqEntry) XXX_Size() int {
	return xxx_messageInfo_ReqEntry.Size(m)
}
func (m *ReqEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEntry proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Resp)(nil), "go.micro.service.upload.Resp")
	proto.RegisterType((*ReqEntry)(nil), "go.micro.service.upload.ReqEntry")
}

func init() { proto.RegisterFile("upload.proto", fileDescriptor_91b94b655bd2a7e5) }

var fileDescriptor_91b94b655bd2a7e5 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2d, 0xc8, 0xc9,
	0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c,
	0x2e, 0xca, 0xd7, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0x48, 0x2b, 0x39, 0x70,
	0xb1, 0x04, 0xa5, 0x16, 0x17, 0x08, 0x09, 0x71, 0xb1, 0x38, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x20, 0x26, 0x48, 0x95, 0x4b, 0x62, 0x49, 0xa2, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x4f, 0x10, 0x98, 0xad, 0xc4, 0xc5, 0xc5, 0x11, 0x94, 0x5a, 0xe8, 0x9a, 0x57, 0x52, 0x54,
	0x69, 0x94, 0xc4, 0xc5, 0x1b, 0x0a, 0x36, 0x37, 0x18, 0x62, 0x8b, 0x50, 0x20, 0x17, 0x37, 0x44,
	0x00, 0x2c, 0x2f, 0xa4, 0xa8, 0x87, 0xc3, 0x1d, 0x7a, 0x30, 0x23, 0xa4, 0x64, 0xf1, 0x28, 0x29,
	0x2e, 0x50, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x77, 0x79,
	0xe0, 0x23, 0xe1, 0x00, 0x00, 0x00,
}