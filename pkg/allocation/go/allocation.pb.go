// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/allocation/allocation.proto

package allocation

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "google.golang.org/genproto/googleapis/api/annotations"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AllocationRequest_SchedulingStrategy int32

const (
	AllocationRequest_Packed      AllocationRequest_SchedulingStrategy = 0
	AllocationRequest_Distributed AllocationRequest_SchedulingStrategy = 1
)

var AllocationRequest_SchedulingStrategy_name = map[int32]string{
	0: "Packed",
	1: "Distributed",
}
var AllocationRequest_SchedulingStrategy_value = map[string]int32{
	"Packed":      0,
	"Distributed": 1,
}

func (x AllocationRequest_SchedulingStrategy) String() string {
	return proto.EnumName(AllocationRequest_SchedulingStrategy_name, int32(x))
}
func (AllocationRequest_SchedulingStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{0, 0}
}

type AllocationRequest struct {
	// The k8s namespace that is hosting the targeted fleet of gameservers to be allocated
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// If specified, multi-cluster policies are applied. Otherwise, allocation will happen locally.
	MultiClusterSetting *MultiClusterSetting `protobuf:"bytes,2,opt,name=multiClusterSetting,proto3" json:"multiClusterSetting,omitempty"`
	// The required allocation. Defaults to all GameServers.
	RequiredGameServerSelector *LabelSelector `protobuf:"bytes,3,opt,name=requiredGameServerSelector,proto3" json:"requiredGameServerSelector,omitempty"`
	// The ordered list of preferred allocations out of the `required` set.
	// If the first selector is not matched, the selection attempts the second selector, and so on.
	PreferredGameServerSelectors []*LabelSelector `protobuf:"bytes,4,rep,name=preferredGameServerSelectors,proto3" json:"preferredGameServerSelectors,omitempty"`
	// Scheduling strategy. Defaults to "Packed".
	Scheduling AllocationRequest_SchedulingStrategy `protobuf:"varint,5,opt,name=scheduling,proto3,enum=allocation.AllocationRequest_SchedulingStrategy" json:"scheduling,omitempty"`
	// MetaPatch is optional custom metadata that is added to the game server at
	// allocation You can use this to tell the server necessary session data
	MetaPatch            *MetaPatch `protobuf:"bytes,6,opt,name=metaPatch,proto3" json:"metaPatch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AllocationRequest) Reset()         { *m = AllocationRequest{} }
func (m *AllocationRequest) String() string { return proto.CompactTextString(m) }
func (*AllocationRequest) ProtoMessage()    {}
func (*AllocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{0}
}
func (m *AllocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationRequest.Unmarshal(m, b)
}
func (m *AllocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationRequest.Marshal(b, m, deterministic)
}
func (dst *AllocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationRequest.Merge(dst, src)
}
func (m *AllocationRequest) XXX_Size() int {
	return xxx_messageInfo_AllocationRequest.Size(m)
}
func (m *AllocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationRequest proto.InternalMessageInfo

func (m *AllocationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AllocationRequest) GetMultiClusterSetting() *MultiClusterSetting {
	if m != nil {
		return m.MultiClusterSetting
	}
	return nil
}

func (m *AllocationRequest) GetRequiredGameServerSelector() *LabelSelector {
	if m != nil {
		return m.RequiredGameServerSelector
	}
	return nil
}

func (m *AllocationRequest) GetPreferredGameServerSelectors() []*LabelSelector {
	if m != nil {
		return m.PreferredGameServerSelectors
	}
	return nil
}

func (m *AllocationRequest) GetScheduling() AllocationRequest_SchedulingStrategy {
	if m != nil {
		return m.Scheduling
	}
	return AllocationRequest_Packed
}

func (m *AllocationRequest) GetMetaPatch() *MetaPatch {
	if m != nil {
		return m.MetaPatch
	}
	return nil
}

type AllocationResponse struct {
	GameServerName       string                                     `protobuf:"bytes,2,opt,name=gameServerName,proto3" json:"gameServerName,omitempty"`
	Ports                []*AllocationResponse_GameServerStatusPort `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	Address              string                                     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	NodeName             string                                     `protobuf:"bytes,5,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AllocationResponse) Reset()         { *m = AllocationResponse{} }
func (m *AllocationResponse) String() string { return proto.CompactTextString(m) }
func (*AllocationResponse) ProtoMessage()    {}
func (*AllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{1}
}
func (m *AllocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationResponse.Unmarshal(m, b)
}
func (m *AllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationResponse.Marshal(b, m, deterministic)
}
func (dst *AllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationResponse.Merge(dst, src)
}
func (m *AllocationResponse) XXX_Size() int {
	return xxx_messageInfo_AllocationResponse.Size(m)
}
func (m *AllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationResponse proto.InternalMessageInfo

func (m *AllocationResponse) GetGameServerName() string {
	if m != nil {
		return m.GameServerName
	}
	return ""
}

func (m *AllocationResponse) GetPorts() []*AllocationResponse_GameServerStatusPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *AllocationResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AllocationResponse) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

// The gameserver port info that is allocated.
type AllocationResponse_GameServerStatusPort struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationResponse_GameServerStatusPort) Reset() {
	*m = AllocationResponse_GameServerStatusPort{}
}
func (m *AllocationResponse_GameServerStatusPort) String() string { return proto.CompactTextString(m) }
func (*AllocationResponse_GameServerStatusPort) ProtoMessage()    {}
func (*AllocationResponse_GameServerStatusPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{1, 0}
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Unmarshal(m, b)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Marshal(b, m, deterministic)
}
func (dst *AllocationResponse_GameServerStatusPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationResponse_GameServerStatusPort.Merge(dst, src)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Size() int {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Size(m)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationResponse_GameServerStatusPort.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationResponse_GameServerStatusPort proto.InternalMessageInfo

func (m *AllocationResponse_GameServerStatusPort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AllocationResponse_GameServerStatusPort) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Specifies settings for multi-cluster allocation.
type MultiClusterSetting struct {
	// If set to true, multi-cluster allocation is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Selects multi-cluster allocation policies to apply. If not specified, all multi-cluster allocation policies are to be applied.
	PolicySelector       *LabelSelector `protobuf:"bytes,2,opt,name=policySelector,proto3" json:"policySelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MultiClusterSetting) Reset()         { *m = MultiClusterSetting{} }
func (m *MultiClusterSetting) String() string { return proto.CompactTextString(m) }
func (*MultiClusterSetting) ProtoMessage()    {}
func (*MultiClusterSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{2}
}
func (m *MultiClusterSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterSetting.Unmarshal(m, b)
}
func (m *MultiClusterSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterSetting.Marshal(b, m, deterministic)
}
func (dst *MultiClusterSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterSetting.Merge(dst, src)
}
func (m *MultiClusterSetting) XXX_Size() int {
	return xxx_messageInfo_MultiClusterSetting.Size(m)
}
func (m *MultiClusterSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterSetting.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterSetting proto.InternalMessageInfo

func (m *MultiClusterSetting) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *MultiClusterSetting) GetPolicySelector() *LabelSelector {
	if m != nil {
		return m.PolicySelector
	}
	return nil
}

// MetaPatch is the metadata used to patch the GameServer metadata on allocation
type MetaPatch struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetaPatch) Reset()         { *m = MetaPatch{} }
func (m *MetaPatch) String() string { return proto.CompactTextString(m) }
func (*MetaPatch) ProtoMessage()    {}
func (*MetaPatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{3}
}
func (m *MetaPatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaPatch.Unmarshal(m, b)
}
func (m *MetaPatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaPatch.Marshal(b, m, deterministic)
}
func (dst *MetaPatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaPatch.Merge(dst, src)
}
func (m *MetaPatch) XXX_Size() int {
	return xxx_messageInfo_MetaPatch.Size(m)
}
func (m *MetaPatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaPatch.DiscardUnknown(m)
}

