// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/app.proto

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
	App_Create_FullMethodName         = "/spaceone.api.identity.v2.App/create"
	App_Update_FullMethodName         = "/spaceone.api.identity.v2.App/update"
	App_GenerateApiKey_FullMethodName = "/spaceone.api.identity.v2.App/generate_api_key"
	App_Enable_FullMethodName         = "/spaceone.api.identity.v2.App/enable"
	App_Disable_FullMethodName        = "/spaceone.api.identity.v2.App/disable"
	App_Delete_FullMethodName         = "/spaceone.api.identity.v2.App/delete"
	App_Get_FullMethodName            = "/spaceone.api.identity.v2.App/get"
	App_Check_FullMethodName          = "/spaceone.api.identity.v2.App/check"
	App_List_FullMethodName           = "/spaceone.api.identity.v2.App/list"
	App_Stat_FullMethodName           = "/spaceone.api.identity.v2.App/stat"
)

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	Create(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	Update(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	GenerateApiKey(ctx context.Context, in *GenerateAPIKeyAppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	Enable(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	Disable(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	Delete(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error)
	Check(ctx context.Context, in *AppCheckRequest, opts ...grpc.CallOption) (*CheckAppInfo, error)
	List(ctx context.Context, in *AppSearchQuery, opts ...grpc.CallOption) (*AppsInfo, error)
	Stat(ctx context.Context, in *AppStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Update(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GenerateApiKey(ctx context.Context, in *GenerateAPIKeyAppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_GenerateApiKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Enable(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Disable(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Delete(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, App_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Get(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppInfo, error) {
	out := new(AppInfo)
	err := c.cc.Invoke(ctx, App_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Check(ctx context.Context, in *AppCheckRequest, opts ...grpc.CallOption) (*CheckAppInfo, error) {
	out := new(CheckAppInfo)
	err := c.cc.Invoke(ctx, App_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) List(ctx context.Context, in *AppSearchQuery, opts ...grpc.CallOption) (*AppsInfo, error) {
	out := new(AppsInfo)
	err := c.cc.Invoke(ctx, App_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Stat(ctx context.Context, in *AppStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, App_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	Create(context.Context, *CreateAppRequest) (*AppInfo, error)
	Update(context.Context, *UpdateAppRequest) (*AppInfo, error)
	GenerateApiKey(context.Context, *GenerateAPIKeyAppRequest) (*AppInfo, error)
	Enable(context.Context, *AppRequest) (*AppInfo, error)
	Disable(context.Context, *AppRequest) (*AppInfo, error)
	Delete(context.Context, *AppRequest) (*empty.Empty, error)
	Get(context.Context, *AppRequest) (*AppInfo, error)
	Check(context.Context, *AppCheckRequest) (*CheckAppInfo, error)
	List(context.Context, *AppSearchQuery) (*AppsInfo, error)
	Stat(context.Context, *AppStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) Create(context.Context, *CreateAppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAppServer) Update(context.Context, *UpdateAppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAppServer) GenerateApiKey(context.Context, *GenerateAPIKeyAppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateApiKey not implemented")
}
func (UnimplementedAppServer) Enable(context.Context, *AppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedAppServer) Disable(context.Context, *AppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedAppServer) Delete(context.Context, *AppRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAppServer) Get(context.Context, *AppRequest) (*AppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAppServer) Check(context.Context, *AppCheckRequest) (*CheckAppInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAppServer) List(context.Context, *AppSearchQuery) (*AppsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAppServer) Stat(context.Context, *AppStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Update(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GenerateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAPIKeyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GenerateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_GenerateApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GenerateApiKey(ctx, req.(*GenerateAPIKeyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Enable(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Disable(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Delete(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Get(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Check(ctx, req.(*AppCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).List(ctx, req.(*AppSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Stat(ctx, req.(*AppStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _App_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _App_Update_Handler,
		},
		{
			MethodName: "generate_api_key",
			Handler:    _App_GenerateApiKey_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _App_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _App_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _App_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _App_Get_Handler,
		},
		{
			MethodName: "check",
			Handler:    _App_Check_Handler,
		},
		{
			MethodName: "list",
			Handler:    _App_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _App_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/app.proto",
}