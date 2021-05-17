// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tx

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)
	// BroadcastTx broadcast transaction.
	BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error)
	// GetTxsEvent fetches txs by event.
	GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error)
}

type serviceClient struct {
	cc           grpc.ClientConnInterface
	_Simulate    types.Invoker
	_GetTx       types.Invoker
	_BroadcastTx types.Invoker
	_GetTxsEvent types.Invoker
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc: cc}
}

func (c *serviceClient) Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error) {
	if invoker := c._Simulate; invoker != nil {
		var out SimulateResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Simulate, err = invokerConn.Invoker("/cosmos.tx.v1beta1.ServiceSimulate")
		if err != nil {
			var out SimulateResponse
			err = c._Simulate(ctx, in, &out)
			return &out, err
		}
	}
	out := new(SimulateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.ServiceSimulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTx(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error) {
	if invoker := c._GetTx; invoker != nil {
		var out GetTxResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetTx, err = invokerConn.Invoker("/cosmos.tx.v1beta1.ServiceGetTx")
		if err != nil {
			var out GetTxResponse
			err = c._GetTx(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.ServiceGetTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BroadcastTx(ctx context.Context, in *BroadcastTxRequest, opts ...grpc.CallOption) (*BroadcastTxResponse, error) {
	if invoker := c._BroadcastTx; invoker != nil {
		var out BroadcastTxResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._BroadcastTx, err = invokerConn.Invoker("/cosmos.tx.v1beta1.ServiceBroadcastTx")
		if err != nil {
			var out BroadcastTxResponse
			err = c._BroadcastTx(ctx, in, &out)
			return &out, err
		}
	}
	out := new(BroadcastTxResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.ServiceBroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTxsEvent(ctx context.Context, in *GetTxsEventRequest, opts ...grpc.CallOption) (*GetTxsEventResponse, error) {
	if invoker := c._GetTxsEvent; invoker != nil {
		var out GetTxsEventResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._GetTxsEvent, err = invokerConn.Invoker("/cosmos.tx.v1beta1.ServiceGetTxsEvent")
		if err != nil {
			var out GetTxsEventResponse
			err = c._GetTxsEvent(ctx, in, &out)
			return &out, err
		}
	}
	out := new(GetTxsEventResponse)
	err := c.cc.Invoke(ctx, "/cosmos.tx.v1beta1.ServiceGetTxsEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)
	// GetTx fetches a tx by hash.
	GetTx(context.Context, *GetTxRequest) (*GetTxResponse, error)
	// BroadcastTx broadcast transaction.
	BroadcastTx(context.Context, *BroadcastTxRequest) (*BroadcastTxResponse, error)
	// GetTxsEvent fetches txs by event.
	GetTxsEvent(context.Context, *GetTxsEventRequest) (*GetTxsEventResponse, error)
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Simulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Simulate(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.ServiceSimulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Simulate(types.UnwrapSDKContext(ctx), req.(*SimulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTx(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.ServiceGetTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTx(types.UnwrapSDKContext(ctx), req.(*GetTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BroadcastTx(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.ServiceBroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BroadcastTx(types.UnwrapSDKContext(ctx), req.(*BroadcastTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTxsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTxsEvent(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.tx.v1beta1.ServiceGetTxsEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTxsEvent(types.UnwrapSDKContext(ctx), req.(*GetTxsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.tx.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simulate",
			Handler:    _Service_Simulate_Handler,
		},
		{
			MethodName: "GetTx",
			Handler:    _Service_GetTx_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _Service_BroadcastTx_Handler,
		},
		{
			MethodName: "GetTxsEvent",
			Handler:    _Service_GetTxsEvent_Handler,
		},
	},
	Metadata: "cosmos/tx/v1beta1/service.proto",
}

const (
	ServiceSimulateMethod    = "/cosmos.tx.v1beta1.ServiceSimulate"
	ServiceGetTxMethod       = "/cosmos.tx.v1beta1.ServiceGetTx"
	ServiceBroadcastTxMethod = "/cosmos.tx.v1beta1.ServiceBroadcastTx"
	ServiceGetTxsEventMethod = "/cosmos.tx.v1beta1.ServiceGetTxsEvent"
)
