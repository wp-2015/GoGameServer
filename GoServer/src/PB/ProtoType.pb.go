// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProtoType.proto

package protocol

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

//协议号
type MSGTYPE int32

const (
	MSGTYPE_Mt_C2S_TestMessage MSGTYPE = 0
	MSGTYPE_Mt_C2S_Join        MSGTYPE = 1
	MSGTYPE_Mt_S2C_TestMessage MSGTYPE = 10000
	MSGTYPE_Mt_S2C_MapInit     MSGTYPE = 10001
	MSGTYPE_Mt_S2C_MapRand     MSGTYPE = 10002
)

var MSGTYPE_name = map[int32]string{
	0:     "Mt_C2S_TestMessage",
	1:     "Mt_C2S_Join",
	10000: "Mt_S2C_TestMessage",
	10001: "Mt_S2C_MapInit",
	10002: "Mt_S2C_MapRand",
}

var MSGTYPE_value = map[string]int32{
	"Mt_C2S_TestMessage": 0,
	"Mt_C2S_Join":        1,
	"Mt_S2C_TestMessage": 10000,
	"Mt_S2C_MapInit":     10001,
	"Mt_S2C_MapRand":     10002,
}

func (x MSGTYPE) String() string {
	return proto.EnumName(MSGTYPE_name, int32(x))
}

func (MSGTYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f0a87896b3b6436, []int{0}
}

func init() {
	proto.RegisterEnum("protocol.MSGTYPE", MSGTYPE_name, MSGTYPE_value)
}

func init() { proto.RegisterFile("ProtoType.proto", fileDescriptor_3f0a87896b3b6436) }

var fileDescriptor_3f0a87896b3b6436 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0f, 0x28, 0xca, 0x2f,
	0xc9, 0x0f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x00, 0xb1, 0x84, 0x38, 0xc0, 0x54, 0x72, 0x7e, 0x8e,
	0x56, 0x29, 0x17, 0xbb, 0x6f, 0xb0, 0x7b, 0x48, 0x64, 0x80, 0xab, 0x90, 0x18, 0x97, 0x90, 0x6f,
	0x49, 0xbc, 0xb3, 0x51, 0x70, 0x7c, 0x48, 0x6a, 0x71, 0x89, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a,
	0xaa, 0x00, 0x83, 0x10, 0x3f, 0x17, 0x37, 0x54, 0xdc, 0x2b, 0x3f, 0x33, 0x4f, 0x80, 0x51, 0x48,
	0x1c, 0xac, 0x30, 0xd8, 0xc8, 0x19, 0x45, 0xe1, 0x04, 0x3f, 0x21, 0x61, 0x2e, 0x3e, 0xa8, 0x84,
	0x6f, 0x62, 0x81, 0x67, 0x5e, 0x66, 0x89, 0xc0, 0x44, 0x34, 0xc1, 0xa0, 0xc4, 0xbc, 0x14, 0x81,
	0x49, 0x7e, 0x49, 0x6c, 0x60, 0x07, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xee, 0x76, 0xf9,
	0x05, 0x9a, 0x00, 0x00, 0x00,
}
