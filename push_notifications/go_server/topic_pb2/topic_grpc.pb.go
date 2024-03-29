// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/topic.proto

package topic_pb2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TopicSelectionClient is the client API for TopicSelection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicSelectionClient interface {
	// Random cat facts stream
	StreamMeowFacts(ctx context.Context, in *FactRequest, opts ...grpc.CallOption) (TopicSelection_StreamMeowFactsClient, error)
	// Color information format stream
	StreamColorInfo(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (TopicSelection_StreamColorInfoClient, error)
}

type topicSelectionClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicSelectionClient(cc grpc.ClientConnInterface) TopicSelectionClient {
	return &topicSelectionClient{cc}
}

func (c *topicSelectionClient) StreamMeowFacts(ctx context.Context, in *FactRequest, opts ...grpc.CallOption) (TopicSelection_StreamMeowFactsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TopicSelection_ServiceDesc.Streams[0], "/topic_selection.TopicSelection/StreamMeowFacts", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicSelectionStreamMeowFactsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TopicSelection_StreamMeowFactsClient interface {
	Recv() (*FactResponse, error)
	grpc.ClientStream
}

type topicSelectionStreamMeowFactsClient struct {
	grpc.ClientStream
}

func (x *topicSelectionStreamMeowFactsClient) Recv() (*FactResponse, error) {
	m := new(FactResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicSelectionClient) StreamColorInfo(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (TopicSelection_StreamColorInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &TopicSelection_ServiceDesc.Streams[1], "/topic_selection.TopicSelection/StreamColorInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicSelectionStreamColorInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TopicSelection_StreamColorInfoClient interface {
	Recv() (*ColorResponse, error)
	grpc.ClientStream
}

type topicSelectionStreamColorInfoClient struct {
	grpc.ClientStream
}

func (x *topicSelectionStreamColorInfoClient) Recv() (*ColorResponse, error) {
	m := new(ColorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TopicSelectionServer is the server API for TopicSelection service.
// All implementations must embed UnimplementedTopicSelectionServer
// for forward compatibility
type TopicSelectionServer interface {
	// Random cat facts stream
	StreamMeowFacts(*FactRequest, TopicSelection_StreamMeowFactsServer) error
	// Color information format stream
	StreamColorInfo(*ColorRequest, TopicSelection_StreamColorInfoServer) error
	mustEmbedUnimplementedTopicSelectionServer()
}

// UnimplementedTopicSelectionServer must be embedded to have forward compatible implementations.
type UnimplementedTopicSelectionServer struct {
}

func (UnimplementedTopicSelectionServer) StreamMeowFacts(*FactRequest, TopicSelection_StreamMeowFactsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMeowFacts not implemented")
}
func (UnimplementedTopicSelectionServer) StreamColorInfo(*ColorRequest, TopicSelection_StreamColorInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamColorInfo not implemented")
}
func (UnimplementedTopicSelectionServer) mustEmbedUnimplementedTopicSelectionServer() {}

// UnsafeTopicSelectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicSelectionServer will
// result in compilation errors.
type UnsafeTopicSelectionServer interface {
	mustEmbedUnimplementedTopicSelectionServer()
}

func RegisterTopicSelectionServer(s grpc.ServiceRegistrar, srv TopicSelectionServer) {
	s.RegisterService(&TopicSelection_ServiceDesc, srv)
}

func _TopicSelection_StreamMeowFacts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FactRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicSelectionServer).StreamMeowFacts(m, &topicSelectionStreamMeowFactsServer{stream})
}

type TopicSelection_StreamMeowFactsServer interface {
	Send(*FactResponse) error
	grpc.ServerStream
}

type topicSelectionStreamMeowFactsServer struct {
	grpc.ServerStream
}

func (x *topicSelectionStreamMeowFactsServer) Send(m *FactResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TopicSelection_StreamColorInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ColorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicSelectionServer).StreamColorInfo(m, &topicSelectionStreamColorInfoServer{stream})
}

type TopicSelection_StreamColorInfoServer interface {
	Send(*ColorResponse) error
	grpc.ServerStream
}

type topicSelectionStreamColorInfoServer struct {
	grpc.ServerStream
}

func (x *topicSelectionStreamColorInfoServer) Send(m *ColorResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TopicSelection_ServiceDesc is the grpc.ServiceDesc for TopicSelection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicSelection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "topic_selection.TopicSelection",
	HandlerType: (*TopicSelectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMeowFacts",
			Handler:       _TopicSelection_StreamMeowFacts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamColorInfo",
			Handler:       _TopicSelection_StreamColorInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/topic.proto",
}
