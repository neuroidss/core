// Code generated by protoc-gen-go. DO NOT EDIT.
// source: miner.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NATType int32

const (
	NATType_NONE                   NATType = 0
	NATType_BLOCKED                NATType = 1
	NATType_FULL                   NATType = 2
	NATType_SYMMETRIC              NATType = 3
	NATType_RESTRICTED             NATType = 4
	NATType_PORT_RESTRICTED        NATType = 5
	NATType_SYMMETRIC_UDP_FIREWALL NATType = 6
	NATType_UNKNOWN                NATType = 7
)

var NATType_name = map[int32]string{
	0: "NONE",
	1: "BLOCKED",
	2: "FULL",
	3: "SYMMETRIC",
	4: "RESTRICTED",
	5: "PORT_RESTRICTED",
	6: "SYMMETRIC_UDP_FIREWALL",
	7: "UNKNOWN",
}
var NATType_value = map[string]int32{
	"NONE":                   0,
	"BLOCKED":                1,
	"FULL":                   2,
	"SYMMETRIC":              3,
	"RESTRICTED":             4,
	"PORT_RESTRICTED":        5,
	"SYMMETRIC_UDP_FIREWALL": 6,
	"UNKNOWN":                7,
}

func (x NATType) String() string {
	return proto.EnumName(NATType_name, int32(x))
}
func (NATType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MinerInfoRequest struct {
}

func (m *MinerInfoRequest) Reset()                    { *m = MinerInfoRequest{} }
func (m *MinerInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerInfoRequest) ProtoMessage()               {}
func (*MinerInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MinerHandshakeRequest struct {
	Hub string `protobuf:"bytes,1,opt,name=hub" json:"hub,omitempty"`
}

func (m *MinerHandshakeRequest) Reset()                    { *m = MinerHandshakeRequest{} }
func (m *MinerHandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeRequest) ProtoMessage()               {}
func (*MinerHandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MinerHandshakeRequest) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

type MinerHandshakeReply struct {
	Miner        string        `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,2,opt,name=capabilities" json:"capabilities,omitempty"`
	NatType      NATType       `protobuf:"varint,3,opt,name=natType,enum=sonm.NATType" json:"natType,omitempty"`
}

func (m *MinerHandshakeReply) Reset()                    { *m = MinerHandshakeReply{} }
func (m *MinerHandshakeReply) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeReply) ProtoMessage()               {}
func (*MinerHandshakeReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *MinerHandshakeReply) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerHandshakeReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *MinerHandshakeReply) GetNatType() NATType {
	if m != nil {
		return m.NatType
	}
	return NATType_NONE
}

type MinerStartRequest struct {
	Id            string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Registry      string                    `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string                    `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string                    `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	RestartPolicy *ContainerRestartPolicy   `protobuf:"bytes,5,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	Usage         *TaskResourceRequirements `protobuf:"bytes,6,opt,name=usage" json:"usage,omitempty"`
	PublicKeyData string                    `protobuf:"bytes,7,opt,name=PublicKeyData,json=publicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool                      `protobuf:"varint,8,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string         `protobuf:"bytes,9,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartRequest) Reset()                    { *m = MinerStartRequest{} }
func (m *MinerStartRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStartRequest) ProtoMessage()               {}
func (*MinerStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *MinerStartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MinerStartRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *MinerStartRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *MinerStartRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *MinerStartRequest) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *MinerStartRequest) GetUsage() *TaskResourceRequirements {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *MinerStartRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *MinerStartRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *MinerStartRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type MinerStartReply struct {
	Container string                          `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Ports     map[string]*MinerStartReplyPort `protobuf:"bytes,2,rep,name=Ports,json=ports" json:"Ports,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartReply) Reset()                    { *m = MinerStartReply{} }
func (m *MinerStartReply) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReply) ProtoMessage()               {}
func (*MinerStartReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *MinerStartReply) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *MinerStartReply) GetPorts() map[string]*MinerStartReplyPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type MinerStartReplyPort struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *MinerStartReplyPort) Reset()                    { *m = MinerStartReplyPort{} }
func (m *MinerStartReplyPort) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReplyPort) ProtoMessage()               {}
func (*MinerStartReplyPort) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4, 0} }

func (m *MinerStartReplyPort) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MinerStartReplyPort) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type MinerStatusMapRequest struct {
}

