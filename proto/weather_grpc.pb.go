// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: weather.proto

package WeatherFetcher

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	WeatherFetcher_GetWeather_FullMethodName = "/weather.WeatherFetcher/GetWeather"
)

// WeatherFetcherClient is the client API for WeatherFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The WeatherFetcher service definition
type WeatherFetcherClient interface {
	GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
}

type weatherFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherFetcherClient(cc grpc.ClientConnInterface) WeatherFetcherClient {
	return &weatherFetcherClient{cc}
}

func (c *weatherFetcherClient) GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, WeatherFetcher_GetWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherFetcherServer is the server API for WeatherFetcher service.
// All implementations must embed UnimplementedWeatherFetcherServer
// for forward compatibility
//
// The WeatherFetcher service definition
type WeatherFetcherServer interface {
	GetWeather(context.Context, *WeatherRequest) (*WeatherResponse, error)
	mustEmbedUnimplementedWeatherFetcherServer()
}

// UnimplementedWeatherFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherFetcherServer struct {
}

func (UnimplementedWeatherFetcherServer) GetWeather(context.Context, *WeatherRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}
func (UnimplementedWeatherFetcherServer) mustEmbedUnimplementedWeatherFetcherServer() {}

// UnsafeWeatherFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherFetcherServer will
// result in compilation errors.
type UnsafeWeatherFetcherServer interface {
	mustEmbedUnimplementedWeatherFetcherServer()
}

func RegisterWeatherFetcherServer(s grpc.ServiceRegistrar, srv WeatherFetcherServer) {
	s.RegisterService(&WeatherFetcher_ServiceDesc, srv)
}

func _WeatherFetcher_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherFetcherServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherFetcher_GetWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherFetcherServer).GetWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherFetcher_ServiceDesc is the grpc.ServiceDesc for WeatherFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherFetcher",
	HandlerType: (*WeatherFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _WeatherFetcher_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}