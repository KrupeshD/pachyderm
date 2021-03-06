// Code generated by protoc-gen-go.
// source: server/pkg/metrics/metrics.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	server/pkg/metrics/metrics.proto

It has these top-level messages:
	Metrics
*/
package metrics

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

type Metrics struct {
	ClusterID        string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	PodID            string `protobuf:"bytes,2,opt,name=pod_id,json=podId" json:"pod_id,omitempty"`
	Nodes            int64  `protobuf:"varint,3,opt,name=nodes" json:"nodes,omitempty"`
	Version          string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Repos            int64  `protobuf:"varint,5,opt,name=repos" json:"repos,omitempty"`
	Commits          int64  `protobuf:"varint,6,opt,name=commits" json:"commits,omitempty"`
	Files            int64  `protobuf:"varint,7,opt,name=files" json:"files,omitempty"`
	Bytes            int64  `protobuf:"varint,8,opt,name=bytes" json:"bytes,omitempty"`
	Jobs             int64  `protobuf:"varint,9,opt,name=jobs" json:"jobs,omitempty"`
	Pipelines        int64  `protobuf:"varint,10,opt,name=pipelines" json:"pipelines,omitempty"`
	ArchivedCommits  int64  `protobuf:"varint,11,opt,name=archived_commits,json=archivedCommits" json:"archived_commits,omitempty"`
	CancelledCommits int64  `protobuf:"varint,12,opt,name=cancelled_commits,json=cancelledCommits" json:"cancelled_commits,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Metrics)(nil), "Metrics")
}

func init() { proto.RegisterFile("server/pkg/metrics/metrics.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xdb, 0x24, 0x66, 0x14, 0xac, 0x8b, 0xc2, 0x1c, 0x14, 0x82, 0xa7, 0x8a, 0x60,
	0x0f, 0x3e, 0x82, 0xa7, 0x1e, 0xbc, 0xf4, 0x05, 0x4a, 0xb3, 0x3b, 0xea, 0x6a, 0x92, 0x5d, 0x76,
	0x62, 0xc0, 0x07, 0xf5, 0x7d, 0x24, 0xb3, 0x59, 0xf5, 0x94, 0x7c, 0xdf, 0xff, 0x1d, 0x86, 0x85,
	0x9a, 0x29, 0x8c, 0x14, 0xb6, 0xfe, 0xe3, 0x75, 0xdb, 0xd1, 0x10, 0xac, 0xe6, 0xf4, 0x7d, 0xf0,
	0xc1, 0x0d, 0xee, 0xf6, 0x7b, 0x01, 0xe5, 0x73, 0x34, 0xea, 0x06, 0x40, 0xb7, 0x9f, 0x3c, 0x50,
	0x38, 0x58, 0x83, 0x59, 0x9d, 0x6d, 0xaa, 0x7d, 0x35, 0x9b, 0x9d, 0x51, 0x57, 0x50, 0x78, 0x67,
	0xa6, 0x69, 0x21, 0x53, 0xee, 0x9d, 0xd9, 0x19, 0x75, 0x09, 0x79, 0xef, 0x0c, 0x31, 0x2e, 0xeb,
	0x6c, 0xb3, 0xdc, 0x47, 0x50, 0x08, 0xe5, 0x48, 0x81, 0xad, 0xeb, 0x71, 0x25, 0x75, 0xc2, 0xa9,
	0x0f, 0xe4, 0x1d, 0x63, 0x1e, 0x7b, 0x81, 0xa9, 0xd7, 0xae, 0xeb, 0xec, 0xc0, 0x58, 0x88, 0x4f,
	0x38, 0xf5, 0x2f, 0xb6, 0x25, 0xc6, 0x32, 0xf6, 0x02, 0x93, 0x6d, 0xbe, 0x06, 0x62, 0x3c, 0x89,
	0x56, 0x40, 0x29, 0x58, 0xbd, 0xbb, 0x86, 0xb1, 0x12, 0x29, 0xff, 0xea, 0x1a, 0x2a, 0x6f, 0x3d,
	0xb5, 0xb6, 0x27, 0x46, 0x90, 0xe1, 0x4f, 0xa8, 0x3b, 0x58, 0x1f, 0x83, 0x7e, 0xb3, 0x23, 0x99,
	0x43, 0x3a, 0xe0, 0x54, 0xa2, 0xf3, 0xe4, 0x9f, 0xe6, 0x43, 0xee, 0xe1, 0x42, 0x1f, 0x7b, 0x4d,
	0x6d, 0xfb, 0xaf, 0x3d, 0x93, 0x76, 0xfd, 0x3b, 0xcc, 0x71, 0x53, 0xc8, 0xf3, 0x3e, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x59, 0x42, 0xa4, 0x5d, 0x82, 0x01, 0x00, 0x00,
}
