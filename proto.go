// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer.proto

package peer

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

type PeerRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerRequest) Reset()         { *m = PeerRequest{} }
func (m *PeerRequest) String() string { return proto.CompactTextString(m) }
func (*PeerRequest) ProtoMessage()    {}
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_055ae5a865fc1c9e, []int{0}
}

func (m *PeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerRequest.Unmarshal(m, b)
}
func (m *PeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerRequest.Marshal(b, m, deterministic)
}
func (m *PeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerRequest.Merge(m, src)
}
func (m *PeerRequest) XXX_Size() int {
	return xxx_messageInfo_PeerRequest.Size(m)
}
func (m *PeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerRequest proto.InternalMessageInfo

func (m *PeerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PeerResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_055ae5a865fc1c9e, []int{1}
}

func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerResponse.Unmarshal(m, b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
}
func (m *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(m, src)
}
func (m *PeerResponse) XXX_Size() int {
	return xxx_messageInfo_PeerResponse.Size(m)
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerRequest)(nil), "peer.PeerRequest")
	proto.RegisterType((*PeerResponse)(nil), "peer.PeerResponse")
}

func init() { proto.RegisterFile("peer.proto", fileDescriptor_055ae5a865fc1c9e) }

var fileDescriptor_055ae5a865fc1c9e = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x4d, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xf4, 0xb9, 0xb8, 0x03, 0x52,
	0x53, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3,
	0x25, 0x98, 0xc0, 0x02, 0x20, 0xa6, 0x92, 0x01, 0x17, 0x0f, 0x44, 0x43, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0x2a, 0x61, 0x1d, 0x46, 0x76, 0x5c, 0x1c, 0x20, 0x1d, 0xae, 0xc9, 0x19, 0xf9, 0x42, 0x46,
	0x5c, 0xec, 0x30, 0xab, 0x04, 0xf5, 0xc0, 0x8e, 0x41, 0xb2, 0x5d, 0x4a, 0x08, 0x59, 0x08, 0x62,
	0xbe, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xf9,
	0xe5, 0x77, 0xbd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeerEchoClient is the client API for PeerEcho service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerEchoClient interface {
	Request(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerResponse, error)
}

type peerEchoClient struct {
	cc *grpc.ClientConn
}

func NewPeerEchoClient(cc *grpc.ClientConn) PeerEchoClient {
	return &peerEchoClient{cc}
}

func (c *peerEchoClient) Request(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/peer.PeerEcho/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerEchoServer is the server API for PeerEcho service.
type PeerEchoServer interface {
	Request(context.Context, *PeerRequest) (*PeerResponse, error)
}

// UnimplementedPeerEchoServer can be embedded to have forward compatible implementations.
type UnimplementedPeerEchoServer struct {
}

func (*UnimplementedPeerEchoServer) Request(ctx context.Context, req *PeerRequest) (*PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}

func RegisterPeerEchoServer(s *grpc.Server, srv PeerEchoServer) {
	s.RegisterService(&_PeerEcho_serviceDesc, srv)
}

func _PeerEcho_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerEchoServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.PeerEcho/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerEchoServer).Request(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerEcho_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peer.PeerEcho",
	HandlerType: (*PeerEchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _PeerEcho_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer.proto",
}
