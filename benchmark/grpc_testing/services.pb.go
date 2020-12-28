// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "github.com/Hyperledger-TWGC/grpc"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BenchmarkServiceClient is the client API for BenchmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/Hyperledger-TWGC/grpc#ClientConn.NewStream.
type BenchmarkServiceClient interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_StreamingCallClient, error)
}

type benchmarkServiceClient struct {
	cc *grpc.ClientConn
}

func NewBenchmarkServiceClient(cc *grpc.ClientConn) BenchmarkServiceClient {
	return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.BenchmarkService/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchmarkServiceClient) StreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_StreamingCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BenchmarkService_serviceDesc.Streams[0], "/grpc.testing.BenchmarkService/StreamingCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkServiceStreamingCallClient{stream}
	return x, nil
}

type BenchmarkService_StreamingCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type benchmarkServiceStreamingCallClient struct {
	grpc.ClientStream
}

func (x *benchmarkServiceStreamingCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *benchmarkServiceStreamingCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BenchmarkServiceServer is the server API for BenchmarkService service.
type BenchmarkServiceServer interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(BenchmarkService_StreamingCallServer) error
}

func RegisterBenchmarkServiceServer(s *grpc.Server, srv BenchmarkServiceServer) {
	s.RegisterService(&_BenchmarkService_serviceDesc, srv)
}

func _BenchmarkService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.BenchmarkService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchmarkService_StreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).StreamingCall(&benchmarkServiceStreamingCallServer{stream})
}

type BenchmarkService_StreamingCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type benchmarkServiceStreamingCallServer struct {
	grpc.ServerStream
}

func (x *benchmarkServiceStreamingCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *benchmarkServiceStreamingCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BenchmarkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _BenchmarkService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingCall",
			Handler:       _BenchmarkService_StreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/Hyperledger-TWGC/grpc#ClientConn.NewStream.
type WorkerServiceClient interface {
	// Start server with specified workload.
	// First request sent specifies the ServerConfig followed by ServerStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test server
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunServer(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunServerClient, error)
	// Start client with specified workload.
	// First request sent specifies the ClientConfig followed by ClientStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test client
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunClient(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunClientClient, error)
	// Just return the core count - unary call
	CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error)
	// Quit this worker
	QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
}

type workerServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkerServiceClient(cc *grpc.ClientConn) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) RunServer(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkerService_serviceDesc.Streams[0], "/grpc.testing.WorkerService/RunServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerServiceRunServerClient{stream}
	return x, nil
}

type WorkerService_RunServerClient interface {
	Send(*ServerArgs) error
	Recv() (*ServerStatus, error)
	grpc.ClientStream
}

type workerServiceRunServerClient struct {
	grpc.ClientStream
}

