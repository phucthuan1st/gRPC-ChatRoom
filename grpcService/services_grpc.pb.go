// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: grpcService/services.proto

package grpcService

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

// ChatRoomClient is the client API for ChatRoom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatRoomClient interface {
	// broadcast message to every one in room chat
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error)
	// send private message to user (no broadcast)
	SendPrivateMessage(ctx context.Context, in *PrivateChatMessage, opts ...grpc.CallOption) (*SentMessageStatus, error)
	// Register for a new client account
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthenticationResult, error)
	// login using a pair of username and password
	Login(ctx context.Context, in *UserLoginCredentials, opts ...grpc.CallOption) (*AuthenticationResult, error)
	// Get a list of information of connected peers or specific peers
	GetConnectedPeers(ctx context.Context, in *Command, opts ...grpc.CallOption) (*UserList, error)
	// Get a peer information (except password)
	GetPeerInfomations(ctx context.Context, in *Command, opts ...grpc.CallOption) (*User, error)
}

type chatRoomClient struct {
	cc grpc.ClientConnInterface
}

func NewChatRoomClient(cc grpc.ClientConnInterface) ChatRoomClient {
	return &chatRoomClient{cc}
}

func (c *chatRoomClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatRoom_ServiceDesc.Streams[0], "/grpcService.ChatRoom/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRoomChatClient{stream}
	return x, nil
}

type ChatRoom_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatRoomChatClient struct {
	grpc.ClientStream
}

func (x *chatRoomChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRoomChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatRoomClient) SendPrivateMessage(ctx context.Context, in *PrivateChatMessage, opts ...grpc.CallOption) (*SentMessageStatus, error) {
	out := new(SentMessageStatus)
	err := c.cc.Invoke(ctx, "/grpcService.ChatRoom/SendPrivateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthenticationResult, error) {
	out := new(AuthenticationResult)
	err := c.cc.Invoke(ctx, "/grpcService.ChatRoom/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomClient) Login(ctx context.Context, in *UserLoginCredentials, opts ...grpc.CallOption) (*AuthenticationResult, error) {
	out := new(AuthenticationResult)
	err := c.cc.Invoke(ctx, "/grpcService.ChatRoom/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomClient) GetConnectedPeers(ctx context.Context, in *Command, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/grpcService.ChatRoom/GetConnectedPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomClient) GetPeerInfomations(ctx context.Context, in *Command, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/grpcService.ChatRoom/GetPeerInfomations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatRoomServer is the server API for ChatRoom service.
// All implementations must embed UnimplementedChatRoomServer
// for forward compatibility
type ChatRoomServer interface {
	// broadcast message to every one in room chat
	Chat(ChatRoom_ChatServer) error
	// send private message to user (no broadcast)
	SendPrivateMessage(context.Context, *PrivateChatMessage) (*SentMessageStatus, error)
	// Register for a new client account
	Register(context.Context, *User) (*AuthenticationResult, error)
	// login using a pair of username and password
	Login(context.Context, *UserLoginCredentials) (*AuthenticationResult, error)
	// Get a list of information of connected peers or specific peers
	GetConnectedPeers(context.Context, *Command) (*UserList, error)
	// Get a peer information (except password)
	GetPeerInfomations(context.Context, *Command) (*User, error)
	mustEmbedUnimplementedChatRoomServer()
}

// UnimplementedChatRoomServer must be embedded to have forward compatible implementations.
type UnimplementedChatRoomServer struct {
}

func (UnimplementedChatRoomServer) Chat(ChatRoom_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatRoomServer) SendPrivateMessage(context.Context, *PrivateChatMessage) (*SentMessageStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPrivateMessage not implemented")
}
func (UnimplementedChatRoomServer) Register(context.Context, *User) (*AuthenticationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChatRoomServer) Login(context.Context, *UserLoginCredentials) (*AuthenticationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatRoomServer) GetConnectedPeers(context.Context, *Command) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectedPeers not implemented")
}
func (UnimplementedChatRoomServer) GetPeerInfomations(context.Context, *Command) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerInfomations not implemented")
}
func (UnimplementedChatRoomServer) mustEmbedUnimplementedChatRoomServer() {}

// UnsafeChatRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatRoomServer will
// result in compilation errors.
type UnsafeChatRoomServer interface {
	mustEmbedUnimplementedChatRoomServer()
}

func RegisterChatRoomServer(s grpc.ServiceRegistrar, srv ChatRoomServer) {
	s.RegisterService(&ChatRoom_ServiceDesc, srv)
}

func _ChatRoom_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatRoomServer).Chat(&chatRoomChatServer{stream})
}

type ChatRoom_ChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatRoomChatServer struct {
	grpc.ServerStream
}

func (x *chatRoomChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRoomChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatRoom_SendPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).SendPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.ChatRoom/SendPrivateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).SendPrivateMessage(ctx, req.(*PrivateChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoom_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.ChatRoom/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoom_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.ChatRoom/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).Login(ctx, req.(*UserLoginCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoom_GetConnectedPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).GetConnectedPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.ChatRoom/GetConnectedPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).GetConnectedPeers(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoom_GetPeerInfomations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).GetPeerInfomations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.ChatRoom/GetPeerInfomations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).GetPeerInfomations(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatRoom_ServiceDesc is the grpc.ServiceDesc for ChatRoom service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatRoom_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcService.ChatRoom",
	HandlerType: (*ChatRoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPrivateMessage",
			Handler:    _ChatRoom_SendPrivateMessage_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ChatRoom_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ChatRoom_Login_Handler,
		},
		{
			MethodName: "GetConnectedPeers",
			Handler:    _ChatRoom_GetConnectedPeers_Handler,
		},
		{
			MethodName: "GetPeerInfomations",
			Handler:    _ChatRoom_GetPeerInfomations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _ChatRoom_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcService/services.proto",
}