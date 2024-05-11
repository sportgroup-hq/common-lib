// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/api.proto

package api

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error)
	UserExistsByEmail(ctx context.Context, in *UserExistsByEmailRequest, opts ...grpc.CallOption) (*UserExistsByEmailResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error) {
	out := new(GetUserByIDResponse)
	err := c.cc.Invoke(ctx, "/Api/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error) {
	out := new(GetUserByEmailResponse)
	err := c.cc.Invoke(ctx, "/Api/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UserExistsByEmail(ctx context.Context, in *UserExistsByEmailRequest, opts ...grpc.CallOption) (*UserExistsByEmailResponse, error) {
	out := new(UserExistsByEmailResponse)
	err := c.cc.Invoke(ctx, "/Api/UserExistsByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Api/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	UserExistsByEmail(context.Context, *UserExistsByEmailRequest) (*UserExistsByEmailResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedApiServer) GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedApiServer) UserExistsByEmail(context.Context, *UserExistsByEmailRequest) (*UserExistsByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserExistsByEmail not implemented")
}
func (UnimplementedApiServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetUserByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetUserByEmail(ctx, req.(*GetUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UserExistsByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserExistsByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UserExistsByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/UserExistsByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UserExistsByEmail(ctx, req.(*UserExistsByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByID",
			Handler:    _Api_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _Api_GetUserByEmail_Handler,
		},
		{
			MethodName: "UserExistsByEmail",
			Handler:    _Api_UserExistsByEmail_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Api_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
