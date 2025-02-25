// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accounting

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

// AccountingServiceClient is the client API for AccountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingServiceClient interface {
	// Returns the amount of funds in GAS token for the requested NeoFS account.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// balance has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type accountingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingServiceClient(cc grpc.ClientConnInterface) AccountingServiceClient {
	return &accountingServiceClient{cc}
}

func (c *accountingServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.accounting.AccountingService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServiceServer is the server API for AccountingService service.
// All implementations should embed UnimplementedAccountingServiceServer
// for forward compatibility
type AccountingServiceServer interface {
	// Returns the amount of funds in GAS token for the requested NeoFS account.
	//
	// Statuses:
	// - **OK** (0, SECTION_SUCCESS):
	// balance has been successfully read;
	// - Common failures (SECTION_FAILURE_COMMON).
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

// UnimplementedAccountingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountingServiceServer struct {
}

func (UnimplementedAccountingServiceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

// UnsafeAccountingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingServiceServer will
// result in compilation errors.
type UnsafeAccountingServiceServer interface {
	mustEmbedUnimplementedAccountingServiceServer()
}

func RegisterAccountingServiceServer(s grpc.ServiceRegistrar, srv AccountingServiceServer) {
	s.RegisterService(&AccountingService_ServiceDesc, srv)
}

func _AccountingService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.accounting.AccountingService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountingService_ServiceDesc is the grpc.ServiceDesc for AccountingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neo.fs.v2.accounting.AccountingService",
	HandlerType: (*AccountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _AccountingService_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounting/grpc/service.proto",
}
