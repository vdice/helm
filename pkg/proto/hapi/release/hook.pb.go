// Code generated by protoc-gen-go.
// source: hapi/release/hook.proto
// DO NOT EDIT!

/*
Package release is a generated protocol buffer package.

It is generated from these files:
	hapi/release/hook.proto
	hapi/release/info.proto
	hapi/release/release.proto
	hapi/release/status.proto
	hapi/release/test_run.proto
	hapi/release/test_suite.proto

It has these top-level messages:
	Hook
	Info
	Release
	Status
	TestRun
	TestSuite
*/
package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hook_Event int32

const (
	Hook_UNKNOWN       Hook_Event = 0
	Hook_PRE_INSTALL   Hook_Event = 1
	Hook_POST_INSTALL  Hook_Event = 2
	Hook_PRE_DELETE    Hook_Event = 3
	Hook_POST_DELETE   Hook_Event = 4
	Hook_PRE_UPGRADE   Hook_Event = 5
	Hook_POST_UPGRADE  Hook_Event = 6
	Hook_PRE_ROLLBACK  Hook_Event = 7
	Hook_POST_ROLLBACK Hook_Event = 8
	Hook_RELEASE_TEST  Hook_Event = 9
)

var Hook_Event_name = map[int32]string{
	0: "UNKNOWN",
	1: "PRE_INSTALL",
	2: "POST_INSTALL",
	3: "PRE_DELETE",
	4: "POST_DELETE",
	5: "PRE_UPGRADE",
	6: "POST_UPGRADE",
	7: "PRE_ROLLBACK",
	8: "POST_ROLLBACK",
	9: "RELEASE_TEST",
}
var Hook_Event_value = map[string]int32{
	"UNKNOWN":       0,
	"PRE_INSTALL":   1,
	"POST_INSTALL":  2,
	"PRE_DELETE":    3,
	"POST_DELETE":   4,
	"PRE_UPGRADE":   5,
	"POST_UPGRADE":  6,
	"PRE_ROLLBACK":  7,
	"POST_ROLLBACK": 8,
	"RELEASE_TEST":  9,
}

func (x Hook_Event) String() string {
	return proto.EnumName(Hook_Event_name, int32(x))
}
func (Hook_Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Hook defines a hook object.
type Hook struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Kind is the Kubernetes kind.
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	// Path is the chart-relative path to the template.
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// Manifest is the manifest contents.
	Manifest string `protobuf:"bytes,4,opt,name=manifest" json:"manifest,omitempty"`
	// Events are the events that this hook fires on.
	Events []Hook_Event `protobuf:"varint,5,rep,packed,name=events,enum=hapi.release.Hook_Event" json:"events,omitempty"`
	// LastRun indicates the date/time this was last run.
	LastRun *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_run,json=lastRun" json:"last_run,omitempty"`
}

func (m *Hook) Reset()                    { *m = Hook{} }
func (m *Hook) String() string            { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()               {}
func (*Hook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hook) GetLastRun() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastRun
	}
	return nil
}

func init() {
	proto.RegisterType((*Hook)(nil), "hapi.release.Hook")
	proto.RegisterEnum("hapi.release.Hook_Event", Hook_Event_name, Hook_Event_value)
}

func init() { proto.RegisterFile("hapi/release/hook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x6e, 0xb2, 0x40,
	0x14, 0x86, 0x3f, 0x14, 0x41, 0x8f, 0x7e, 0x96, 0xce, 0xa6, 0x13, 0x37, 0x35, 0xae, 0x5c, 0x0d,
	0x8d, 0x4d, 0x2f, 0x00, 0xeb, 0xa4, 0x6d, 0x24, 0x68, 0x06, 0x4c, 0x93, 0x6e, 0xcc, 0x98, 0x8e,
	0x4a, 0x94, 0x9f, 0x08, 0xf6, 0xc2, 0x7a, 0x7d, 0x5d, 0x94, 0x19, 0x7e, 0xd2, 0xdd, 0xe1, 0x39,
	0xcf, 0x79, 0xe1, 0x05, 0xee, 0x8e, 0x3c, 0x0d, 0xed, 0x8b, 0x38, 0x0b, 0x9e, 0x09, 0xfb, 0x98,
	0x24, 0x27, 0x92, 0x5e, 0x92, 0x3c, 0x41, 0x03, 0xb9, 0x20, 0xd5, 0x62, 0x74, 0x7f, 0x48, 0x92,
	0xc3, 0x59, 0xd8, 0x6a, 0xb7, 0xbb, 0xee, 0xed, 0x3c, 0x8c, 0x44, 0x96, 0xf3, 0x28, 0x2d, 0xf5,
	0xc9, 0x4f, 0x0b, 0xf4, 0xd7, 0xe2, 0x1a, 0x21, 0xd0, 0x63, 0x1e, 0x09, 0xac, 0x8d, 0xb5, 0x69,
	0x8f, 0xa9, 0x59, 0xb2, 0x53, 0x18, 0x7f, 0xe2, 0x56, 0xc9, 0xe4, 0x2c, 0x59, 0xca, 0xf3, 0x23,
	0x6e, 0x97, 0x4c, 0xce, 0x68, 0x04, 0xdd, 0x88, 0xc7, 0xe1, 0xbe, 0x48, 0xc6, 0xba, 0xe2, 0xcd,
	0x33, 0x7a, 0x00, 0x43, 0x7c, 0x89, 0x38, 0xcf, 0x70, 0x67, 0xdc, 0x9e, 0x0e, 0x67, 0x98, 0xfc,
	0xfd, 0x40, 0x22, 0xdf, 0x4d, 0xa8, 0x14, 0x58, 0xe5, 0xa1, 0x27, 0xe8, 0x9e, 0x79, 0x96, 0x6f,
	0x2f, 0xd7, 0x18, 0x1b, 0x45, 0x5a, 0x7f, 0x36, 0x22, 0x65, 0x0d, 0x52, 0xd7, 0x20, 0x41, 0x5d,
	0x83, 0x99, 0xd2, 0x65, 0xd7, 0x78, 0xf2, 0xad, 0x41, 0x47, 0x05, 0xa1, 0x3e, 0x98, 0x1b, 0x6f,
	0xe9, 0xad, 0xde, 0x3d, 0xeb, 0x1f, 0xba, 0x81, 0xfe, 0x9a, 0xd1, 0xed, 0x9b, 0xe7, 0x07, 0x8e,
	0xeb, 0x5a, 0x1a, 0xb2, 0x60, 0xb0, 0x5e, 0xf9, 0x41, 0x43, 0x5a, 0x68, 0x08, 0x20, 0x95, 0x05,
	0x75, 0x69, 0x40, 0xad, 0xb6, 0x3a, 0x91, 0x46, 0x05, 0xf4, 0x3a, 0x63, 0xb3, 0x7e, 0x61, 0xce,
	0x82, 0x5a, 0x9d, 0x26, 0xa3, 0x26, 0x86, 0x22, 0x85, 0xc2, 0x56, 0xae, 0x3b, 0x77, 0x9e, 0x97,
	0x96, 0x89, 0x6e, 0xe1, 0xbf, 0x72, 0x1a, 0xd4, 0x95, 0x12, 0x2b, 0x32, 0x1d, 0x9f, 0x6e, 0x03,
	0xea, 0x07, 0x56, 0x6f, 0xde, 0xfb, 0x30, 0xab, 0x3f, 0xb1, 0x33, 0x54, 0xb9, 0xc7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x35, 0xb7, 0x2a, 0x22, 0xda, 0x01, 0x00, 0x00,
}
