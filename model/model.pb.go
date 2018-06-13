// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	Span
*/
package model

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Span struct {
	TraceID TraceID `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3,customtype=TraceID" json:"trace_id"`
	SpanID  SpanID  `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3,customtype=SpanID" json:"span_id"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptorModel, []int{0} }

func init() {
	proto.RegisterType((*Span)(nil), "model.Span")
}

func init() { proto.RegisterFile("model.proto", fileDescriptorModel) }

var fileDescriptorModel = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x74, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xb2, 0x49,
	0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x74, 0x29, 0x15, 0x70, 0xb1, 0x04, 0x17, 0x24,
	0xe6, 0x09, 0x99, 0x72, 0x71, 0x94, 0x14, 0x25, 0x26, 0xa7, 0xc6, 0x67, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x38, 0x49, 0x9d, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0x3c, 0x7b, 0x08, 0x48,
	0xdc, 0xd3, 0xe5, 0x11, 0x82, 0x19, 0xc4, 0x0e, 0x56, 0xeb, 0x99, 0x22, 0x64, 0xc8, 0xc5, 0x5e,
	0x5c, 0x90, 0x98, 0x07, 0xd2, 0xc5, 0x04, 0xd6, 0x25, 0x01, 0xd5, 0xc5, 0x06, 0x32, 0x15, 0xac,
	0x09, 0xca, 0x0a, 0x62, 0x03, 0x29, 0xf4, 0x4c, 0x49, 0x62, 0x03, 0x5b, 0x6c, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0xcd, 0xcb, 0x3d, 0xbd, 0x00, 0x00, 0x00,
}