var xxx_messageInfo_MetaPatch proto.InternalMessageInfo

func (m *MetaPatch) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetaPatch) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

// LabelSelector used for finding a GameServer with matching labels.
type LabelSelector struct {
	// Labels to match.
	MatchLabels          map[string]string `protobuf:"bytes,1,rep,name=matchLabels,proto3" json:"matchLabels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_d5752a3ca48e8a0b, []int{4}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelector.Unmarshal(m, b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
}
func (dst *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(dst, src)
}
func (m *LabelSelector) XXX_Size() int {
	return xxx_messageInfo_LabelSelector.Size(m)
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationRequest)(nil), "allocation.AllocationRequest")
	proto.RegisterType((*AllocationResponse)(nil), "allocation.AllocationResponse")
	proto.RegisterType((*AllocationResponse_GameServerStatusPort)(nil), "allocation.AllocationResponse.GameServerStatusPort")
	proto.RegisterType((*MultiClusterSetting)(nil), "allocation.MultiClusterSetting")
	proto.RegisterType((*MetaPatch)(nil), "allocation.MetaPatch")
	proto.RegisterMapType((map[string]string)(nil), "allocation.MetaPatch.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "allocation.MetaPatch.LabelsEntry")
	proto.RegisterType((*LabelSelector)(nil), "allocation.LabelSelector")
	proto.RegisterMapType((map[string]string)(nil), "allocation.LabelSelector.MatchLabelsEntry")
	proto.RegisterEnum("allocation.AllocationRequest_SchedulingStrategy", AllocationRequest_SchedulingStrategy_name, AllocationRequest_SchedulingStrategy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AllocationServiceClient is the client API for AllocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllocationServiceClient interface {
	Allocate(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error)
}

type allocationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAllocationServiceClient(cc *grpc.ClientConn) AllocationServiceClient {
	return &allocationServiceClient{cc}
}

func (c *allocationServiceClient) Allocate(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error) {
	out := new(AllocationResponse)
	err := c.cc.Invoke(ctx, "/allocation.AllocationService/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocationServiceServer is the server API for AllocationService service.
type AllocationServiceServer interface {
	Allocate(context.Context, *AllocationRequest) (*AllocationResponse, error)
}

func RegisterAllocationServiceServer(s *grpc.Server, srv AllocationServiceServer) {
	s.RegisterService(&_AllocationService_serviceDesc, srv)
}

func _AllocationService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allocation.AllocationService/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Allocate(ctx, req.(*AllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AllocationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allocation.AllocationService",
	HandlerType: (*AllocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocate",
			Handler:    _AllocationService_Allocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/allocation/allocation.proto",
}

func init() {
	proto.RegisterFile("proto/allocation/allocation.proto", fileDescriptor_allocation_d5752a3ca48e8a0b)
}

var fileDescriptor_allocation_d5752a3ca48e8a0b = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x65, 0x93, 0x26, 0x34, 0x13, 0x51, 0xc2, 0xb4, 0x95, 0x96, 0x55, 0x81, 0x74, 0x0f, 0x55,
	0xd5, 0x43, 0x02, 0xed, 0x85, 0xf6, 0x50, 0xa9, 0x02, 0x04, 0x48, 0x2d, 0x0a, 0x9b, 0x13, 0x07,
	0x0e, 0xce, 0xee, 0x90, 0xae, 0xea, 0xac, 0xb7, 0xb6, 0xb7, 0x52, 0x6e, 0x88, 0x2b, 0x47, 0xce,
	0xfc, 0x2a, 0xfe, 0x02, 0x3f, 0x82, 0x0b, 0x12, 0xb2, 0xf3, 0xe5, 0xa6, 0x69, 0x50, 0x6f, 0xf6,
	0xcc, 0x7b, 0xcf, 0x33, 0xcf, 0x63, 0xc3, 0x76, 0x2e, 0x85, 0x16, 0x6d, 0xc6, 0xb9, 0x88, 0x99,
	0x4e, 0x45, 0xe6, 0x2c, 0x5b, 0x36, 0x87, 0x30, 0x8b, 0x04, 0x5b, 0x7d, 0x21, 0xfa, 0x9c, 0xda,
	0x2c, 0x4f, 0xdb, 0x2c, 0xcb, 0x84, 0xb6, 0x61, 0x35, 0x42, 0x86, 0x7f, 0xca, 0xf0, 0xe8, 0x64,
	0x0a, 0x8e, 0xe8, 0xb2, 0x20, 0xa5, 0x71, 0x0b, 0x6a, 0x19, 0x1b, 0x90, 0xca, 0x59, 0x4c, 0xbe,
	0xd7, 0xf4, 0x76, 0x6b, 0xd1, 0x2c, 0x80, 0x1f, 0x61, 0x7d, 0x50, 0x70, 0x9d, 0xbe, 0xe2, 0x85,
	0xd2, 0x24, 0xbb, 0xa4, 0x75, 0x9a, 0xf5, 0xfd, 0x52, 0xd3, 0xdb, 0xad, 0xef, 0x3f, 0x6b, 0x39,
	0xd5, 0x9c, 0xdd, 0x84, 0x45, 0x8b, 0xb8, 0xf8, 0x09, 0x02, 0x49, 0x97, 0x45, 0x2a, 0x29, 0x79,
	0xcb, 0x06, 0xd4, 0x25, 0x79, 0x65, 0x92, 0x9c, 0x62, 0x2d, 0xa4, 0x5f, 0xb6, 0xca, 0x8f, 0x5d,
	0xe5, 0x53, 0xd6, 0x23, 0x3e, 0x01, 0x44, 0x4b, 0xc8, 0xf8, 0x19, 0xb6, 0x72, 0x49, 0x5f, 0x48,
	0x2e, 0x4c, 0x2b, 0x7f, 0xa5, 0x59, 0x5e, 0x2e, 0xbe, 0x94, 0x8e, 0x1d, 0x00, 0x15, 0x9f, 0x53,
	0x52, 0x70, 0xe3, 0x41, 0xa5, 0xe9, 0xed, 0xae, 0xed, 0x3f, 0x77, 0xc5, 0x6e, 0xb8, 0xdb, 0xea,
	0x4e, 0xf1, 0x5d, 0x2d, 0x99, 0xa6, 0xfe, 0x30, 0x72, 0x34, 0xf0, 0x00, 0x6a, 0x03, 0xd2, 0xac,
	0xc3, 0x74, 0x7c, 0xee, 0x57, 0x6d, 0xeb, 0x9b, 0xd7, 0x4c, 0x9d, 0x24, 0xa3, 0x19, 0x2e, 0x7c,
	0x01, 0x78, 0x53, 0x16, 0x01, 0xaa, 0x1d, 0x16, 0x5f, 0x50, 0xd2, 0xb8, 0x87, 0x0f, 0xa1, 0xfe,
	0x3a, 0x55, 0x5a, 0xa6, 0xbd, 0x42, 0x53, 0xd2, 0xf0, 0xc2, 0xbf, 0x1e, 0xa0, 0x5b, 0x9c, 0xca,
	0x45, 0xa6, 0x08, 0x77, 0x60, 0xad, 0x3f, 0xed, 0xf3, 0x03, 0x1b, 0x90, 0xbd, 0xd8, 0x5a, 0x34,
	0x17, 0xc5, 0xf7, 0x50, 0xc9, 0x85, 0xd4, 0xca, 0x2f, 0x5b, 0x03, 0x0f, 0x6e, 0xeb, 0x79, 0x24,
	0xdb, 0x72, 0xbc, 0xd3, 0x4c, 0x17, 0xaa, 0x23, 0xa4, 0x8e, 0x46, 0x0a, 0xe8, 0xc3, 0x7d, 0x96,
	0x24, 0x92, 0x94, 0xb9, 0x0d, 0x73, 0xd6, 0x64, 0x8b, 0x01, 0xac, 0x66, 0x22, 0x21, 0x5b, 0x46,
	0xc5, 0xa6, 0xa6, 0xfb, 0xe0, 0x18, 0x36, 0x16, 0x89, 0x22, 0xc2, 0x8a, 0x99, 0xd5, 0xf1, 0xdc,
	0xda, 0xb5, 0x89, 0x99, 0xa3, 0x6c, 0x2b, 0x95, 0xc8, 0xae, 0x43, 0x09, 0xeb, 0x0b, 0xe6, 0xd3,
	0x14, 0x43, 0x19, 0xeb, 0x71, 0x4a, 0xac, 0xc2, 0x6a, 0x34, 0xd9, 0xe2, 0x09, 0xac, 0xe5, 0x82,
	0xa7, 0xf1, 0x70, 0x3a, 0x98, 0xa5, 0xff, 0x0d, 0xe6, 0x1c, 0x21, 0xfc, 0x5e, 0x82, 0xda, 0xf4,
	0xfe, 0xf0, 0x10, 0xaa, 0xdc, 0xc0, 0x95, 0xef, 0x59, 0x0f, 0xb7, 0x17, 0x5e, 0xf3, 0x48, 0x52,
	0xbd, 0xc9, 0xb4, 0x1c, 0x46, 0x63, 0x02, 0xbe, 0x83, 0xba, 0xf3, 0x98, 0xfd, 0x92, 0xe5, 0xef,
	0x2c, 0xe6, 0x9f, 0xcc, 0x80, 0x23, 0x11, 0x97, 0x1a, 0x1c, 0x42, 0xdd, 0x39, 0x00, 0x1b, 0x50,
	0xbe, 0xa0, 0xe1, 0xd8, 0x3c, 0xb3, 0xc4, 0x0d, 0xa8, 0x5c, 0x31, 0x5e, 0x4c, 0xe6, 0x60, 0xb4,
	0x39, 0x2a, 0xbd, 0xf4, 0x82, 0x63, 0x68, 0xcc, 0x6b, 0xdf, 0x85, 0x1f, 0xfe, 0xf4, 0xe0, 0xc1,
	0x35, 0xbf, 0xf0, 0x14, 0xea, 0x03, 0x53, 0xf3, 0xa9, 0x6b, 0xcb, 0xde, 0xad, 0xfe, 0xb6, 0xce,
	0x66, 0xe0, 0x71, 0x6b, 0x0e, 0xdd, 0xd4, 0x37, 0x0f, 0xb8, 0x4b, 0x7d, 0xfb, 0x5f, 0x3d, 0xf7,
	0x73, 0x34, 0x83, 0x96, 0xc6, 0x84, 0x17, 0xb0, 0x3a, 0x0e, 0x12, 0x3e, 0x59, 0xfa, 0xd2, 0x83,
	0xa7, 0xcb, 0x1f, 0x45, 0xd8, 0xfc, 0xf6, 0xeb, 0xf7, 0x8f, 0x52, 0x10, 0x6e, 0xb6, 0xcd, 0xe3,
	0x52, 0x76, 0x92, 0x67, 0x8c, 0x23, 0x6f, 0xaf, 0x57, 0xb5, 0xdf, 0xf4, 0xc1, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x15, 0x3d, 0x53, 0xf5, 0x05, 0x00, 0x00,
}