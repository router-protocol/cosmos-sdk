// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v2alpha1

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReflectionServiceClient interface {
	// GetAuthnDescriptor returns information on how to authenticate transactions in the application
	// NOTE: this RPC is still experimental and might be subject to breaking changes or removal in
	// future releases of the cosmos-sdk.
	GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error)
	// GetChainDescriptor returns the description of the chain
	GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error)
	// GetCodecDescriptor returns the descriptor of the codec of the application
	GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error)
	// GetConfigurationDescriptor returns the descriptor for the sdk.Config of the application
	GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error)
	// GetQueryServicesDescriptor returns the available gRPC queryable services of the application
	GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error)
	// GetTxDescriptor returns information on the used transaction object and available msgs that can be used
	GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error)
}

type reflectionServiceClient struct {
	cc                          grpc.ClientConnInterface
	_GetAuthnDescriptor         types.Invoker
	_GetChainDescriptor         types.Invoker
	_GetCodecDescriptor         types.Invoker
	_GetConfigurationDescriptor types.Invoker
	_GetQueryServicesDescriptor types.Invoker
	_GetTxDescriptor            types.Invoker
}

func NewReflectionServiceClient(cc grpc.ClientConnInterface) ReflectionServiceClient {
	return &reflectionServiceClient{cc: cc}
}

func (c *reflectionServiceClient) GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error) {
	if invoker := c._GetAuthnDescriptor; invoker != nil {
		var out GetAuthnDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetAuthnDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetAuthnDescriptor")
		if err != nil {
			var out GetAuthnDescriptorResponse
			err = c._GetAuthnDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetAuthnDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetAuthnDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error) {
	if invoker := c._GetChainDescriptor; invoker != nil {
		var out GetChainDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetChainDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetChainDescriptor")
		if err != nil {
			var out GetChainDescriptorResponse
			err = c._GetChainDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetChainDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetChainDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error) {
	if invoker := c._GetCodecDescriptor; invoker != nil {
		var out GetCodecDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetCodecDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetCodecDescriptor")
		if err != nil {
			var out GetCodecDescriptorResponse
			err = c._GetCodecDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetCodecDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetCodecDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error) {
	if invoker := c._GetConfigurationDescriptor; invoker != nil {
		var out GetConfigurationDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetConfigurationDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetConfigurationDescriptor")
		if err != nil {
			var out GetConfigurationDescriptorResponse
			err = c._GetConfigurationDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetConfigurationDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetConfigurationDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error) {
	if invoker := c._GetQueryServicesDescriptor; invoker != nil {
		var out GetQueryServicesDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetQueryServicesDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetQueryServicesDescriptor")
		if err != nil {
			var out GetQueryServicesDescriptorResponse
			err = c._GetQueryServicesDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetQueryServicesDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetQueryServicesDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error) {
	if invoker := c._GetTxDescriptor; invoker != nil {
		var out GetTxDescriptorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetTxDescriptor, err = invokerConn.Invoker("/cosmos.base.reflection.v2alpha1.ReflectionServiceGetTxDescriptor")
		if err != nil {
			var out GetTxDescriptorResponse
			err = c._GetTxDescriptor(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetTxDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetTxDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
type ReflectionServiceServer interface {
	// GetAuthnDescriptor returns information on how to authenticate transactions in the application
	// NOTE: this RPC is still experimental and might be subject to breaking changes or removal in
	// future releases of the cosmos-sdk.
	GetAuthnDescriptor(context.Context, *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error)
	// GetChainDescriptor returns the description of the chain
	GetChainDescriptor(context.Context, *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error)
	// GetCodecDescriptor returns the descriptor of the codec of the application
	GetCodecDescriptor(context.Context, *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error)
	// GetConfigurationDescriptor returns the descriptor for the sdk.Config of the application
	GetConfigurationDescriptor(context.Context, *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error)
	// GetQueryServicesDescriptor returns the available gRPC queryable services of the application
	GetQueryServicesDescriptor(context.Context, *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error)
	// GetTxDescriptor returns information on the used transaction object and available msgs that can be used
	GetTxDescriptor(context.Context, *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error)
}

func RegisterReflectionServiceServer(s grpc.ServiceRegistrar, srv ReflectionServiceServer) {
	s.RegisterService(&ReflectionService_ServiceDesc, srv)
}

func _ReflectionService_GetAuthnDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthnDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetAuthnDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(types.UnwrapSDKContext(ctx), req.(*GetAuthnDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetChainDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetChainDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetChainDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetChainDescriptor(types.UnwrapSDKContext(ctx), req.(*GetChainDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetCodecDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodecDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetCodecDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(types.UnwrapSDKContext(ctx), req.(*GetCodecDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetConfigurationDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetConfigurationDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(types.UnwrapSDKContext(ctx), req.(*GetConfigurationDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetQueryServicesDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryServicesDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetQueryServicesDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(types.UnwrapSDKContext(ctx), req.(*GetQueryServicesDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetTxDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetTxDescriptor(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetTxDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetTxDescriptor(types.UnwrapSDKContext(ctx), req.(*GetTxDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflectionService_ServiceDesc is the grpc.ServiceDesc for ReflectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v2alpha1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthnDescriptor",
			Handler:    _ReflectionService_GetAuthnDescriptor_Handler,
		},
		{
			MethodName: "GetChainDescriptor",
			Handler:    _ReflectionService_GetChainDescriptor_Handler,
		},
		{
			MethodName: "GetCodecDescriptor",
			Handler:    _ReflectionService_GetCodecDescriptor_Handler,
		},
		{
			MethodName: "GetConfigurationDescriptor",
			Handler:    _ReflectionService_GetConfigurationDescriptor_Handler,
		},
		{
			MethodName: "GetQueryServicesDescriptor",
			Handler:    _ReflectionService_GetQueryServicesDescriptor_Handler,
		},
		{
			MethodName: "GetTxDescriptor",
			Handler:    _ReflectionService_GetTxDescriptor_Handler,
		},
	},
	Metadata: "cosmos/base/reflection/v2alpha1/reflection.proto",
}

const (
	ReflectionServiceGetAuthnDescriptorMethod         = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetAuthnDescriptor"
	ReflectionServiceGetChainDescriptorMethod         = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetChainDescriptor"
	ReflectionServiceGetCodecDescriptorMethod         = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetCodecDescriptor"
	ReflectionServiceGetConfigurationDescriptorMethod = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetConfigurationDescriptor"
	ReflectionServiceGetQueryServicesDescriptorMethod = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetQueryServicesDescriptor"
	ReflectionServiceGetTxDescriptorMethod            = "/cosmos.base.reflection.v2alpha1.ReflectionServiceGetTxDescriptor"
)
