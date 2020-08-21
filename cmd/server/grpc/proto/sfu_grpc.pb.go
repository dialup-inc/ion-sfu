// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SFUClient is the client API for SFU service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SFUClient interface {
	Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingRequest, error)
}

type sFUClient struct {
	cc grpc.ClientConnInterface
}

func NewSFUClient(cc grpc.ClientConnInterface) SFUClient {
	return &sFUClient{cc}
}

func (c *sFUClient) Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SFU_serviceDesc.Streams[0], "/sfu.SFU/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &sFUSignalClient{stream}
	return x, nil
}

type SFU_SignalClient interface {
	Send(*SignalRequest) error
	Recv() (*SignalReply, error)
	grpc.ClientStream
}

type sFUSignalClient struct {
	grpc.ClientStream
}

func (x *sFUSignalClient) Send(m *SignalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sFUSignalClient) Recv() (*SignalReply, error) {
	m := new(SignalReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sFUClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingRequest, error) {
	out := new(PingRequest)
	err := c.cc.Invoke(ctx, "/sfu.SFU/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SFUServer is the server API for SFU service.
// All implementations must embed UnimplementedSFUServer
// for forward compatibility
type SFUServer interface {
	Signal(SFU_SignalServer) error
	Ping(context.Context, *PingRequest) (*PingRequest, error)
	mustEmbedUnimplementedSFUServer()
}

// UnimplementedSFUServer must be embedded to have forward compatible implementations.
type UnimplementedSFUServer struct {
}

func (*UnimplementedSFUServer) Signal(SFU_SignalServer) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (*UnimplementedSFUServer) Ping(context.Context, *PingRequest) (*PingRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedSFUServer) mustEmbedUnimplementedSFUServer() {}

func RegisterSFUServer(s *grpc.Server, srv SFUServer) {
	s.RegisterService(&_SFU_serviceDesc, srv)
}

func _SFU_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SFUServer).Signal(&sFUSignalServer{stream})
}

type SFU_SignalServer interface {
	Send(*SignalReply) error
	Recv() (*SignalRequest, error)
	grpc.ServerStream
}

type sFUSignalServer struct {
	grpc.ServerStream
}

func (x *sFUSignalServer) Send(m *SignalReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sFUSignalServer) Recv() (*SignalRequest, error) {
	m := new(SignalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SFU_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SFUServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sfu.SFU/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SFUServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SFU_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sfu.SFU",
	HandlerType: (*SFUServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SFU_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _SFU_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cmd/server/grpc/proto/sfu.proto",
}
