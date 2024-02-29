// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/routeRaydar.proto

package routeRaydar

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

// RouteServiceClient is the client API for RouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteServiceClient interface {
	// GetRoute returns the details of a specific route.
	GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error)
	SubmitGrid(ctx context.Context, in *SubmitGridRequest, opts ...grpc.CallOption) (*SubmitGridResponse, error)
	SendCoordinates(ctx context.Context, in *SendCoordinatesRequest, opts ...grpc.CallOption) (*SendCoordinatesResponse, error)
	StreamRide(ctx context.Context, in *StreamRideRequest, opts ...grpc.CallOption) (RouteService_StreamRideClient, error)
}

type routeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteServiceClient(cc grpc.ClientConnInterface) RouteServiceClient {
	return &routeServiceClient{cc}
}

func (c *routeServiceClient) GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error) {
	out := new(GetRouteResponse)
	err := c.cc.Invoke(ctx, "/routeRaydar.RouteService/GetRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) SubmitGrid(ctx context.Context, in *SubmitGridRequest, opts ...grpc.CallOption) (*SubmitGridResponse, error) {
	out := new(SubmitGridResponse)
	err := c.cc.Invoke(ctx, "/routeRaydar.RouteService/SubmitGrid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) SendCoordinates(ctx context.Context, in *SendCoordinatesRequest, opts ...grpc.CallOption) (*SendCoordinatesResponse, error) {
	out := new(SendCoordinatesResponse)
	err := c.cc.Invoke(ctx, "/routeRaydar.RouteService/SendCoordinates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) StreamRide(ctx context.Context, in *StreamRideRequest, opts ...grpc.CallOption) (RouteService_StreamRideClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteService_ServiceDesc.Streams[0], "/routeRaydar.RouteService/StreamRide", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeServiceStreamRideClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteService_StreamRideClient interface {
	Recv() (*Coordinates, error)
	grpc.ClientStream
}

type routeServiceStreamRideClient struct {
	grpc.ClientStream
}

func (x *routeServiceStreamRideClient) Recv() (*Coordinates, error) {
	m := new(Coordinates)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteServiceServer is the server API for RouteService service.
// All implementations must embed UnimplementedRouteServiceServer
// for forward compatibility
type RouteServiceServer interface {
	// GetRoute returns the details of a specific route.
	GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error)
	SubmitGrid(context.Context, *SubmitGridRequest) (*SubmitGridResponse, error)
	SendCoordinates(context.Context, *SendCoordinatesRequest) (*SendCoordinatesResponse, error)
	StreamRide(*StreamRideRequest, RouteService_StreamRideServer) error
	mustEmbedUnimplementedRouteServiceServer()
}

// UnimplementedRouteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServiceServer struct {
}

func (UnimplementedRouteServiceServer) GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}
func (UnimplementedRouteServiceServer) SubmitGrid(context.Context, *SubmitGridRequest) (*SubmitGridResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGrid not implemented")
}
func (UnimplementedRouteServiceServer) SendCoordinates(context.Context, *SendCoordinatesRequest) (*SendCoordinatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCoordinates not implemented")
}
func (UnimplementedRouteServiceServer) StreamRide(*StreamRideRequest, RouteService_StreamRideServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRide not implemented")
}
func (UnimplementedRouteServiceServer) mustEmbedUnimplementedRouteServiceServer() {}

// UnsafeRouteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServiceServer will
// result in compilation errors.
type UnsafeRouteServiceServer interface {
	mustEmbedUnimplementedRouteServiceServer()
}

func RegisterRouteServiceServer(s grpc.ServiceRegistrar, srv RouteServiceServer) {
	s.RegisterService(&RouteService_ServiceDesc, srv)
}

func _RouteService_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeRaydar.RouteService/GetRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).GetRoute(ctx, req.(*GetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_SubmitGrid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitGridRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).SubmitGrid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeRaydar.RouteService/SubmitGrid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).SubmitGrid(ctx, req.(*SubmitGridRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_SendCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCoordinatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).SendCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeRaydar.RouteService/SendCoordinates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).SendCoordinates(ctx, req.(*SendCoordinatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_StreamRide_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRideRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteServiceServer).StreamRide(m, &routeServiceStreamRideServer{stream})
}

type RouteService_StreamRideServer interface {
	Send(*Coordinates) error
	grpc.ServerStream
}

type routeServiceStreamRideServer struct {
	grpc.ServerStream
}

func (x *routeServiceStreamRideServer) Send(m *Coordinates) error {
	return x.ServerStream.SendMsg(m)
}

// RouteService_ServiceDesc is the grpc.ServiceDesc for RouteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeRaydar.RouteService",
	HandlerType: (*RouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoute",
			Handler:    _RouteService_GetRoute_Handler,
		},
		{
			MethodName: "SubmitGrid",
			Handler:    _RouteService_SubmitGrid_Handler,
		},
		{
			MethodName: "SendCoordinates",
			Handler:    _RouteService_SendCoordinates_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRide",
			Handler:       _RouteService_StreamRide_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/routeRaydar.proto",
}
