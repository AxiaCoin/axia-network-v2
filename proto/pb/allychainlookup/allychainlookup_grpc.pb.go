// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: allychainlookup/allychainlookup.proto

package allychainlookup

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

// AllychainLookupClient is the client API for AllychainLookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllychainLookupClient interface {
	AllychainID(ctx context.Context, in *AllychainIDRequest, opts ...grpc.CallOption) (*AllychainIDResponse, error)
}

type allychainLookupClient struct {
	cc grpc.ClientConnInterface
}

func NewAllychainLookupClient(cc grpc.ClientConnInterface) AllychainLookupClient {
	return &allychainLookupClient{cc}
}

func (c *allychainLookupClient) AllychainID(ctx context.Context, in *AllychainIDRequest, opts ...grpc.CallOption) (*AllychainIDResponse, error) {
	out := new(AllychainIDResponse)
	err := c.cc.Invoke(ctx, "/allychainlookup.AllychainLookup/AllychainID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllychainLookupServer is the server API for AllychainLookup service.
// All implementations must embed UnimplementedAllychainLookupServer
// for forward compatibility
type AllychainLookupServer interface {
	AllychainID(context.Context, *AllychainIDRequest) (*AllychainIDResponse, error)
	mustEmbedUnimplementedAllychainLookupServer()
}

// UnimplementedAllychainLookupServer must be embedded to have forward compatible implementations.
type UnimplementedAllychainLookupServer struct {
}

func (UnimplementedAllychainLookupServer) AllychainID(context.Context, *AllychainIDRequest) (*AllychainIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllychainID not implemented")
}
func (UnimplementedAllychainLookupServer) mustEmbedUnimplementedAllychainLookupServer() {}

// UnsafeAllychainLookupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllychainLookupServer will
// result in compilation errors.
type UnsafeAllychainLookupServer interface {
	mustEmbedUnimplementedAllychainLookupServer()
}

func RegisterAllychainLookupServer(s grpc.ServiceRegistrar, srv AllychainLookupServer) {
	s.RegisterService(&AllychainLookup_ServiceDesc, srv)
}

func _AllychainLookup_AllychainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllychainIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllychainLookupServer).AllychainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allychainlookup.AllychainLookup/AllychainID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllychainLookupServer).AllychainID(ctx, req.(*AllychainIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AllychainLookup_ServiceDesc is the grpc.ServiceDesc for AllychainLookup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AllychainLookup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "allychainlookup.AllychainLookup",
	HandlerType: (*AllychainLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllychainID",
			Handler:    _AllychainLookup_AllychainID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allychainlookup/allychainlookup.proto",
}
