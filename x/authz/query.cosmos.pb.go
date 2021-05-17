// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authz

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error)
}

type queryClient struct {
	cc      grpc.ClientConnInterface
	_Grants types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error) {
	if invoker := c._Grants; invoker != nil {
		var out QueryGrantsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Grants, err = invokerConn.Invoker("/cosmos.authz.v1beta1.QueryGrants")
		if err != nil {
			var out QueryGrantsResponse
			err = c._Grants(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.QueryGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Grants(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.QueryGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Grants(types.UnwrapSDKContext(ctx), req.(*QueryGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _Query_Grants_Handler,
		},
	},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}

const (
	QueryGrantsMethod = "/cosmos.authz.v1beta1.QueryGrants"
)
