// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protobufs/todo.proto

package todo

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

// TodoListClient is the client API for TodoList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoListClient interface {
	AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoResponse, error)
	GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
}

type todoListClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListClient(cc grpc.ClientConnInterface) TodoListClient {
	return &todoListClient{cc}
}

func (c *todoListClient) AddTodo(ctx context.Context, in *AddTodoRequest, opts ...grpc.CallOption) (*AddTodoResponse, error) {
	out := new(AddTodoResponse)
	err := c.cc.Invoke(ctx, "/todoapp.TodoList/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error) {
	out := new(GetTodosResponse)
	err := c.cc.Invoke(ctx, "/todoapp.TodoList/GetTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/todoapp.TodoList/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todoapp.TodoList/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServer is the server API for TodoList service.
// All implementations must embed UnimplementedTodoListServer
// for forward compatibility
type TodoListServer interface {
	AddTodo(context.Context, *AddTodoRequest) (*AddTodoResponse, error)
	GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	mustEmbedUnimplementedTodoListServer()
}

// UnimplementedTodoListServer must be embedded to have forward compatible implementations.
type UnimplementedTodoListServer struct {
}

func (UnimplementedTodoListServer) AddTodo(context.Context, *AddTodoRequest) (*AddTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodoListServer) GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}
func (UnimplementedTodoListServer) UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoListServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoListServer) mustEmbedUnimplementedTodoListServer() {}

// UnsafeTodoListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoListServer will
// result in compilation errors.
type UnsafeTodoListServer interface {
	mustEmbedUnimplementedTodoListServer()
}

func RegisterTodoListServer(s grpc.ServiceRegistrar, srv TodoListServer) {
	s.RegisterService(&TodoList_ServiceDesc, srv)
}

func _TodoList_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapp.TodoList/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).AddTodo(ctx, req.(*AddTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapp.TodoList/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).GetTodos(ctx, req.(*GetTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapp.TodoList/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapp.TodoList/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoList_ServiceDesc is the grpc.ServiceDesc for TodoList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoapp.TodoList",
	HandlerType: (*TodoListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _TodoList_AddTodo_Handler,
		},
		{
			MethodName: "GetTodos",
			Handler:    _TodoList_GetTodos_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoList_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoList_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/todo.proto",
}
