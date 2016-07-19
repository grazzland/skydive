// Code generated by protoc-gen-go.
// source: flow/request.proto
// DO NOT EDIT!

package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BoolFilterOp int32

const (
	BoolFilterOp_OR  BoolFilterOp = 0
	BoolFilterOp_AND BoolFilterOp = 1
	BoolFilterOp_NOT BoolFilterOp = 2
)

var BoolFilterOp_name = map[int32]string{
	0: "OR",
	1: "AND",
	2: "NOT",
}
var BoolFilterOp_value = map[string]int32{
	"OR":  0,
	"AND": 1,
	"NOT": 2,
}

func (x BoolFilterOp) String() string {
	return proto.EnumName(BoolFilterOp_name, int32(x))
}
func (BoolFilterOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type TermStringFilter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *TermStringFilter) Reset()                    { *m = TermStringFilter{} }
func (m *TermStringFilter) String() string            { return proto.CompactTextString(m) }
func (*TermStringFilter) ProtoMessage()               {}
func (*TermStringFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type TermInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *TermInt64Filter) Reset()                    { *m = TermInt64Filter{} }
func (m *TermInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*TermInt64Filter) ProtoMessage()               {}
func (*TermInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type NeStringFilter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *NeStringFilter) Reset()                    { *m = NeStringFilter{} }
func (m *NeStringFilter) String() string            { return proto.CompactTextString(m) }
func (*NeStringFilter) ProtoMessage()               {}
func (*NeStringFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type NeInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *NeInt64Filter) Reset()                    { *m = NeInt64Filter{} }
func (m *NeInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*NeInt64Filter) ProtoMessage()               {}
func (*NeInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type GtInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *GtInt64Filter) Reset()                    { *m = GtInt64Filter{} }
func (m *GtInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*GtInt64Filter) ProtoMessage()               {}
func (*GtInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type LtInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *LtInt64Filter) Reset()                    { *m = LtInt64Filter{} }
func (m *LtInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*LtInt64Filter) ProtoMessage()               {}
func (*LtInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type GteInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *GteInt64Filter) Reset()                    { *m = GteInt64Filter{} }
func (m *GteInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*GteInt64Filter) ProtoMessage()               {}
func (*GteInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type LteInt64Filter struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *LteInt64Filter) Reset()                    { *m = LteInt64Filter{} }
func (m *LteInt64Filter) String() string            { return proto.CompactTextString(m) }
func (*LteInt64Filter) ProtoMessage()               {}
func (*LteInt64Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type Filter struct {
	TermStringFilter *TermStringFilter `protobuf:"bytes,1,opt,name=TermStringFilter" json:"TermStringFilter,omitempty"`
	TermInt64Filter  *TermInt64Filter  `protobuf:"bytes,2,opt,name=TermInt64Filter" json:"TermInt64Filter,omitempty"`
	GtInt64Filter    *GtInt64Filter    `protobuf:"bytes,3,opt,name=GtInt64Filter" json:"GtInt64Filter,omitempty"`
	LtInt64Filter    *LtInt64Filter    `protobuf:"bytes,4,opt,name=LtInt64Filter" json:"LtInt64Filter,omitempty"`
	GteInt64Filter   *GteInt64Filter   `protobuf:"bytes,5,opt,name=GteInt64Filter" json:"GteInt64Filter,omitempty"`
	LteInt64Filter   *LteInt64Filter   `protobuf:"bytes,6,opt,name=LteInt64Filter" json:"LteInt64Filter,omitempty"`
	BoolFilter       *BoolFilter       `protobuf:"bytes,7,opt,name=BoolFilter" json:"BoolFilter,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *Filter) GetTermStringFilter() *TermStringFilter {
	if m != nil {
		return m.TermStringFilter
	}
	return nil
}

func (m *Filter) GetTermInt64Filter() *TermInt64Filter {
	if m != nil {
		return m.TermInt64Filter
	}
	return nil
}

func (m *Filter) GetGtInt64Filter() *GtInt64Filter {
	if m != nil {
		return m.GtInt64Filter
	}
	return nil
}

func (m *Filter) GetLtInt64Filter() *LtInt64Filter {
	if m != nil {
		return m.LtInt64Filter
	}
	return nil
}

func (m *Filter) GetGteInt64Filter() *GteInt64Filter {
	if m != nil {
		return m.GteInt64Filter
	}
	return nil
}

func (m *Filter) GetLteInt64Filter() *LteInt64Filter {
	if m != nil {
		return m.LteInt64Filter
	}
	return nil
}

func (m *Filter) GetBoolFilter() *BoolFilter {
	if m != nil {
		return m.BoolFilter
	}
	return nil
}

type BoolFilter struct {
	Op      BoolFilterOp `protobuf:"varint,1,opt,name=Op,enum=flow.BoolFilterOp" json:"Op,omitempty"`
	Filters []*Filter    `protobuf:"bytes,2,rep,name=Filters" json:"Filters,omitempty"`
}

func (m *BoolFilter) Reset()                    { *m = BoolFilter{} }
func (m *BoolFilter) String() string            { return proto.CompactTextString(m) }
func (*BoolFilter) ProtoMessage()               {}
func (*BoolFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *BoolFilter) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type Range struct {
	From int64 `protobuf:"varint,1,opt,name=From" json:"From,omitempty"`
	To   int64 `protobuf:"varint,2,opt,name=To" json:"To,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

type FlowSearchQuery struct {
	Filter *Filter `protobuf:"bytes,1,opt,name=Filter" json:"Filter,omitempty"`
	Range  *Range  `protobuf:"bytes,2,opt,name=Range" json:"Range,omitempty"`
}

func (m *FlowSearchQuery) Reset()                    { *m = FlowSearchQuery{} }
func (m *FlowSearchQuery) String() string            { return proto.CompactTextString(m) }
func (*FlowSearchQuery) ProtoMessage()               {}
func (*FlowSearchQuery) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *FlowSearchQuery) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *FlowSearchQuery) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type FlowSearchReply struct {
	FlowSet *FlowSet `protobuf:"bytes,1,opt,name=FlowSet" json:"FlowSet,omitempty"`
}

