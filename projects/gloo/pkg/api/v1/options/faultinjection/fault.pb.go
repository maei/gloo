// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto

package faultinjection

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RouteAbort struct {
	// Percentage of requests that should be aborted, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// This should be a standard HTTP status, i.e. 503. Defaults to 0.
	HttpStatus           uint32   `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAbort) Reset()         { *m = RouteAbort{} }
func (m *RouteAbort) String() string { return proto.CompactTextString(m) }
func (*RouteAbort) ProtoMessage()    {}
func (*RouteAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{0}
}
func (m *RouteAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAbort.Unmarshal(m, b)
}
func (m *RouteAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAbort.Marshal(b, m, deterministic)
}
func (m *RouteAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAbort.Merge(m, src)
}
func (m *RouteAbort) XXX_Size() int {
	return xxx_messageInfo_RouteAbort.Size(m)
}
func (m *RouteAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAbort.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAbort proto.InternalMessageInfo

func (m *RouteAbort) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteAbort) GetHttpStatus() uint32 {
	if m != nil {
		return m.HttpStatus
	}
	return 0
}

type RouteDelay struct {
	// Percentage of requests that should be delayed, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Fixed delay, defaulting to 0.
	FixedDelay           *time.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,proto3,stdduration" json:"fixed_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RouteDelay) Reset()         { *m = RouteDelay{} }
func (m *RouteDelay) String() string { return proto.CompactTextString(m) }
func (*RouteDelay) ProtoMessage()    {}
func (*RouteDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{1}
}
func (m *RouteDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDelay.Unmarshal(m, b)
}
func (m *RouteDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDelay.Marshal(b, m, deterministic)
}
func (m *RouteDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDelay.Merge(m, src)
}
func (m *RouteDelay) XXX_Size() int {
	return xxx_messageInfo_RouteDelay.Size(m)
}
func (m *RouteDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDelay.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDelay proto.InternalMessageInfo

func (m *RouteDelay) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteDelay) GetFixedDelay() *time.Duration {
	if m != nil {
		return m.FixedDelay
	}
	return nil
}

type RouteFaults struct {
	Abort                *RouteAbort `protobuf:"bytes,1,opt,name=abort,proto3" json:"abort,omitempty"`
	Delay                *RouteDelay `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RouteFaults) Reset()         { *m = RouteFaults{} }
func (m *RouteFaults) String() string { return proto.CompactTextString(m) }
func (*RouteFaults) ProtoMessage()    {}
func (*RouteFaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{2}
}
func (m *RouteFaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFaults.Unmarshal(m, b)
}
func (m *RouteFaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFaults.Marshal(b, m, deterministic)
}
func (m *RouteFaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFaults.Merge(m, src)
}
func (m *RouteFaults) XXX_Size() int {
	return xxx_messageInfo_RouteFaults.Size(m)
}
func (m *RouteFaults) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFaults.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFaults proto.InternalMessageInfo

func (m *RouteFaults) GetAbort() *RouteAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *RouteFaults) GetDelay() *RouteDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteAbort)(nil), "fault.options.gloo.solo.io.RouteAbort")
	proto.RegisterType((*RouteDelay)(nil), "fault.options.gloo.solo.io.RouteDelay")
	proto.RegisterType((*RouteFaults)(nil), "fault.options.gloo.solo.io.RouteFaults")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto", fileDescriptor_2594ab72fbfb9a06)
}

