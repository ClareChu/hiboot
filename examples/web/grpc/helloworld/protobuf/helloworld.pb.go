// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package protobuf

/*
target package name
*/

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
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
	return fileDescriptor_helloworld_ddbc9c942d6b2c79, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
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

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_ddbc9c942d6b2c79, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type HolaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HolaRequest) Reset()         { *m = HolaRequest{} }
func (m *HolaRequest) String() string { return proto.CompactTextString(m) }
func (*HolaRequest) ProtoMessage()    {}
func (*HolaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_ddbc9c942d6b2c79, []int{2}
}
func (m *HolaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HolaRequest.Unmarshal(m, b)
}
func (m *HolaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HolaRequest.Marshal(b, m, deterministic)
}
func (dst *HolaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HolaRequest.Merge(dst, src)
}
func (m *HolaRequest) XXX_Size() int {
	return xxx_messageInfo_HolaRequest.Size(m)
}
func (m *HolaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HolaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HolaRequest proto.InternalMessageInfo

func (m *HolaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HolaReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HolaReply) Reset()         { *m = HolaReply{} }
func (m *HolaReply) String() string { return proto.CompactTextString(m) }
func (*HolaReply) ProtoMessage()    {}
func (*HolaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_ddbc9c942d6b2c79, []int{3}
}
func (m *HolaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HolaReply.Unmarshal(m, b)
}
func (m *HolaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HolaReply.Marshal(b, m, deterministic)
}
func (dst *HolaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HolaReply.Merge(dst, src)
}
func (m *HolaReply) XXX_Size() int {
	return xxx_messageInfo_HolaReply.Size(m)
}
func (m *HolaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HolaReply.DiscardUnknown(m)
}

var xxx_messageInfo_HolaReply proto.InternalMessageInfo

func (m *HolaReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protobuf.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "protobuf.HelloReply")
	proto.RegisterType((*HolaRequest)(nil), "protobuf.HolaRequest")
	proto.RegisterType((*HolaReply)(nil), "protobuf.HolaReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Sends a hello greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/protobuf.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Sends a hello greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

// HolaServiceClient is the client API for HolaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HolaServiceClient interface {
	// Sends a hola greeting
	SayHola(ctx context.Context, in *HolaRequest, opts ...grpc.CallOption) (*HolaReply, error)
}

type holaServiceClient struct {
	cc *grpc.ClientConn
}

func NewHolaServiceClient(cc *grpc.ClientConn) HolaServiceClient {
	return &holaServiceClient{cc}
}

func (c *holaServiceClient) SayHola(ctx context.Context, in *HolaRequest, opts ...grpc.CallOption) (*HolaReply, error) {
	out := new(HolaReply)
	err := c.cc.Invoke(ctx, "/protobuf.HolaService/SayHola", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HolaServiceServer is the server API for HolaService service.
type HolaServiceServer interface {
	// Sends a hola greeting
	SayHola(context.Context, *HolaRequest) (*HolaReply, error)
}

func RegisterHolaServiceServer(s *grpc.Server, srv HolaServiceServer) {
	s.RegisterService(&_HolaService_serviceDesc, srv)
}

func _HolaService_SayHola_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HolaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolaServiceServer).SayHola(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.HolaService/SayHola",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolaServiceServer).SayHola(ctx, req.(*HolaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HolaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.HolaService",
	HandlerType: (*HolaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHola",
			Handler:    _HolaService_SayHola_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_helloworld_ddbc9c942d6b2c79) }

var fileDescriptor_helloworld_ddbc9c942d6b2c79 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0x49, 0xa5, 0x69, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0xd9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b, 0xaa, 0xa6, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37,
	0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x08, 0xc6, 0x55, 0x52, 0xe4, 0xe2, 0xf6, 0xc8, 0xcf, 0x49,
	0xc4, 0x67, 0x94, 0x2a, 0x17, 0x27, 0x44, 0x09, 0x5e, 0x93, 0x8c, 0xbc, 0xa0, 0xae, 0x0a, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xb2, 0xe2, 0xe2, 0x08, 0x4e, 0xac, 0x04, 0x0b, 0x09, 0x89,
	0xe9, 0xc1, 0x1c, 0xaf, 0x87, 0xec, 0x72, 0x29, 0x11, 0x0c, 0xf1, 0x82, 0x9c, 0x4a, 0x25, 0x06,
	0x23, 0x37, 0x88, 0xab, 0x60, 0x46, 0x99, 0x73, 0xb1, 0x83, 0x8c, 0xca, 0xcf, 0x49, 0x14, 0x12,
	0x45, 0xd2, 0x81, 0x70, 0xb7, 0x94, 0x30, 0xba, 0x30, 0xd8, 0x1c, 0x27, 0x03, 0x2e, 0xe9, 0xcc,
	0x7c, 0xbd, 0xf4, 0xa2, 0x82, 0x64, 0xbd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0xd4, 0x62, 0x3d,
	0x44, 0xc0, 0x3a, 0xf1, 0x83, 0x2d, 0x0d, 0x07, 0xb1, 0x03, 0x40, 0xda, 0x03, 0x18, 0x93, 0xd8,
	0xc0, 0xe6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x83, 0x96, 0x55, 0x0c, 0x80, 0x01, 0x00,
	0x00,
}
