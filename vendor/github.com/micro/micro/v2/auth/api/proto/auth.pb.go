// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/micro/auth/api/proto/auth.proto

package go_micro_api_auth

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

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2514cc94062b9131, []int{0}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2514cc94062b9131, []int{1}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "go.micro.api.auth.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "go.micro.api.auth.VerifyResponse")
}

func init() {
	proto.RegisterFile("micro/micro/auth/api/proto/auth.proto", fileDescriptor_2514cc94062b9131)
}

var fileDescriptor_2514cc94062b9131 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x87, 0x90, 0x89, 0xa5, 0x25, 0x19, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x05, 0x45, 0xf9,
	0x25, 0x10, 0xae, 0x1e, 0x98, 0x29, 0x24, 0x98, 0x9e, 0xaf, 0x07, 0x56, 0xa3, 0x97, 0x58, 0x90,
	0xa9, 0x07, 0x92, 0x50, 0x52, 0xe5, 0xe2, 0x0d, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x0c, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x04, 0xb8, 0xf8, 0x60, 0xca, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x8d, 0xc2, 0xb9, 0x58, 0x1c, 0x4b, 0x4b, 0x32, 0x84, 0xfc, 0xb9, 0xd8, 0x20, 0x32,
	0x42, 0x0a, 0x7a, 0x18, 0xc6, 0xeb, 0xa1, 0x98, 0x2d, 0xa5, 0x88, 0x47, 0x05, 0xc4, 0x58, 0x25,
	0x86, 0x24, 0x36, 0xb0, 0x5b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x03, 0xec, 0x41,
	0xd4, 0x00, 0x00, 0x00,
}