func (m *MinerStatusMapRequest) Reset()                    { *m = MinerStatusMapRequest{} }
func (m *MinerStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStatusMapRequest) ProtoMessage()               {}
func (*MinerStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func init() {
	proto.RegisterType((*MinerInfoRequest)(nil), "sonm.MinerInfoRequest")
	proto.RegisterType((*MinerHandshakeRequest)(nil), "sonm.MinerHandshakeRequest")
	proto.RegisterType((*MinerHandshakeReply)(nil), "sonm.MinerHandshakeReply")
	proto.RegisterType((*MinerStartRequest)(nil), "sonm.MinerStartRequest")
	proto.RegisterType((*MinerStartReply)(nil), "sonm.MinerStartReply")
	proto.RegisterType((*MinerStartReplyPort)(nil), "sonm.MinerStartReply.port")
	proto.RegisterType((*MinerStatusMapRequest)(nil), "sonm.MinerStatusMapRequest")
	proto.RegisterEnum("sonm.NATType", NATType_name, NATType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Miner service

type MinerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error)
	Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error)
	Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error)
	TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error)
}

type minerClient struct {
	cc *grpc.ClientConn
}

func NewMinerClient(cc *grpc.ClientConn) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error) {
	out := new(MinerHandshakeReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error) {
	out := new(MinerStartReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[0], c.cc, "/sonm.Miner/TasksStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTasksStatusClient{stream}
	return x, nil
}

type Miner_TasksStatusClient interface {
	Send(*MinerStatusMapRequest) error
	Recv() (*StatusMapReply, error)
	grpc.ClientStream
}

type minerTasksStatusClient struct {
	grpc.ClientStream
}

func (x *minerTasksStatusClient) Send(m *MinerStatusMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerTasksStatusClient) Recv() (*StatusMapReply, error) {
	m := new(StatusMapReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/TaskDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[1], c.cc, "/sonm.Miner/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Miner_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type minerTaskLogsClient struct {
	grpc.ClientStream
}

func (x *minerTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Miner service

type MinerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Info(context.Context, *MinerInfoRequest) (*InfoReply, error)
	Handshake(context.Context, *MinerHandshakeRequest) (*MinerHandshakeReply, error)
	Start(context.Context, *MinerStartRequest) (*MinerStartReply, error)
	Stop(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TasksStatus(Miner_TasksStatusServer) error
	TaskDetails(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	TaskLogs(*TaskLogsRequest, Miner_TaskLogsServer) error
}

func RegisterMinerServer(s *grpc.Server, srv MinerServer) {
	s.RegisterService(&_Miner_serviceDesc, srv)
}

func _Miner_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Info(ctx, req.(*MinerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerHandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Handshake(ctx, req.(*MinerHandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Start(ctx, req.(*MinerStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Stop(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TasksStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).TasksStatus(&minerTasksStatusServer{stream})
}

type Miner_TasksStatusServer interface {
	Send(*StatusMapReply) error
	Recv() (*MinerStatusMapRequest, error)
	grpc.ServerStream
}

type minerTasksStatusServer struct {
	grpc.ServerStream
}

func (x *minerTasksStatusServer) Send(m *StatusMapReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerTasksStatusServer) Recv() (*MinerStatusMapRequest, error) {
	m := new(MinerStatusMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_TaskDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).TaskDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/TaskDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).TaskDetails(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServer).TaskLogs(m, &minerTaskLogsServer{stream})
}

type Miner_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type minerTaskLogsServer struct {
	grpc.ServerStream
}

func (x *minerTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Miner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Miner_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Miner_Info_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Miner_Handshake_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Miner_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Miner_Stop_Handler,
		},
		{
			MethodName: "TaskDetails",
			Handler:    _Miner_TaskDetails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TasksStatus",
			Handler:       _Miner_TasksStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskLogs",
			Handler:       _Miner_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "miner.proto",
}

func init() { proto.RegisterFile("miner.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x25, 0xd2, 0x92, 0x46, 0x91, 0xcd, 0x8c, 0xf3, 0xc3, 0xb2, 0x41, 0x21, 0x10, 0x05,
	0xaa, 0x06, 0x85, 0xe0, 0xaa, 0x81, 0xd1, 0xe6, 0x52, 0x24, 0x16, 0x8d, 0x1a, 0x96, 0x25, 0x62,
	0x25, 0x23, 0xe8, 0xc9, 0x58, 0xcb, 0x5b, 0x7b, 0x21, 0x8a, 0x64, 0xc9, 0xa5, 0x01, 0x3e, 0x41,
	0x2f, 0x7d, 0x81, 0x3e, 0x60, 0x0f, 0x7d, 0x8b, 0x62, 0x77, 0x49, 0x85, 0x56, 0xe4, 0x13, 0x67,
	0xbe, 0x9d, 0x6f, 0xf6, 0x9b, 0x1f, 0x2e, 0x74, 0xd7, 0x3c, 0x62, 0xe9, 0x30, 0x49, 0x63, 0x11,
	0xa3, 0x99, 0xc5, 0xd1, 0xda, 0x3d, 0xe4, 0x91, 0xfc, 0x46, 0x9c, 0x6a, 0xd8, 0xc5, 0x25, 0x4d,
	0xe8, 0x0d, 0x0f, 0xb9, 0xe0, 0x2c, 0xd3, 0x98, 0x87, 0x60, 0x5f, 0x4a, 0xe6, 0x79, 0xf4, 0x47,
	0x4c, 0xd8, 0x9f, 0x39, 0xcb, 0x84, 0xf7, 0x3d, 0xbc, 0x54, 0xd8, 0x6f, 0x34, 0xba, 0xcd, 0xee,
	0xe9, 0x8a, 0x95, 0x07, 0x68, 0x43, 0xf3, 0x3e, 0xbf, 0x71, 0x8c, 0xbe, 0x31, 0xe8, 0x10, 0x69,
	0x7a, 0x7f, 0x1b, 0x70, 0xb4, 0x1d, 0x9b, 0x84, 0x05, 0xbe, 0x00, 0x4b, 0x09, 0x2a, 0x63, 0xb5,
	0x83, 0x27, 0xf0, 0xac, 0x2e, 0xc1, 0x69, 0xf4, 0x8d, 0x41, 0x77, 0x84, 0x43, 0x29, 0x73, 0x78,
	0x5a, 0x3b, 0x21, 0x8f, 0xe2, 0xf0, 0x3b, 0x68, 0x45, 0x54, 0x2c, 0x8a, 0x84, 0x39, 0xcd, 0xbe,
	0x31, 0x38, 0x18, 0xf5, 0x34, 0x65, 0xfa, 0x61, 0x21, 0x41, 0x52, 0x9d, 0x7a, 0xff, 0x34, 0xe1,
	0xb9, 0x92, 0x33, 0x17, 0x34, 0x15, 0x95, 0xec, 0x03, 0x68, 0xf0, 0xdb, 0x52, 0x49, 0x83, 0xdf,
	0xa2, 0x0b, 0xed, 0x94, 0xdd, 0xf1, 0x4c, 0xa4, 0x85, 0x92, 0xd0, 0x21, 0x1b, 0x5f, 0x0a, 0xe7,
	0x6b, 0x7a, 0xa7, 0x2f, 0xea, 0x10, 0xed, 0x20, 0x82, 0x49, 0x73, 0x71, 0xef, 0x98, 0x0a, 0x54,
	0x36, 0x7e, 0x84, 0x5e, 0xca, 0x32, 0x79, 0x4f, 0x10, 0x87, 0x7c, 0x59, 0x38, 0x96, 0xaa, 0xe6,
	0x4d, 0x59, 0x4d, 0x1c, 0x09, 0x2a, 0x95, 0x90, 0x7a, 0x0c, 0x79, 0x4c, 0xc1, 0x77, 0x60, 0xe5,
	0x99, 0xbc, 0x6d, 0x5f, 0x71, 0xbf, 0xd1, 0xdc, 0x05, 0xcd, 0x56, 0x84, 0x65, 0x71, 0x9e, 0x2e,
	0x55, 0xeb, 0x79, 0xca, 0xd6, 0x2c, 0x12, 0x19, 0xd1, 0xc1, 0xf8, 0x2d, 0xf4, 0x82, 0xfc, 0x26,
	0xe4, 0xcb, 0x0b, 0x56, 0x8c, 0xa9, 0xa0, 0x4e, 0x4b, 0xc9, 0xea, 0x25, 0x75, 0x10, 0x3d, 0x78,
	0xb6, 0x8c, 0xd7, 0x6b, 0x2e, 0x66, 0xd1, 0x5c, 0xc4, 0x89, 0xd3, 0xee, 0x1b, 0x83, 0x36, 0x79,
	0x84, 0xe1, 0x08, 0x9a, 0x2c, 0x7a, 0x70, 0x3a, 0xfd, 0xe6, 0xa0, 0x3b, 0xea, 0xeb, 0xdb, 0xbf,
	0xe8, 0xdf, 0xd0, 0x8f, 0x1e, 0xfc, 0x48, 0xa4, 0x05, 0x91, 0xc1, 0xee, 0x09, 0xb4, 0x2b, 0x40,
	0x2e, 0xc4, 0x8a, 0x15, 0xd5, 0x42, 0xac, 0x98, 0xea, 0xdf, 0x03, 0x0d, 0x73, 0x56, 0x36, 0x56,
	0x3b, 0xef, 0x1b, 0x3f, 0x1b, 0xde, 0x7f, 0x06, 0x1c, 0xd6, 0x73, 0xcb, 0x35, 0x79, 0x03, 0x9d,
	0x65, 0xd5, 0xa8, 0x32, 0xcb, 0x67, 0x00, 0x4f, 0xc0, 0x0a, 0xe2, 0x54, 0xc8, 0x3d, 0x79, 0x42,
	0x5f, 0x12, 0x16, 0x43, 0x15, 0xa2, 0xf5, 0x59, 0x89, 0xb4, 0xdd, 0xb7, 0x60, 0x4a, 0x43, 0xce,
	0xfd, 0x3c, 0xd8, 0xcc, 0x3d, 0x90, 0x53, 0x94, 0x78, 0x29, 0x4d, 0xd9, 0xee, 0x02, 0xe0, 0x73,
	0x82, 0x1d, 0xf5, 0x1c, 0xd7, 0xeb, 0xe9, 0x8e, 0xdc, 0xdd, 0x1a, 0x64, 0xaa, 0x7a, 0xad, 0xaf,
	0xcb, 0x3f, 0x68, 0x2e, 0xa8, 0xc8, 0xb3, 0x4b, 0x9a, 0x94, 0xad, 0x7c, 0xfb, 0x97, 0x01, 0xad,
	0x72, 0x6b, 0xb1, 0x0d, 0xe6, 0x74, 0x36, 0xf5, 0xed, 0x3d, 0xec, 0x42, 0xeb, 0xe3, 0x64, 0x76,
	0x7a, 0xe1, 0x8f, 0x6d, 0x43, 0xc2, 0x67, 0x57, 0x93, 0x89, 0xdd, 0xc0, 0x1e, 0x74, 0xe6, 0xbf,
	0x5f, 0x5e, 0xfa, 0x0b, 0x72, 0x7e, 0x6a, 0x37, 0xf1, 0x00, 0x80, 0xf8, 0x73, 0xe9, 0x2c, 0xfc,
	0xb1, 0x6d, 0xe2, 0x11, 0x1c, 0x06, 0x33, 0xb2, 0xb8, 0xae, 0x81, 0x16, 0xba, 0xf0, 0x6a, 0xc3,
	0xb9, 0xbe, 0x1a, 0x07, 0xd7, 0x67, 0xe7, 0xc4, 0xff, 0xf4, 0x61, 0x32, 0xb1, 0xf7, 0xe5, 0x35,
	0x57, 0xd3, 0x8b, 0xe9, 0xec, 0xd3, 0xd4, 0x6e, 0x8d, 0xfe, 0x6d, 0x82, 0xa5, 0x34, 0xe2, 0x0f,
	0x60, 0x06, 0x3c, 0xba, 0xc3, 0xe7, 0xba, 0x36, 0x69, 0x97, 0x72, 0xdd, 0xc3, 0x3a, 0x94, 0x84,
	0x85, 0xb7, 0x87, 0x3f, 0x82, 0x29, 0xdf, 0x0a, 0x7c, 0x55, 0xeb, 0x44, 0xed, 0xf1, 0xa8, 0x28,
	0x1a, 0xd2, 0x14, 0x1f, 0x3a, 0x9b, 0xe7, 0x01, 0xbf, 0xae, 0xf1, 0xb6, 0x1f, 0x18, 0xf7, 0xab,
	0xdd, 0x87, 0x3a, 0xcd, 0x2f, 0x60, 0xa9, 0x96, 0xe3, 0xeb, 0x27, 0x16, 0xd5, 0x7d, 0xb9, 0x73,
	0x3a, 0xde, 0x1e, 0xbe, 0x03, 0x53, 0xed, 0x7b, 0x19, 0x20, 0x6d, 0xfd, 0x93, 0x69, 0xde, 0xd1,
	0x36, 0xac, 0x59, 0x67, 0xd0, 0x95, 0x6e, 0xa6, 0xa7, 0xf8, 0x48, 0xf9, 0xf6, 0x60, 0xdd, 0x17,
	0x55, 0x8a, 0x0d, 0xae, 0x72, 0x0c, 0x8c, 0x63, 0x03, 0x7f, 0xd5, 0x79, 0xc6, 0x4c, 0x50, 0x1e,
	0x66, 0x95, 0x7c, 0x09, 0xe9, 0xf0, 0x2d, 0xf9, 0xf5, 0x03, 0x2d, 0xe4, 0x3d, 0xb4, 0x25, 0x38,
	0x89, 0xef, 0x32, 0xac, 0x05, 0x49, 0x7f, 0xab, 0x84, 0x0a, 0x3e, 0xbd, 0xcf, 0xa3, 0x95, 0xb7,
	0x77, 0x6c, 0xdc, 0xec, 0xab, 0x77, 0xfe, 0xa7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x1b,
	0x99, 0x26, 0x21, 0x06, 0x00, 0x00,
}
