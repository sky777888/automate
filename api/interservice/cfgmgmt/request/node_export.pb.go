// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/request/node_export.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NodeExport struct {
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Sorting              *Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty" toml:"sorting,omitempty" mapstructure:"sorting,omitempty"`
	OutputType           string   `protobuf:"bytes,3,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty" toml:"output_type,omitempty" mapstructure:"output_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodeExport) Reset()         { *m = NodeExport{} }
func (m *NodeExport) String() string { return proto.CompactTextString(m) }
func (*NodeExport) ProtoMessage()    {}
func (*NodeExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_14256eaca27cb923, []int{0}
}

func (m *NodeExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExport.Unmarshal(m, b)
}
func (m *NodeExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExport.Marshal(b, m, deterministic)
}
func (m *NodeExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExport.Merge(m, src)
}
func (m *NodeExport) XXX_Size() int {
	return xxx_messageInfo_NodeExport.Size(m)
}
func (m *NodeExport) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExport.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExport proto.InternalMessageInfo

func (m *NodeExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodeExport) GetSorting() *Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *NodeExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

type ReportExport struct {
	Filter               []string             `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	OutputType           string               `protobuf:"bytes,2,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty" toml:"output_type,omitempty" mapstructure:"output_type,omitempty"`
	NodeId               string               `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ReportExport) Reset()         { *m = ReportExport{} }
func (m *ReportExport) String() string { return proto.CompactTextString(m) }
func (*ReportExport) ProtoMessage()    {}
func (*ReportExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_14256eaca27cb923, []int{1}
}

func (m *ReportExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportExport.Unmarshal(m, b)
}
func (m *ReportExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportExport.Marshal(b, m, deterministic)
}
func (m *ReportExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportExport.Merge(m, src)
}
func (m *ReportExport) XXX_Size() int {
	return xxx_messageInfo_ReportExport.Size(m)
}
func (m *ReportExport) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportExport.DiscardUnknown(m)
}

var xxx_messageInfo_ReportExport proto.InternalMessageInfo

func (m *ReportExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReportExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

func (m *ReportExport) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReportExport) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ReportExport) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeExport)(nil), "chef.automate.domain.cfgmgmt.request.NodeExport")
	proto.RegisterType((*ReportExport)(nil), "chef.automate.domain.cfgmgmt.request.ReportExport")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/request/node_export.proto", fileDescriptor_14256eaca27cb923)
}

var fileDescriptor_14256eaca27cb923 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x96, 0xb6, 0xaa, 0xcb, 0xe4, 0x01, 0xa2, 0x2e, 0xad, 0x2a, 0x86, 0x0e, 0x60,
	0x43, 0x99, 0x10, 0x1b, 0x12, 0x42, 0x2c, 0x0c, 0xa1, 0x13, 0x4b, 0xe5, 0x24, 0x97, 0xd4, 0x52,
	0x9d, 0x33, 0xce, 0x05, 0xd1, 0x3f, 0xc1, 0x6f, 0xe2, 0xa7, 0xa1, 0xc4, 0xc9, 0xd2, 0x25, 0x8c,
	0x67, 0xbd, 0xf7, 0xee, 0xf9, 0x3b, 0xb6, 0x51, 0x56, 0x4b, 0x5d, 0x10, 0xb8, 0x12, 0xdc, 0x97,
	0x4e, 0x40, 0x26, 0x59, 0x6e, 0x72, 0x43, 0xd2, 0xc1, 0x67, 0x05, 0x25, 0xc9, 0x02, 0x53, 0xd8,
	0xc1, 0xb7, 0x45, 0x47, 0xc2, 0x3a, 0x24, 0xe4, 0x57, 0xc9, 0x1e, 0x32, 0xa1, 0x2a, 0x42, 0xa3,
	0x08, 0x44, 0x8a, 0x46, 0xe9, 0x42, 0xb4, 0x3e, 0xd1, 0xfa, 0xe6, 0x8b, 0x1c, 0x31, 0x3f, 0x80,
	0x6c, 0x3c, 0x71, 0x95, 0x49, 0xd2, 0x06, 0x4a, 0x52, 0xc6, 0xfa, 0x98, 0xf9, 0x5d, 0xef, 0x6a,
	0xab, 0x9c, 0x32, 0x50, 0x0b, 0xbc, 0x65, 0xf5, 0x13, 0x30, 0xf6, 0x86, 0x29, 0x3c, 0x37, 0x75,
	0xf8, 0x05, 0x1b, 0x67, 0xfa, 0x40, 0xe0, 0xc2, 0x60, 0x39, 0x5c, 0x4f, 0xa3, 0x76, 0xe2, 0x2f,
	0x6c, 0x52, 0xa2, 0x23, 0x5d, 0xe4, 0xe1, 0x60, 0x19, 0xac, 0x67, 0x9b, 0x1b, 0xf1, 0x9f, 0xca,
	0xe2, 0xdd, 0x9b, 0xa2, 0xce, 0xcd, 0x17, 0x6c, 0x86, 0x15, 0xd9, 0x8a, 0x76, 0x74, 0xb4, 0x10,
	0x0e, 0x97, 0xc1, 0x7a, 0x1a, 0x31, 0xff, 0xb4, 0x3d, 0x5a, 0x58, 0xfd, 0x06, 0xec, 0x3c, 0x82,
	0xba, 0x4c, 0x4f, 0xa5, 0x93, 0xa4, 0xc1, 0x69, 0x12, 0xbf, 0x64, 0x93, 0x86, 0xb4, 0x4e, 0xdb,
	0x35, 0xe3, 0x7a, 0x7c, 0x4d, 0xf9, 0x2d, 0x1b, 0x95, 0xa4, 0x1c, 0x85, 0x67, 0xcd, 0x57, 0xe6,
	0xc2, 0x73, 0x15, 0x1d, 0x57, 0xb1, 0xed, 0xb8, 0x46, 0x5e, 0xc8, 0xaf, 0xd9, 0x10, 0x8a, 0x34,
	0x1c, 0xf5, 0xea, 0x6b, 0xd9, 0xd3, 0xe3, 0xc7, 0x43, 0xae, 0x69, 0x5f, 0xc5, 0x22, 0x41, 0x23,
	0x6b, 0x4e, 0xb2, 0xe3, 0x24, 0xfb, 0x2e, 0x14, 0x8f, 0x9b, 0xd4, 0xfb, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x25, 0x52, 0xfb, 0xb0, 0x47, 0x02, 0x00, 0x00,
}