func (m *FlowSearchReply) Reset()                    { *m = FlowSearchReply{} }
func (m *FlowSearchReply) String() string            { return proto.CompactTextString(m) }
func (*FlowSearchReply) ProtoMessage()               {}
func (*FlowSearchReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *FlowSearchReply) GetFlowSet() *FlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

func init() {
	proto.RegisterType((*TermStringFilter)(nil), "flow.TermStringFilter")
	proto.RegisterType((*TermInt64Filter)(nil), "flow.TermInt64Filter")
	proto.RegisterType((*NeStringFilter)(nil), "flow.NeStringFilter")
	proto.RegisterType((*NeInt64Filter)(nil), "flow.NeInt64Filter")
	proto.RegisterType((*GtInt64Filter)(nil), "flow.GtInt64Filter")
	proto.RegisterType((*LtInt64Filter)(nil), "flow.LtInt64Filter")
	proto.RegisterType((*GteInt64Filter)(nil), "flow.GteInt64Filter")
	proto.RegisterType((*LteInt64Filter)(nil), "flow.LteInt64Filter")
	proto.RegisterType((*Filter)(nil), "flow.Filter")
	proto.RegisterType((*BoolFilter)(nil), "flow.BoolFilter")
	proto.RegisterType((*Range)(nil), "flow.Range")
	proto.RegisterType((*FlowSearchQuery)(nil), "flow.FlowSearchQuery")
	proto.RegisterType((*FlowSearchReply)(nil), "flow.FlowSearchReply")
	proto.RegisterEnum("flow.BoolFilterOp", BoolFilterOp_name, BoolFilterOp_value)
}