func (x *workerServiceRunServerClient) Send(m *ServerArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerServiceRunServerClient) Recv() (*ServerStatus, error) {
	m := new(ServerStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerServiceClient) RunClient(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkerService_serviceDesc.Streams[1], "/grpc.testing.WorkerService/RunClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerServiceRunClientClient{stream}
	return x, nil
}

type WorkerService_RunClientClient interface {
	Send(*ClientArgs) error
	Recv() (*ClientStatus, error)
	grpc.ClientStream
}

type workerServiceRunClientClient struct {
	grpc.ClientStream
}

func (x *workerServiceRunClientClient) Send(m *ClientArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerServiceRunClientClient) Recv() (*ClientStatus, error) {
	m := new(ClientStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerServiceClient) CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error) {
	out := new(CoreResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.WorkerService/CoreCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.testing.WorkerService/QuitWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
type WorkerServiceServer interface {
	// Start server with specified workload.
	// First request sent specifies the ServerConfig followed by ServerStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test server
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunServer(WorkerService_RunServerServer) error
	// Start client with specified workload.
	// First request sent specifies the ClientConfig followed by ClientStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test client
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunClient(WorkerService_RunClientServer) error
	// Just return the core count - unary call
	CoreCount(context.Context, *CoreRequest) (*CoreResponse, error)
	// Quit this worker
	QuitWorker(context.Context, *Void) (*Void, error)
}

func RegisterWorkerServiceServer(s *grpc.Server, srv WorkerServiceServer) {
	s.RegisterService(&_WorkerService_serviceDesc, srv)
}

func _WorkerService_RunServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).RunServer(&workerServiceRunServerServer{stream})
}

type WorkerService_RunServerServer interface {
	Send(*ServerStatus) error
	Recv() (*ServerArgs, error)
	grpc.ServerStream
}

type workerServiceRunServerServer struct {
	grpc.ServerStream
}

func (x *workerServiceRunServerServer) Send(m *ServerStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerServiceRunServerServer) Recv() (*ServerArgs, error) {
	m := new(ServerArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WorkerService_RunClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).RunClient(&workerServiceRunClientServer{stream})
}

type WorkerService_RunClientServer interface {
	Send(*ClientStatus) error
	Recv() (*ClientArgs, error)
	grpc.ServerStream
}

type workerServiceRunClientServer struct {
	grpc.ServerStream
}

func (x *workerServiceRunClientServer) Send(m *ClientStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerServiceRunClientServer) Recv() (*ClientArgs, error) {
	m := new(ClientArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WorkerService_CoreCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CoreCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.WorkerService/CoreCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CoreCount(ctx, req.(*CoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_QuitWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).QuitWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.WorkerService/QuitWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).QuitWorker(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoreCount",
			Handler:    _WorkerService_CoreCount_Handler,
		},
		{
			MethodName: "QuitWorker",
			Handler:    _WorkerService_QuitWorker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunServer",
			Handler:       _WorkerService_RunServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunClient",
			Handler:       _WorkerService_RunClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_services_bf68f4d7cbd0e0a1) }

var fileDescriptor_services_bf68f4d7cbd0e0a1 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x07, 0xa1, 0xc1, 0x2e, 0x92, 0x93, 0x46, 0x1f, 0xc0, 0x53, 0x91, 0xd5, 0x17,
	0x70, 0x8b, 0x1e, 0x05, 0xb7, 0xa8, 0xe7, 0x58, 0x87, 0x1a, 0x36, 0xcd, 0xd4, 0x99, 0x89, 0xe0,
	0x93, 0xf8, 0x0e, 0x3e, 0xa5, 0xec, 0x66, 0x57, 0xd6, 0x92, 0x9b, 0xc7, 0xf9, 0xbf, 0xe1, 0x23,
	0x7f, 0x46, 0xcd, 0x18, 0xe8, 0xc3, 0x75, 0xc0, 0xf5, 0x48, 0x28, 0xa8, 0x8f, 0x7a, 0x1a, 0xbb,
	0x5a, 0x80, 0xc5, 0x85, 0xde, 0xcc, 0x06, 0x60, 0xb6, 0xfd, 0x8e, 0x9a, 0xaa, 0xc3, 0x20, 0x84,
	0x3e, 0x8d, 0xf3, 0xef, 0x42, 0x1d, 0x2f, 0x20, 0x74, 0x6f, 0x83, 0xa5, 0x55, 0x9b, 0x44, 0xfa,
	0x4e, 0x95, 0x8f, 0xc1, 0xd2, 0x67, 0x63, 0xbd, 0xd7, 0x67, 0xf5, 0xbe, 0xaf, 0x6e, 0xdd, 0x30,
	0x7a, 0x58, 0xc2, 0x7b, 0x04, 0x16, 0x73, 0x9e, 0x87, 0x3c, 0x62, 0x60, 0xd0, 0xf7, 0xaa, 0x6a,
	0x85, 0xc0, 0x0e, 0x2e, 0xf4, 0xff, 0x74, 0x5d, 0x14, 0x97, 0xc5, 0xfc, 0xeb, 0x40, 0x55, 0xcf,
	0x48, 0x2b, 0xa0, 0xdd, 0x4b, 0x6f, 0x55, 0xb9, 0x8c, 0x61, 0x3d, 0x01, 0xe9, 0x93, 0x89, 0x60,
	0x93, 0xde, 0x50, 0xcf, 0xc6, 0xe4, 0x48, 0x2b, 0x56, 0x22, 0xaf, 0xc5, 0x5b, 0x4d, 0xe3, 0x1d,
	0x04, 0x99, 0x6a, 0x52, 0x9a, 0xd3, 0x24, 0xb2, 0xa7, 0x59, 0xa8, 0xb2, 0x41, 0x82, 0x06, 0x63,
	0x10, 0x7d, 0x3a, 0x59, 0x46, 0xfa, 0x6d, 0x6a, 0x72, 0x68, 0xfb, 0x67, 0xd7, 0x4a, 0x3d, 0x44,
	0x27, 0xa9, 0xa6, 0xd6, 0x7f, 0x37, 0x9f, 0xd0, 0xbd, 0x9a, 0x4c, 0xf6, 0x72, 0xb8, 0xb9, 0xe6,
	0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x84, 0x02, 0xe3, 0x0c, 0x02, 0x00, 0x00,
}
