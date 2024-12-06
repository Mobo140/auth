// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.3
// source: auth.proto

package auth_v1

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

// AuthV1Client is the client API for AuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthV1Client interface {
	// Get refresh token
	GetRefreshToken(ctx context.Context, in *GetRefreshTokenRequest, opts ...grpc.CallOption) (*GetRefreshTokenResponse, error)
	// Get access token
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
	// Login
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthV1Client(cc grpc.ClientConnInterface) AuthV1Client {
	return &authV1Client{cc}
}

func (c *authV1Client) GetRefreshToken(ctx context.Context, in *GetRefreshTokenRequest, opts ...grpc.CallOption) (*GetRefreshTokenResponse, error) {
	out := new(GetRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/GetRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authV1Client) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthV1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthV1Server is the server API for AuthV1 service.
// All implementations must embed UnimplementedAuthV1Server
// for forward compatibility
type AuthV1Server interface {
	// Get refresh token
	GetRefreshToken(context.Context, *GetRefreshTokenRequest) (*GetRefreshTokenResponse, error)
	// Get access token
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	// Login
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthV1Server()
}

// UnimplementedAuthV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthV1Server struct {
}

func (UnimplementedAuthV1Server) GetRefreshToken(context.Context, *GetRefreshTokenRequest) (*GetRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshToken not implemented")
}
func (UnimplementedAuthV1Server) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthV1Server) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthV1Server) mustEmbedUnimplementedAuthV1Server() {}

// UnsafeAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthV1Server will
// result in compilation errors.
type UnsafeAuthV1Server interface {
	mustEmbedUnimplementedAuthV1Server()
}

func RegisterAuthV1Server(s grpc.ServiceRegistrar, srv AuthV1Server) {
	s.RegisterService(&AuthV1_ServiceDesc, srv)
}

func _AuthV1_GetRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).GetRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/GetRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).GetRefreshToken(ctx, req.(*GetRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthV1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthV1Server).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthV1_ServiceDesc is the grpc.ServiceDesc for AuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.AuthV1",
	HandlerType: (*AuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRefreshToken",
			Handler:    _AuthV1_GetRefreshToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _AuthV1_GetAccessToken_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthV1_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