var fileDescriptor_2594ab72fbfb9a06 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x75, 0xa3, 0xb1, 0xd8, 0x43, 0x8b, 0x43, 0x30, 0xa6, 0x48, 0x42, 0x0a, 0x09, 0x82, 0xbb,
	0x18, 0x5b, 0x1b, 0x43, 0x50, 0x10, 0x44, 0x3c, 0x3b, 0x9b, 0xb0, 0x97, 0xdb, 0x6c, 0x56, 0xcf,
	0x9b, 0xe3, 0x6e, 0x2e, 0x9c, 0x9f, 0xe0, 0x1f, 0xf8, 0x09, 0xe2, 0x17, 0x88, 0x95, 0x7f, 0x22,
	0xd8, 0xf9, 0x01, 0xf6, 0xb2, 0xbb, 0x17, 0x4c, 0xa3, 0xb1, 0x9b, 0x79, 0xfb, 0xde, 0xcc, 0xbe,
	0xc7, 0xd0, 0x0b, 0xa5, 0x71, 0x5a, 0x84, 0x6c, 0x0c, 0x77, 0x3c, 0x87, 0x18, 0xf6, 0x35, 0x70,
	0x15, 0x03, 0xf0, 0x34, 0x83, 0x1b, 0x39, 0xc6, 0xdc, 0x75, 0x22, 0xd5, 0x7c, 0x76, 0xc0, 0x21,
	0x45, 0x0d, 0x49, 0xce, 0x27, 0xa2, 0x88, 0x51, 0x27, 0x86, 0xa0, 0x21, 0x71, 0x2d, 0x4b, 0x33,
	0x40, 0xf0, 0x9b, 0xae, 0xa9, 0x98, 0xcc, 0xa8, 0x99, 0x19, 0xcc, 0x34, 0x34, 0x5b, 0x0a, 0x40,
	0xc5, 0x92, 0x5b, 0x66, 0x58, 0x4c, 0x78, 0x54, 0x64, 0xc2, 0xf0, 0x9c, 0xb6, 0xb9, 0x3d, 0x13,
	0xb1, 0x8e, 0x04, 0x4a, 0x3e, 0x2f, 0xaa, 0x87, 0x2d, 0x05, 0x0a, 0x6c, 0xc9, 0x4d, 0x55, 0xa1,
	0xbe, 0x2c, 0xd1, 0x81, 0xb2, 0xac, 0xd6, 0x77, 0xcf, 0x29, 0x0d, 0xa0, 0x40, 0x79, 0x1c, 0x42,
	0x86, 0x7e, 0x8b, 0xd2, 0x54, 0x66, 0x63, 0x99, 0xa0, 0x50, 0xb2, 0x41, 0x3a, 0xa4, 0x57, 0x0b,
	0x16, 0x10, 0xbf, 0x4d, 0xbd, 0x29, 0x62, 0x3a, 0xca, 0x51, 0x60, 0x91, 0x37, 0x6a, 0x1d, 0xd2,
	0xdb, 0x08, 0xa8, 0x81, 0xae, 0x2c, 0xd2, 0x2d, 0xab, 0x71, 0x43, 0x19, 0x8b, 0xfb, 0xa5, 0xe3,
	0xce, 0xa8, 0x37, 0xd1, 0xa5, 0x8c, 0x46, 0x91, 0xa1, 0xdb, 0x71, 0x5e, 0x7f, 0x87, 0x39, 0xd7,
	0x6c, 0xee, 0x9a, 0x0d, 0x2b, 0xd7, 0x83, 0xcd, 0xc7, 0xf7, 0x36, 0x79, 0xfd, 0x7c, 0x5b, 0xad,
	0x3f, 0x93, 0xda, 0xde, 0x4a, 0x40, 0xad, 0xda, 0xee, 0xea, 0x3e, 0x10, 0xea, 0xd9, 0xd5, 0x27,
	0x26, 0xcf, 0xdc, 0x3f, 0xa2, 0x75, 0x61, 0x3c, 0xd9, 0xb5, 0x5e, 0x7f, 0x97, 0xfd, 0x9e, 0x33,
	0xfb, 0x49, 0x20, 0x70, 0x22, 0xa3, 0x5e, 0xfc, 0xd3, 0x72, 0xb5, 0xfd, 0x44, 0xe0, 0x44, 0x83,
	0xcb, 0x97, 0xaf, 0x35, 0xf2, 0xf4, 0xd1, 0x22, 0xd7, 0xa7, 0xff, 0x3b, 0x97, 0xf4, 0x56, 0xfd,
	0x7d, 0x32, 0xe1, 0xba, 0x4d, 0xe3, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x0d, 0x10, 0x38,
	0x80, 0x02, 0x00, 0x00,
}

func (this *RouteAbort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteAbort)
	if !ok {
		that2, ok := that.(RouteAbort)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.HttpStatus != that1.HttpStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteDelay) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteDelay)
	if !ok {
		that2, ok := that.(RouteDelay)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.FixedDelay != nil && that1.FixedDelay != nil {
		if *this.FixedDelay != *that1.FixedDelay {
			return false
		}
	} else if this.FixedDelay != nil {
		return false
	} else if that1.FixedDelay != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteFaults) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteFaults)
	if !ok {
		that2, ok := that.(RouteFaults)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Abort.Equal(that1.Abort) {
		return false
	}
	if !this.Delay.Equal(that1.Delay) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
