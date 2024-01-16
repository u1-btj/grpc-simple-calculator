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
	StreamMeowFacts(ctx context.Context, in *FactRequest, opts ...grpc.CallOption) (TopicSelection_StreamMeowFactsClient, error)
	StreamAnimeQuotes(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (TopicSelection_StreamAnimeQuotesClient, error)
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

func (c *topicSelectionClient) StreamAnimeQuotes(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (TopicSelection_StreamAnimeQuotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &TopicSelection_ServiceDesc.Streams[1], "/topic_selection.TopicSelection/StreamAnimeQuotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicSelectionStreamAnimeQuotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TopicSelection_StreamAnimeQuotesClient interface {
	Recv() (*QuoteResponse, error)
	grpc.ClientStream
}

type topicSelectionStreamAnimeQuotesClient struct {
	grpc.ClientStream
}

func (x *topicSelectionStreamAnimeQuotesClient) Recv() (*QuoteResponse, error) {
	m := new(QuoteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TopicSelectionServer is the server API for TopicSelection service.
// All implementations must embed UnimplementedTopicSelectionServer
// for forward compatibility
type TopicSelectionServer interface {
	StreamMeowFacts(*FactRequest, TopicSelection_StreamMeowFactsServer) error
	StreamAnimeQuotes(*QuoteRequest, TopicSelection_StreamAnimeQuotesServer) error
	mustEmbedUnimplementedTopicSelectionServer()
}

// UnimplementedTopicSelectionServer must be embedded to have forward compatible implementations.
type UnimplementedTopicSelectionServer struct {
}

func (UnimplementedTopicSelectionServer) StreamMeowFacts(*FactRequest, TopicSelection_StreamMeowFactsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMeowFacts not implemented")
}
func (UnimplementedTopicSelectionServer) StreamAnimeQuotes(*QuoteRequest, TopicSelection_StreamAnimeQuotesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAnimeQuotes not implemented")
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

func _TopicSelection_StreamAnimeQuotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuoteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicSelectionServer).StreamAnimeQuotes(m, &topicSelectionStreamAnimeQuotesServer{stream})
}

type TopicSelection_StreamAnimeQuotesServer interface {
	Send(*QuoteResponse) error
	grpc.ServerStream
}

type topicSelectionStreamAnimeQuotesServer struct {
	grpc.ServerStream
}

func (x *topicSelectionStreamAnimeQuotesServer) Send(m *QuoteResponse) error {
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
			StreamName:    "StreamAnimeQuotes",
			Handler:       _TopicSelection_StreamAnimeQuotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/topic.proto",
}
