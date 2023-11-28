// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: adapterservice.proto

package pb

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

const (
	AdapterService_CreateStation_FullMethodName  = "/pb.AdapterService/CreateStation"
	AdapterService_DestroyStation_FullMethodName = "/pb.AdapterService/DestroyStation"
	AdapterService_Produce_FullMethodName        = "/pb.AdapterService/Produce"
	AdapterService_Consume_FullMethodName        = "/pb.AdapterService/Consume"
)

// AdapterServiceClient is the client API for AdapterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdapterServiceClient interface {
	CreateStation(ctx context.Context, in *CreateStationRequest, opts ...grpc.CallOption) (*Status, error)
	DestroyStation(ctx context.Context, in *DestroyStationRequest, opts ...grpc.CallOption) (*Status, error)
	Produce(ctx context.Context, opts ...grpc.CallOption) (AdapterService_ProduceClient, error)
	Consume(ctx context.Context, opts ...grpc.CallOption) (AdapterService_ConsumeClient, error)
}

type adapterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdapterServiceClient(cc grpc.ClientConnInterface) AdapterServiceClient {
	return &adapterServiceClient{cc}
}

func (c *adapterServiceClient) CreateStation(ctx context.Context, in *CreateStationRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, AdapterService_CreateStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterServiceClient) DestroyStation(ctx context.Context, in *DestroyStationRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, AdapterService_DestroyStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adapterServiceClient) Produce(ctx context.Context, opts ...grpc.CallOption) (AdapterService_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdapterService_ServiceDesc.Streams[0], AdapterService_Produce_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &adapterServiceProduceClient{stream}
	return x, nil
}

type AdapterService_ProduceClient interface {
	Send(*ProduceMessages) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type adapterServiceProduceClient struct {
	grpc.ClientStream
}

func (x *adapterServiceProduceClient) Send(m *ProduceMessages) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adapterServiceProduceClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adapterServiceClient) Consume(ctx context.Context, opts ...grpc.CallOption) (AdapterService_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdapterService_ServiceDesc.Streams[1], AdapterService_Consume_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &adapterServiceConsumeClient{stream}
	return x, nil
}

type AdapterService_ConsumeClient interface {
	Send(*ConsumeMessages) error
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type adapterServiceConsumeClient struct {
	grpc.ClientStream
}

func (x *adapterServiceConsumeClient) Send(m *ConsumeMessages) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adapterServiceConsumeClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdapterServiceServer is the server API for AdapterService service.
// All implementations must embed UnimplementedAdapterServiceServer
// for forward compatibility
type AdapterServiceServer interface {
	CreateStation(context.Context, *CreateStationRequest) (*Status, error)
	DestroyStation(context.Context, *DestroyStationRequest) (*Status, error)
	Produce(AdapterService_ProduceServer) error
	Consume(AdapterService_ConsumeServer) error
	mustEmbedUnimplementedAdapterServiceServer()
}

// UnimplementedAdapterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdapterServiceServer struct {
}

func (UnimplementedAdapterServiceServer) CreateStation(context.Context, *CreateStationRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStation not implemented")
}
func (UnimplementedAdapterServiceServer) DestroyStation(context.Context, *DestroyStationRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyStation not implemented")
}
func (UnimplementedAdapterServiceServer) Produce(AdapterService_ProduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedAdapterServiceServer) Consume(AdapterService_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedAdapterServiceServer) mustEmbedUnimplementedAdapterServiceServer() {}

// UnsafeAdapterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdapterServiceServer will
// result in compilation errors.
type UnsafeAdapterServiceServer interface {
	mustEmbedUnimplementedAdapterServiceServer()
}

func RegisterAdapterServiceServer(s grpc.ServiceRegistrar, srv AdapterServiceServer) {
	s.RegisterService(&AdapterService_ServiceDesc, srv)
}

func _AdapterService_CreateStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServiceServer).CreateStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdapterService_CreateStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServiceServer).CreateStation(ctx, req.(*CreateStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdapterService_DestroyStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdapterServiceServer).DestroyStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdapterService_DestroyStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdapterServiceServer).DestroyStation(ctx, req.(*DestroyStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdapterService_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdapterServiceServer).Produce(&adapterServiceProduceServer{stream})
}

type AdapterService_ProduceServer interface {
	SendAndClose(*Status) error
	Recv() (*ProduceMessages, error)
	grpc.ServerStream
}

type adapterServiceProduceServer struct {
	grpc.ServerStream
}

func (x *adapterServiceProduceServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adapterServiceProduceServer) Recv() (*ProduceMessages, error) {
	m := new(ProduceMessages)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AdapterService_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdapterServiceServer).Consume(&adapterServiceConsumeServer{stream})
}

type AdapterService_ConsumeServer interface {
	Send(*ConsumeResponse) error
	Recv() (*ConsumeMessages, error)
	grpc.ServerStream
}

type adapterServiceConsumeServer struct {
	grpc.ServerStream
}

func (x *adapterServiceConsumeServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adapterServiceConsumeServer) Recv() (*ConsumeMessages, error) {
	m := new(ConsumeMessages)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdapterService_ServiceDesc is the grpc.ServiceDesc for AdapterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdapterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AdapterService",
	HandlerType: (*AdapterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStation",
			Handler:    _AdapterService_CreateStation_Handler,
		},
		{
			MethodName: "DestroyStation",
			Handler:    _AdapterService_DestroyStation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _AdapterService_Produce_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Consume",
			Handler:       _AdapterService_Consume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "adapterservice.proto",
}
