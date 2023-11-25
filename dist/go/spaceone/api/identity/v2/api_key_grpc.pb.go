// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/api_key.proto

package v2

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	APIKey_Create_FullMethodName  = "/spaceone.api.identity.v2.APIKey/create"
	APIKey_Update_FullMethodName  = "/spaceone.api.identity.v2.APIKey/update"
	APIKey_Enable_FullMethodName  = "/spaceone.api.identity.v2.APIKey/enable"
	APIKey_Disable_FullMethodName = "/spaceone.api.identity.v2.APIKey/disable"
	APIKey_Delete_FullMethodName  = "/spaceone.api.identity.v2.APIKey/delete"
	APIKey_Verify_FullMethodName  = "/spaceone.api.identity.v2.APIKey/verify"
	APIKey_Get_FullMethodName     = "/spaceone.api.identity.v2.APIKey/get"
	APIKey_List_FullMethodName    = "/spaceone.api.identity.v2.APIKey/list"
	APIKey_Stat_FullMethodName    = "/spaceone.api.identity.v2.APIKey/stat"
)

// APIKeyClient is the client API for APIKey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIKeyClient interface {
	Create(ctx context.Context, in *CreateAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	Update(ctx context.Context, in *UpdateAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	Enable(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	Disable(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	Delete(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Verify(ctx context.Context, in *VerifyAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	Get(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error)
	List(ctx context.Context, in *APIKeySearchQuery, opts ...grpc.CallOption) (*APIKeysInfo, error)
	Stat(ctx context.Context, in *APIKeyStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type aPIKeyClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIKeyClient(cc grpc.ClientConnInterface) APIKeyClient {
	return &aPIKeyClient{cc}
}

func (c *aPIKeyClient) Create(ctx context.Context, in *CreateAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Update(ctx context.Context, in *UpdateAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Enable(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Disable(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Delete(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, APIKey_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Verify(ctx context.Context, in *VerifyAPIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Get(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*APIKeyInfo, error) {
	out := new(APIKeyInfo)
	err := c.cc.Invoke(ctx, APIKey_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) List(ctx context.Context, in *APIKeySearchQuery, opts ...grpc.CallOption) (*APIKeysInfo, error) {
	out := new(APIKeysInfo)
	err := c.cc.Invoke(ctx, APIKey_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyClient) Stat(ctx context.Context, in *APIKeyStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, APIKey_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIKeyServer is the server API for APIKey service.
// All implementations must embed UnimplementedAPIKeyServer
// for forward compatibility
type APIKeyServer interface {
	Create(context.Context, *CreateAPIKeyRequest) (*APIKeyInfo, error)
	Update(context.Context, *UpdateAPIKeyRequest) (*APIKeyInfo, error)
	Enable(context.Context, *APIKeyRequest) (*APIKeyInfo, error)
	Disable(context.Context, *APIKeyRequest) (*APIKeyInfo, error)
	Delete(context.Context, *APIKeyRequest) (*empty.Empty, error)
	Verify(context.Context, *VerifyAPIKeyRequest) (*APIKeyInfo, error)
	Get(context.Context, *APIKeyRequest) (*APIKeyInfo, error)
	List(context.Context, *APIKeySearchQuery) (*APIKeysInfo, error)
	Stat(context.Context, *APIKeyStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedAPIKeyServer()
}

// UnimplementedAPIKeyServer must be embedded to have forward compatible implementations.
type UnimplementedAPIKeyServer struct {
}

func (UnimplementedAPIKeyServer) Create(context.Context, *CreateAPIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAPIKeyServer) Update(context.Context, *UpdateAPIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAPIKeyServer) Enable(context.Context, *APIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedAPIKeyServer) Disable(context.Context, *APIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedAPIKeyServer) Delete(context.Context, *APIKeyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAPIKeyServer) Verify(context.Context, *VerifyAPIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAPIKeyServer) Get(context.Context, *APIKeyRequest) (*APIKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAPIKeyServer) List(context.Context, *APIKeySearchQuery) (*APIKeysInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAPIKeyServer) Stat(context.Context, *APIKeyStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedAPIKeyServer) mustEmbedUnimplementedAPIKeyServer() {}

// UnsafeAPIKeyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIKeyServer will
// result in compilation errors.
type UnsafeAPIKeyServer interface {
	mustEmbedUnimplementedAPIKeyServer()
}

func RegisterAPIKeyServer(s grpc.ServiceRegistrar, srv APIKeyServer) {
	s.RegisterService(&APIKey_ServiceDesc, srv)
}

func _APIKey_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Create(ctx, req.(*CreateAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Update(ctx, req.(*UpdateAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Enable(ctx, req.(*APIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Disable(ctx, req.(*APIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Delete(ctx, req.(*APIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Verify(ctx, req.(*VerifyAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Get(ctx, req.(*APIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeySearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).List(ctx, req.(*APIKeySearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKey_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKey_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServer).Stat(ctx, req.(*APIKeyStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// APIKey_ServiceDesc is the grpc.ServiceDesc for APIKey service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIKey_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.APIKey",
	HandlerType: (*APIKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _APIKey_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _APIKey_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _APIKey_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _APIKey_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _APIKey_Delete_Handler,
		},
		{
			MethodName: "verify",
			Handler:    _APIKey_Verify_Handler,
		},
		{
			MethodName: "get",
			Handler:    _APIKey_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _APIKey_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _APIKey_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/api_key.proto",
}
