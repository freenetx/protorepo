// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: channel/channel.proto

package channelpb

import (
	context "context"
	protorepo_share_go "github.com/freenetx/protorepo/gens/go/protorepo-share-go"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChannelServiceClient is the client API for ChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelServiceClient interface {
	Test(ctx context.Context, in *protorepo_share_go.Empty, opts ...grpc.CallOption) (*protorepo_share_go.Empty, error)
	Channel(ctx context.Context, opts ...grpc.CallOption) (ChannelService_ChannelClient, error)
}

type channelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelServiceClient(cc grpc.ClientConnInterface) ChannelServiceClient {
	return &channelServiceClient{cc}
}

func (c *channelServiceClient) Test(ctx context.Context, in *protorepo_share_go.Empty, opts ...grpc.CallOption) (*protorepo_share_go.Empty, error) {
	out := new(protorepo_share_go.Empty)
	err := c.cc.Invoke(ctx, "/channel.api.ChannelService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) Channel(ctx context.Context, opts ...grpc.CallOption) (ChannelService_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChannelService_ServiceDesc.Streams[0], "/channel.api.ChannelService/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelServiceChannelClient{stream}
	return x, nil
}

type ChannelService_ChannelClient interface {
	Send(*ChannelRequest) error
	Recv() (*ReceiveResponse, error)
	grpc.ClientStream
}

type channelServiceChannelClient struct {
	grpc.ClientStream
}

func (x *channelServiceChannelClient) Send(m *ChannelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *channelServiceChannelClient) Recv() (*ReceiveResponse, error) {
	m := new(ReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelServiceServer is the server API for ChannelService service.
// All implementations must embed UnimplementedChannelServiceServer
// for forward compatibility
type ChannelServiceServer interface {
	Test(context.Context, *protorepo_share_go.Empty) (*protorepo_share_go.Empty, error)
	Channel(ChannelService_ChannelServer) error
	mustEmbedUnimplementedChannelServiceServer()
}

// UnimplementedChannelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServiceServer struct {
}

func (UnimplementedChannelServiceServer) Test(context.Context, *protorepo_share_go.Empty) (*protorepo_share_go.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedChannelServiceServer) Channel(ChannelService_ChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method Channel not implemented")
}
func (UnimplementedChannelServiceServer) mustEmbedUnimplementedChannelServiceServer() {}

// UnsafeChannelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServiceServer will
// result in compilation errors.
type UnsafeChannelServiceServer interface {
	mustEmbedUnimplementedChannelServiceServer()
}

func RegisterChannelServiceServer(s grpc.ServiceRegistrar, srv ChannelServiceServer) {
	s.RegisterService(&ChannelService_ServiceDesc, srv)
}

func _ChannelService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protorepo_share_go.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.api.ChannelService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).Test(ctx, req.(*protorepo_share_go.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChannelServiceServer).Channel(&channelServiceChannelServer{stream})
}

type ChannelService_ChannelServer interface {
	Send(*ReceiveResponse) error
	Recv() (*ChannelRequest, error)
	grpc.ServerStream
}

type channelServiceChannelServer struct {
	grpc.ServerStream
}

func (x *channelServiceChannelServer) Send(m *ReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *channelServiceChannelServer) Recv() (*ChannelRequest, error) {
	m := new(ChannelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelService_ServiceDesc is the grpc.ServiceDesc for ChannelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.api.ChannelService",
	HandlerType: (*ChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _ChannelService_Test_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _ChannelService_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "channel/channel.proto",
}