var fileDescriptor2 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xbd, 0x6c, 0x7e, 0xe0, 0x5c, 0x2f, 0x17, 0xd6, 0x2a, 0xc1, 0x27, 0x0d, 0xa2, 0x45,
	0xe1, 0x94, 0x53, 0xd4, 0x4a, 0x41, 0x2c, 0x72, 0x45, 0x0c, 0x39, 0xdc, 0x1e, 0x45, 0x7c, 0xbb,
	0xca, 0x5e, 0x2c, 0xa6, 0xb7, 0x71, 0x6f, 0x4f, 0xc9, 0x9f, 0xe5, 0x7f, 0xe8, 0x66, 0x77, 0xcf,
	0x26, 0x9b, 0x87, 0x96, 0x3c, 0x65, 0x32, 0xb3, 0x9f, 0xf9, 0xb2, 0x33, 0x5f, 0x16, 0xf0, 0xaa,
	0x60, 0x7f, 0x9e, 0x73, 0xfa, 0x6b, 0x4b, 0x37, 0x62, 0x52, 0x72, 0x26, 0x18, 0x76, 0xeb, 0xdc,
	0xfd, 0x50, 0x55, 0x36, 0xd4, 0x64, 0x93, 0x77, 0x10, 0x2d, 0x28, 0xbf, 0x3c, 0x15, 0xfc, 0x62,
	0x9d, 0xcf, 0x2e, 0x0a, 0x41, 0x39, 0x8e, 0x00, 0x7d, 0xa6, 0x55, 0x3c, 0x78, 0x30, 0x38, 0xb8,
	0x4d, 0xd0, 0x4f, 0x5a, 0xe1, 0x7d, 0xf0, 0x7e, 0x2f, 0x8b, 0x2d, 0x8d, 0x1d, 0x95, 0xd3, 0x3f,
	0xc9, 0x21, 0x8c, 0x6b, 0xf6, 0xd3, 0x5a, 0xbc, 0x7e, 0x75, 0x33, 0x14, 0xed, 0xd0, 0xb7, 0x10,
	0x66, 0xf4, 0x7a, 0xd1, 0xb3, 0xae, 0xe8, 0x1b, 0x18, 0x65, 0xf4, 0x5a, 0xc9, 0xb3, 0xae, 0xa4,
	0x04, 0x4f, 0x44, 0x4f, 0x30, 0xed, 0x05, 0xca, 0x4b, 0x9e, 0x08, 0xda, 0x93, 0x4c, 0xfb, 0x91,
	0x7f, 0x11, 0xf8, 0x06, 0x39, 0xee, 0xae, 0x56, 0xf1, 0xc3, 0xe9, 0xbd, 0x49, 0xed, 0x82, 0x89,
	0x5d, 0x25, 0x91, 0xb0, 0xad, 0xf0, 0xbe, 0xb3, 0x62, 0x25, 0x37, 0x9c, 0xde, 0xbd, 0x6a, 0xd1,
	0x28, 0x92, 0xb1, 0xb0, 0x0c, 0x71, 0x68, 0x4d, 0x3d, 0x46, 0x0a, 0xbf, 0xa3, 0xf1, 0x56, 0x89,
	0x8c, 0x72, 0x61, 0xa1, 0xad, 0xb9, 0xc7, 0x6e, 0x13, 0x4d, 0xdb, 0x68, 0xd1, 0x42, 0x8f, 0xec,
	0xc9, 0xc7, 0x9e, 0x62, 0xf7, 0x77, 0xb2, 0xcd, 0x1a, 0x09, 0xf3, 0xf6, 0xac, 0x8f, 0xec, 0xe9,
	0xc7, 0x7e, 0x93, 0x4e, 0x2d, 0xba, 0x68, 0xd3, 0x2f, 0x00, 0x8e, 0x19, 0x2b, 0x0c, 0x19, 0x28,
	0x32, 0xd2, 0xe4, 0x55, 0x9e, 0xc0, 0xf9, 0xff, 0x38, 0xf9, 0xda, 0x24, 0x70, 0x02, 0xce, 0xbc,
	0x54, 0x8b, 0x0a, 0xa7, 0xd8, 0xe6, 0xe6, 0x25, 0x71, 0x58, 0x89, 0x1f, 0x43, 0xa0, 0xff, 0x37,
	0x72, 0x1d, 0x48, 0x0a, 0xec, 0xe9, 0x83, 0xa6, 0x79, 0xb0, 0xd2, 0xc5, 0xe4, 0x19, 0x78, 0x64,
	0xb9, 0xce, 0x29, 0xc6, 0xe0, 0xce, 0x38, 0xbb, 0x54, 0x6d, 0x11, 0x71, 0x57, 0x32, 0xc6, 0x21,
	0x38, 0x0b, 0x66, 0xdc, 0xe3, 0xc8, 0xa7, 0xe0, 0x1b, 0x8c, 0x67, 0xb2, 0xc9, 0x29, 0x5d, 0xf2,
	0xef, 0x3f, 0xbe, 0x6c, 0x29, 0xaf, 0xf0, 0xa3, 0x9d, 0x99, 0x8c, 0x71, 0xda, 0x32, 0xbe, 0x96,
	0xc1, 0x0f, 0x8d, 0x8a, 0xb1, 0xc6, 0x50, 0x1f, 0x52, 0x29, 0xe2, 0xf1, 0xfa, 0x23, 0x9f, 0x99,
	0x46, 0x6f, 0x42, 0xcb, 0xa2, 0xc2, 0x4f, 0xe4, 0x1d, 0x54, 0x4a, 0x98, 0xe6, 0x23, 0xd3, 0x5c,
	0x27, 0xe5, 0x25, 0x74, 0xf0, 0xf4, 0x00, 0xf6, 0x9a, 0x03, 0xc0, 0xbe, 0x1c, 0x10, 0x89, 0x6e,
	0xe1, 0x00, 0xd0, 0x87, 0xec, 0x63, 0x34, 0xa8, 0x83, 0x6c, 0xbe, 0x88, 0x9c, 0x73, 0x5f, 0xbd,
	0x69, 0x2f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xfd, 0xde, 0x5d, 0xff, 0x04, 0x00, 0x00,
}