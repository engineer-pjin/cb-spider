// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meerkat.proto

package common

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// The request message with empty
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// The response message containing a ChildKat status
type Status struct {
	ServerID             string   `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Reserved             string   `protobuf:"bytes,4,opt,name=Reserved,proto3" json:"Reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Status) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Status) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

// The input message containing for a ChildKat to run
type Command struct {
	CMDID                string   `protobuf:"bytes,1,opt,name=CMDID,proto3" json:"CMDID,omitempty"`
	CMDTYPE              string   `protobuf:"bytes,2,opt,name=CMDTYPE,proto3" json:"CMDTYPE,omitempty"`
	CMD                  string   `protobuf:"bytes,3,opt,name=CMD,proto3" json:"CMD,omitempty"`
	Time                 string   `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
	Reserved             string   `protobuf:"bytes,5,opt,name=Reserved,proto3" json:"Reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{2}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Command.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return m.Size()
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCMDID() string {
	if m != nil {
		return m.CMDID
	}
	return ""
}

func (m *Command) GetCMDTYPE() string {
	if m != nil {
		return m.CMDTYPE
	}
	return ""
}

func (m *Command) GetCMD() string {
	if m != nil {
		return m.CMD
	}
	return ""
}

func (m *Command) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Command) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

// The response message containing a ChildKat Command Result
type CommandResult struct {
	ServerID             string   `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	CMD                  string   `protobuf:"bytes,2,opt,name=CMD,proto3" json:"CMD,omitempty"`
	Result               string   `protobuf:"bytes,3,opt,name=Result,proto3" json:"Result,omitempty"`
	Time                 string   `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
	Reserved             string   `protobuf:"bytes,5,opt,name=Reserved,proto3" json:"Reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResult) Reset()         { *m = CommandResult{} }
func (m *CommandResult) String() string { return proto.CompactTextString(m) }
func (*CommandResult) ProtoMessage()    {}
func (*CommandResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{3}
}
func (m *CommandResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommandResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommandResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommandResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResult.Merge(m, src)
}
func (m *CommandResult) XXX_Size() int {
	return m.Size()
}
func (m *CommandResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResult.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResult proto.InternalMessageInfo

func (m *CommandResult) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *CommandResult) GetCMD() string {
	if m != nil {
		return m.CMD
	}
	return ""
}

func (m *CommandResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *CommandResult) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *CommandResult) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "meerkatrpc.Empty")
	proto.RegisterType((*Status)(nil), "meerkatrpc.Status")
	proto.RegisterType((*Command)(nil), "meerkatrpc.Command")
	proto.RegisterType((*CommandResult)(nil), "meerkatrpc.CommandResult")
}

func init() { proto.RegisterFile("meerkat.proto", fileDescriptor_c32248fb0568de66) }

var fileDescriptor_c32248fb0568de66 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x9b, 0xfe, 0x7e, 0xdf, 0x91, 0x8a, 0x8e, 0x52, 0x62, 0x17, 0x41, 0xb2, 0x72, 0xd3,
	0x04, 0xea, 0xca, 0x8d, 0x0b, 0x93, 0xa2, 0x82, 0x05, 0x49, 0x0b, 0xa2, 0xbb, 0xfc, 0x0c, 0x76,
	0x30, 0x93, 0x09, 0x93, 0x19, 0x41, 0x5c, 0x7b, 0x0f, 0x5e, 0x92, 0x4b, 0x2f, 0x41, 0xea, 0x8d,
	0x48, 0x93, 0x49, 0x1a, 0x45, 0x05, 0x77, 0xf3, 0xbc, 0x99, 0x73, 0x9e, 0x99, 0x39, 0x81, 0x3e,
	0xc5, 0x98, 0xdf, 0xf9, 0xc2, 0x4a, 0x39, 0x13, 0x0c, 0x81, 0x42, 0x9e, 0x86, 0x66, 0x0f, 0x3a,
	0x13, 0x9a, 0x8a, 0x07, 0x33, 0x86, 0xee, 0x4c, 0xf8, 0x42, 0x66, 0x68, 0x08, 0xff, 0x66, 0x98,
	0xdf, 0x63, 0x7e, 0xee, 0xea, 0xda, 0xbe, 0x76, 0xf0, 0xdf, 0xab, 0x18, 0x0d, 0xca, 0x5d, 0x7a,
	0x33, 0xff, 0x52, 0xd6, 0x20, 0x68, 0xcf, 0x09, 0xc5, 0x7a, 0x2b, 0x4f, 0xf3, 0xf5, 0xaa, 0x8f,
	0x87, 0xb3, 0x55, 0x65, 0xa4, 0xb7, 0x8b, 0x3e, 0x25, 0x9b, 0x8f, 0xd0, 0x73, 0x18, 0xa5, 0x7e,
	0x12, 0xa1, 0x5d, 0xe8, 0x38, 0x53, 0xb7, 0x72, 0x15, 0x80, 0x74, 0xe8, 0x39, 0x53, 0x77, 0x7e,
	0x7d, 0x39, 0x51, 0xa6, 0x12, 0xd1, 0x16, 0xb4, 0x9c, 0xa9, 0xab, 0x4c, 0xab, 0x65, 0x25, 0x6f,
	0xff, 0x20, 0xef, 0x7c, 0x91, 0x3f, 0x69, 0xd0, 0x57, 0x76, 0x0f, 0x67, 0x32, 0x16, 0xbf, 0x5e,
	0x59, 0xf9, 0x9a, 0x6b, 0xdf, 0x00, 0xba, 0x45, 0x9d, 0x3a, 0x84, 0xa2, 0xbf, 0x9e, 0x63, 0x7c,
	0x06, 0x1b, 0xce, 0x82, 0xc4, 0x91, 0x7a, 0xc3, 0x23, 0xd8, 0x3c, 0xc5, 0xa2, 0x9e, 0x6c, 0x5b,
	0xeb, 0x49, 0x59, 0xf9, 0x98, 0x86, 0xa8, 0x1e, 0x15, 0xdb, 0xcc, 0xc6, 0xf8, 0x02, 0xc0, 0x93,
	0x49, 0xf9, 0xa2, 0xc7, 0x9f, 0x68, 0xa7, 0x5e, 0xa1, 0xc2, 0xe1, 0xde, 0x37, 0x61, 0x71, 0x0b,
	0xb3, 0x71, 0x72, 0xf5, 0xb2, 0x34, 0xb4, 0xd7, 0xa5, 0xa1, 0xbd, 0x2d, 0x0d, 0xed, 0xf9, 0xdd,
	0x68, 0xdc, 0x4c, 0x6e, 0x89, 0x58, 0xc8, 0xc0, 0x0a, 0x19, 0xb5, 0xc3, 0x98, 0xc9, 0x68, 0x14,
	0xf8, 0x9c, 0x64, 0xc2, 0xb7, 0xc3, 0x60, 0x94, 0xa5, 0x24, 0xc2, 0xdc, 0xf6, 0x53, 0x32, 0xe2,
	0x32, 0x11, 0x84, 0x62, 0x5b, 0x35, 0xaf, 0x38, 0x64, 0x94, 0xb2, 0x24, 0xe8, 0xe6, 0xff, 0xdf,
	0xe1, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x6b, 0x8c, 0x82, 0x90, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChildStatusClient is the client API for ChildStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChildStatusClient interface {
	// Sends a request of Resource status
	GetChildStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
}

type childStatusClient struct {
	cc *grpc.ClientConn
}

func NewChildStatusClient(cc *grpc.ClientConn) ChildStatusClient {
	return &childStatusClient{cc}
}

func (c *childStatusClient) GetChildStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/meerkatrpc.ChildStatus/GetChildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChildStatusServer is the server API for ChildStatus service.
type ChildStatusServer interface {
	// Sends a request of Resource status
	GetChildStatus(context.Context, *Empty) (*Status, error)
}

// UnimplementedChildStatusServer can be embedded to have forward compatible implementations.
type UnimplementedChildStatusServer struct {
}

func (*UnimplementedChildStatusServer) GetChildStatus(ctx context.Context, req *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildStatus not implemented")
}

func RegisterChildStatusServer(s *grpc.Server, srv ChildStatusServer) {
	s.RegisterService(&_ChildStatus_serviceDesc, srv)
}

func _ChildStatus_GetChildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChildStatusServer).GetChildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meerkatrpc.ChildStatus/GetChildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChildStatusServer).GetChildStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChildStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meerkatrpc.ChildStatus",
	HandlerType: (*ChildStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChildStatus",
			Handler:    _ChildStatus_GetChildStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meerkat.proto",
}

// RunCommandClient is the client API for RunCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunCommandClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResult, error)
}

type runCommandClient struct {
	cc *grpc.ClientConn
}

func NewRunCommandClient(cc *grpc.ClientConn) RunCommandClient {
	return &runCommandClient{cc}
}

func (c *runCommandClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResult, error) {
	out := new(CommandResult)
	err := c.cc.Invoke(ctx, "/meerkatrpc.RunCommand/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunCommandServer is the server API for RunCommand service.
type RunCommandServer interface {
	RunCommand(context.Context, *Command) (*CommandResult, error)
}

// UnimplementedRunCommandServer can be embedded to have forward compatible implementations.
type UnimplementedRunCommandServer struct {
}

func (*UnimplementedRunCommandServer) RunCommand(ctx context.Context, req *Command) (*CommandResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}

func RegisterRunCommandServer(s *grpc.Server, srv RunCommandServer) {
	s.RegisterService(&_RunCommand_serviceDesc, srv)
}

func _RunCommand_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunCommandServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meerkatrpc.RunCommand/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunCommandServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meerkatrpc.RunCommand",
	HandlerType: (*RunCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _RunCommand_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meerkat.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Command) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Command) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Command) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CMD) > 0 {
		i -= len(m.CMD)
		copy(dAtA[i:], m.CMD)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.CMD)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CMDTYPE) > 0 {
		i -= len(m.CMDTYPE)
		copy(dAtA[i:], m.CMDTYPE)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.CMDTYPE)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CMDID) > 0 {
		i -= len(m.CMDID)
		copy(dAtA[i:], m.CMDID)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.CMDID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommandResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommandResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommandResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CMD) > 0 {
		i -= len(m.CMD)
		copy(dAtA[i:], m.CMD)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.CMD)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeerkat(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeerkat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Command) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CMDID)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.CMDTYPE)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.CMD)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommandResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.CMD)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMeerkat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeerkat(x uint64) (n int) {
	return sovMeerkat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Command) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Command: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Command: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CMDID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CMDID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CMDTYPE", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CMDTYPE = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CMD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CMD = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommandResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommandResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommandResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CMD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CMD = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMeerkat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeerkat
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMeerkat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMeerkat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeerkat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeerkat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeerkat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeerkat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeerkat = fmt.Errorf("proto: unexpected end of group")
)