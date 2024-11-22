// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/opsflow/v1/task.proto

package v1

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Task_Create_FullMethodName       = "/spaceone.api.opsflow.v1.Task/create"
	Task_Update_FullMethodName       = "/spaceone.api.opsflow.v1.Task/update"
	Task_ChangeStatus_FullMethodName = "/spaceone.api.opsflow.v1.Task/change_status"
	Task_Delete_FullMethodName       = "/spaceone.api.opsflow.v1.Task/delete"
	Task_Get_FullMethodName          = "/spaceone.api.opsflow.v1.Task/get"
	Task_List_FullMethodName         = "/spaceone.api.opsflow.v1.Task/list"
	Task_Stat_FullMethodName         = "/spaceone.api.opsflow.v1.Task/stat"
)

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	Create(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskInfo, error)
	Update(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskInfo, error)
	ChangeStatus(ctx context.Context, in *TaskChangeStatusRequest, opts ...grpc.CallOption) (*TaskInfo, error)
	Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskInfo, error)
	List(ctx context.Context, in *TaskSearchQuery, opts ...grpc.CallOption) (*TasksInfo, error)
	Stat(ctx context.Context, in *TaskStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Create(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskInfo)
	err := c.cc.Invoke(ctx, Task_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Update(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskInfo)
	err := c.cc.Invoke(ctx, Task_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) ChangeStatus(ctx context.Context, in *TaskChangeStatusRequest, opts ...grpc.CallOption) (*TaskInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskInfo)
	err := c.cc.Invoke(ctx, Task_ChangeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Delete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Task_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Get(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskInfo)
	err := c.cc.Invoke(ctx, Task_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) List(ctx context.Context, in *TaskSearchQuery, opts ...grpc.CallOption) (*TasksInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TasksInfo)
	err := c.cc.Invoke(ctx, Task_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Stat(ctx context.Context, in *TaskStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Task_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility.
type TaskServer interface {
	Create(context.Context, *TaskCreateRequest) (*TaskInfo, error)
	Update(context.Context, *TaskUpdateRequest) (*TaskInfo, error)
	ChangeStatus(context.Context, *TaskChangeStatusRequest) (*TaskInfo, error)
	Delete(context.Context, *TaskRequest) (*empty.Empty, error)
	Get(context.Context, *TaskRequest) (*TaskInfo, error)
	List(context.Context, *TaskSearchQuery) (*TasksInfo, error)
	Stat(context.Context, *TaskStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskServer struct{}

func (UnimplementedTaskServer) Create(context.Context, *TaskCreateRequest) (*TaskInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskServer) Update(context.Context, *TaskUpdateRequest) (*TaskInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTaskServer) ChangeStatus(context.Context, *TaskChangeStatusRequest) (*TaskInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedTaskServer) Delete(context.Context, *TaskRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaskServer) Get(context.Context, *TaskRequest) (*TaskInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTaskServer) List(context.Context, *TaskSearchQuery) (*TasksInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTaskServer) Stat(context.Context, *TaskStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}
func (UnimplementedTaskServer) testEmbeddedByValue()              {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	// If the following call pancis, it indicates UnimplementedTaskServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Create(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Update(ctx, req.(*TaskUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_ChangeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).ChangeStatus(ctx, req.(*TaskChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Delete(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Get(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).List(ctx, req.(*TaskSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Stat(ctx, req.(*TaskStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.opsflow.v1.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Task_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Task_Update_Handler,
		},
		{
			MethodName: "change_status",
			Handler:    _Task_ChangeStatus_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Task_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Task_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Task_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Task_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/opsflow/v1/task.proto",
}
