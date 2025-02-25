// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package netmap

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

// NetmapServiceClient is the client API for NetmapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetmapServiceClient interface {
	// Get NodeInfo structure from the particular node directly. Node information
	// can be taken from `Netmap` smart contract, but in some cases the one may
	// want to get recent information directly, or to talk to the node not yet
	// present in `Network Map` to find out what API version can be used for
	// further communication. Can also be used to check if node is up and running.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// information about the server has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	LocalNodeInfo(ctx context.Context, in *LocalNodeInfoRequest, opts ...grpc.CallOption) (*LocalNodeInfoResponse, error)
	// Read recent information about the NeoFS network.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// information about the current network state has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	NetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
}

type netmapServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetmapServiceClient(cc grpc.ClientConnInterface) NetmapServiceClient {
	return &netmapServiceClient{cc}
}

func (c *netmapServiceClient) LocalNodeInfo(ctx context.Context, in *LocalNodeInfoRequest, opts ...grpc.CallOption) (*LocalNodeInfoResponse, error) {
	out := new(LocalNodeInfoResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.netmap.NetmapService/LocalNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmapServiceClient) NetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.netmap.NetmapService/NetworkInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetmapServiceServer is the server API for NetmapService service.
// All implementations should embed UnimplementedNetmapServiceServer
// for forward compatibility
type NetmapServiceServer interface {
	// Get NodeInfo structure from the particular node directly. Node information
	// can be taken from `Netmap` smart contract, but in some cases the one may
	// want to get recent information directly, or to talk to the node not yet
	// present in `Network Map` to find out what API version can be used for
	// further communication. Can also be used to check if node is up and running.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// information about the server has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	LocalNodeInfo(context.Context, *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error)
	// Read recent information about the NeoFS network.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// information about the current network state has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	NetworkInfo(context.Context, *NetworkInfoRequest) (*NetworkInfoResponse, error)
}

// UnimplementedNetmapServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNetmapServiceServer struct {
}

func (UnimplementedNetmapServiceServer) LocalNodeInfo(context.Context, *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocalNodeInfo not implemented")
}
func (UnimplementedNetmapServiceServer) NetworkInfo(context.Context, *NetworkInfoRequest) (*NetworkInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkInfo not implemented")
}

// UnsafeNetmapServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetmapServiceServer will
// result in compilation errors.
type UnsafeNetmapServiceServer interface {
	mustEmbedUnimplementedNetmapServiceServer()
}

func RegisterNetmapServiceServer(s grpc.ServiceRegistrar, srv NetmapServiceServer) {
	s.RegisterService(&NetmapService_ServiceDesc, srv)
}

func _NetmapService_LocalNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmapServiceServer).LocalNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.netmap.NetmapService/LocalNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmapServiceServer).LocalNodeInfo(ctx, req.(*LocalNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetmapService_NetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmapServiceServer).NetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.netmap.NetmapService/NetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmapServiceServer).NetworkInfo(ctx, req.(*NetworkInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetmapService_ServiceDesc is the grpc.ServiceDesc for NetmapService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetmapService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neo.fs.v2.netmap.NetmapService",
	HandlerType: (*NetmapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocalNodeInfo",
			Handler:    _NetmapService_LocalNodeInfo_Handler,
		},
		{
			MethodName: "NetworkInfo",
			Handler:    _NetmapService_NetworkInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netmap/grpc/service.proto",
}
