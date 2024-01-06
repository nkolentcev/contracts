// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: exchange/exchange.proto

package exchangev1

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
	Exchange_Register_FullMethodName      = "/exchange.Exchange/Register"
	Exchange_Login_FullMethodName         = "/exchange.Exchange/Login"
	Exchange_UpdateData_FullMethodName    = "/exchange.Exchange/UpdateData"
	Exchange_ReciveCommand_FullMethodName = "/exchange.Exchange/ReciveCommand"
	Exchange_CuntPass_FullMethodName      = "/exchange.Exchange/CuntPass"
	Exchange_CuntDelayed_FullMethodName   = "/exchange.Exchange/CuntDelayed"
	Exchange_CuntOpening_FullMethodName   = "/exchange.Exchange/CuntOpening"
)

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error)
	ReciveCommand(ctx context.Context, in *ResiveCommandRequest, opts ...grpc.CallOption) (*ResiveCommandsResponse, error)
	CuntPass(ctx context.Context, in *CountPassRequest, opts ...grpc.CallOption) (*CountPassResponse, error)
	CuntDelayed(ctx context.Context, in *CuntDelayedRequest, opts ...grpc.CallOption) (*CuntDelayedResponse, error)
	CuntOpening(ctx context.Context, in *CuntOpeningRequest, opts ...grpc.CallOption) (*CuntOpeningResponse, error)
}

type exchangeClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeClient(cc grpc.ClientConnInterface) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Exchange_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Exchange_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error) {
	out := new(UpdateDataResponse)
	err := c.cc.Invoke(ctx, Exchange_UpdateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) ReciveCommand(ctx context.Context, in *ResiveCommandRequest, opts ...grpc.CallOption) (*ResiveCommandsResponse, error) {
	out := new(ResiveCommandsResponse)
	err := c.cc.Invoke(ctx, Exchange_ReciveCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) CuntPass(ctx context.Context, in *CountPassRequest, opts ...grpc.CallOption) (*CountPassResponse, error) {
	out := new(CountPassResponse)
	err := c.cc.Invoke(ctx, Exchange_CuntPass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) CuntDelayed(ctx context.Context, in *CuntDelayedRequest, opts ...grpc.CallOption) (*CuntDelayedResponse, error) {
	out := new(CuntDelayedResponse)
	err := c.cc.Invoke(ctx, Exchange_CuntDelayed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) CuntOpening(ctx context.Context, in *CuntOpeningRequest, opts ...grpc.CallOption) (*CuntOpeningResponse, error) {
	out := new(CuntOpeningResponse)
	err := c.cc.Invoke(ctx, Exchange_CuntOpening_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
// All implementations must embed UnimplementedExchangeServer
// for forward compatibility
type ExchangeServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error)
	ReciveCommand(context.Context, *ResiveCommandRequest) (*ResiveCommandsResponse, error)
	CuntPass(context.Context, *CountPassRequest) (*CountPassResponse, error)
	CuntDelayed(context.Context, *CuntDelayedRequest) (*CuntDelayedResponse, error)
	CuntOpening(context.Context, *CuntOpeningRequest) (*CuntOpeningResponse, error)
	mustEmbedUnimplementedExchangeServer()
}

// UnimplementedExchangeServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (UnimplementedExchangeServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedExchangeServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedExchangeServer) UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedExchangeServer) ReciveCommand(context.Context, *ResiveCommandRequest) (*ResiveCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReciveCommand not implemented")
}
func (UnimplementedExchangeServer) CuntPass(context.Context, *CountPassRequest) (*CountPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CuntPass not implemented")
}
func (UnimplementedExchangeServer) CuntDelayed(context.Context, *CuntDelayedRequest) (*CuntDelayedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CuntDelayed not implemented")
}
func (UnimplementedExchangeServer) CuntOpening(context.Context, *CuntOpeningRequest) (*CuntOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CuntOpening not implemented")
}
func (UnimplementedExchangeServer) mustEmbedUnimplementedExchangeServer() {}

// UnsafeExchangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServer will
// result in compilation errors.
type UnsafeExchangeServer interface {
	mustEmbedUnimplementedExchangeServer()
}

func RegisterExchangeServer(s grpc.ServiceRegistrar, srv ExchangeServer) {
	s.RegisterService(&Exchange_ServiceDesc, srv)
}

func _Exchange_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_UpdateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).UpdateData(ctx, req.(*UpdateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_ReciveCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResiveCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).ReciveCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_ReciveCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).ReciveCommand(ctx, req.(*ResiveCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_CuntPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).CuntPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_CuntPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).CuntPass(ctx, req.(*CountPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_CuntDelayed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CuntDelayedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).CuntDelayed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_CuntDelayed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).CuntDelayed(ctx, req.(*CuntDelayedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_CuntOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CuntOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).CuntOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_CuntOpening_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).CuntOpening(ctx, req.(*CuntOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Exchange_ServiceDesc is the grpc.ServiceDesc for Exchange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exchange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Exchange_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Exchange_Login_Handler,
		},
		{
			MethodName: "UpdateData",
			Handler:    _Exchange_UpdateData_Handler,
		},
		{
			MethodName: "ReciveCommand",
			Handler:    _Exchange_ReciveCommand_Handler,
		},
		{
			MethodName: "CuntPass",
			Handler:    _Exchange_CuntPass_Handler,
		},
		{
			MethodName: "CuntDelayed",
			Handler:    _Exchange_CuntDelayed_Handler,
		},
		{
			MethodName: "CuntOpening",
			Handler:    _Exchange_CuntOpening_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange/exchange.proto",
}
