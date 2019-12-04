// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReplay struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReplay) Reset()         { *m = HelloReplay{} }
func (m *HelloReplay) String() string { return proto.CompactTextString(m) }
func (*HelloReplay) ProtoMessage()    {}
func (*HelloReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReplay.Unmarshal(m, b)
}
func (m *HelloReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReplay.Marshal(b, m, deterministic)
}
func (m *HelloReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReplay.Merge(m, src)
}
func (m *HelloReplay) XXX_Size() int {
	return xxx_messageInfo_HelloReplay.Size(m)
}
func (m *HelloReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReplay.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReplay proto.InternalMessageInfo

func (m *HelloReplay) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMessage) Reset()         { *m = HelloMessage{} }
func (m *HelloMessage) String() string { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()    {}
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *HelloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMessage.Unmarshal(m, b)
}
func (m *HelloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMessage.Marshal(b, m, deterministic)
}
func (m *HelloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMessage.Merge(m, src)
}
func (m *HelloMessage) XXX_Size() int {
	return xxx_messageInfo_HelloMessage.Size(m)
}
func (m *HelloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMessage proto.InternalMessageInfo

func (m *HelloMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protocol.HelloRequest")
	proto.RegisterType((*HelloReplay)(nil), "protocol.HelloReplay")
	proto.RegisterType((*HelloMessage)(nil), "protocol.HelloMessage")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x4a,
	0x5c, 0x3c, 0x1e, 0x20, 0x89, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x9d, 0x8b,
	0x1b, 0xaa, 0xa6, 0x20, 0x27, 0xb1, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x1d, 0xa6, 0x0a, 0xc6, 0x55, 0x52, 0x80, 0x1a, 0xe6, 0x0b, 0xe1, 0x0b, 0x09, 0x70, 0x31, 0xe7,
	0x16, 0xa7, 0x43, 0x55, 0x81, 0x98, 0x46, 0xdd, 0x8c, 0x50, 0xb3, 0x82, 0x53, 0x8b, 0xca, 0x52,
	0x8b, 0x84, 0xac, 0xb9, 0x38, 0x82, 0x13, 0x2b, 0xc1, 0x22, 0x42, 0x62, 0x7a, 0x30, 0x57, 0xe9,
	0x21, 0x3b, 0x49, 0x4a, 0x14, 0x43, 0x1c, 0xe4, 0x0c, 0x25, 0x06, 0x21, 0x7b, 0x2e, 0x6e, 0xf7,
	0xd4, 0x12, 0x88, 0x8d, 0xc5, 0xe9, 0x38, 0xf5, 0xa3, 0x8b, 0x43, 0x5d, 0xa7, 0xc4, 0x90, 0xc4,
	0x06, 0x96, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xc8, 0x93, 0xa6, 0x1c, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerClient interface {
	// 创建第一个接口
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error)
	// 创建第二个接口
	GetHelloMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error)
}

type helloServerClient struct {
	cc *grpc.ClientConn
}

func NewHelloServerClient(cc *grpc.ClientConn) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error) {
	out := new(HelloReplay)
	err := c.cc.Invoke(ctx, "/protocol.HelloServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerClient) GetHelloMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error) {
	out := new(HelloMessage)
	err := c.cc.Invoke(ctx, "/protocol.HelloServer/GetHelloMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
type HelloServerServer interface {
	// 创建第一个接口
	SayHello(context.Context, *HelloRequest) (*HelloReplay, error)
	// 创建第二个接口
	GetHelloMsg(context.Context, *HelloRequest) (*HelloMessage, error)
}

// UnimplementedHelloServerServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (*UnimplementedHelloServerServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServerServer) GetHelloMsg(ctx context.Context, req *HelloRequest) (*HelloMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelloMsg not implemented")
}

func RegisterHelloServerServer(s *grpc.Server, srv HelloServerServer) {
	s.RegisterService(&_HelloServer_serviceDesc, srv)
}

func _HelloServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.HelloServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloServer_GetHelloMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).GetHelloMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.HelloServer/GetHelloMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).GetHelloMsg(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServer_SayHello_Handler,
		},
		{
			MethodName: "GetHelloMsg",
			Handler:    _HelloServer_GetHelloMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}