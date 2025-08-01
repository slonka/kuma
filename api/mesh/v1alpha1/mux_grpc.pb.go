// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.0
// source: api/mesh/v1alpha1/mux.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MultiplexService_StreamMessage_FullMethodName = "/kuma.mesh.v1alpha1.MultiplexService/StreamMessage"
)

// MultiplexServiceClient is the client API for MultiplexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplexServiceClient interface {
	StreamMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error)
}

type multiplexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplexServiceClient(cc grpc.ClientConnInterface) MultiplexServiceClient {
	return &multiplexServiceClient{cc}
}

func (c *multiplexServiceClient) StreamMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MultiplexService_ServiceDesc.Streams[0], MultiplexService_StreamMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Message, Message]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MultiplexService_StreamMessageClient = grpc.BidiStreamingClient[Message, Message]

// MultiplexServiceServer is the server API for MultiplexService service.
// All implementations must embed UnimplementedMultiplexServiceServer
// for forward compatibility.
type MultiplexServiceServer interface {
	StreamMessage(grpc.BidiStreamingServer[Message, Message]) error
	mustEmbedUnimplementedMultiplexServiceServer()
}

// UnimplementedMultiplexServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMultiplexServiceServer struct{}

func (UnimplementedMultiplexServiceServer) StreamMessage(grpc.BidiStreamingServer[Message, Message]) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessage not implemented")
}
func (UnimplementedMultiplexServiceServer) mustEmbedUnimplementedMultiplexServiceServer() {}
func (UnimplementedMultiplexServiceServer) testEmbeddedByValue()                          {}

// UnsafeMultiplexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplexServiceServer will
// result in compilation errors.
type UnsafeMultiplexServiceServer interface {
	mustEmbedUnimplementedMultiplexServiceServer()
}

func RegisterMultiplexServiceServer(s grpc.ServiceRegistrar, srv MultiplexServiceServer) {
	// If the following call pancis, it indicates UnimplementedMultiplexServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MultiplexService_ServiceDesc, srv)
}

func _MultiplexService_StreamMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiplexServiceServer).StreamMessage(&grpc.GenericServerStream[Message, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MultiplexService_StreamMessageServer = grpc.BidiStreamingServer[Message, Message]

// MultiplexService_ServiceDesc is the grpc.ServiceDesc for MultiplexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.MultiplexService",
	HandlerType: (*MultiplexServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessage",
			Handler:       _MultiplexService_StreamMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mesh/v1alpha1/mux.proto",
}
