// Code generated by protoc-gen-go.
// source: adapter/denyChecker/cfg.proto
// DO NOT EDIT!

/*
Package denyChecker is a generated protocol buffer package.

It is generated from these files:
	adapter/denyChecker/cfg.proto

It has these top-level messages:
	Config
*/
package denyChecker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	// The error code to return when denying a request.
	// 	google.rpc.Status error_code = 1;
	ErrorCode int32 `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "denyChecker.Config")
}

func init() { proto.RegisterFile("adapter/denyChecker/cfg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4c, 0x49, 0x2c,
	0x28, 0x49, 0x2d, 0xd2, 0x4f, 0x49, 0xcd, 0xab, 0x74, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0xd2,
	0x4f, 0x4e, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x12, 0x56, 0x52, 0xe7,
	0x62, 0x73, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x17, 0x92, 0xe5, 0xe2, 0x4a, 0x2d, 0x2a, 0xca, 0x2f,
	0x8a, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0xe2, 0x04, 0x8b, 0x38,
	0xe7, 0xa7, 0xa4, 0x26, 0xb1, 0x81, 0x35, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x33, 0x93,
	0x4d, 0xe1, 0x5d, 0x00, 0x00, 0x00,
}