// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: game_service.proto

package game_service

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

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	JoinGame(ctx context.Context, in *JoinGameReq, opts ...grpc.CallOption) (*JoinGameResp, error)
	MovePlayer(ctx context.Context, in *MovePlayerReq, opts ...grpc.CallOption) (*MovePlayerResp, error)
	QuitGame(ctx context.Context, in *QuitGameReq, opts ...grpc.CallOption) (*QuitGameResp, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) JoinGame(ctx context.Context, in *JoinGameReq, opts ...grpc.CallOption) (*JoinGameResp, error) {
	out := new(JoinGameResp)
	err := c.cc.Invoke(ctx, "/game_service.GameService/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) MovePlayer(ctx context.Context, in *MovePlayerReq, opts ...grpc.CallOption) (*MovePlayerResp, error) {
	out := new(MovePlayerResp)
	err := c.cc.Invoke(ctx, "/game_service.GameService/MovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) QuitGame(ctx context.Context, in *QuitGameReq, opts ...grpc.CallOption) (*QuitGameResp, error) {
	out := new(QuitGameResp)
	err := c.cc.Invoke(ctx, "/game_service.GameService/QuitGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	JoinGame(context.Context, *JoinGameReq) (*JoinGameResp, error)
	MovePlayer(context.Context, *MovePlayerReq) (*MovePlayerResp, error)
	QuitGame(context.Context, *QuitGameReq) (*QuitGameResp, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) JoinGame(context.Context, *JoinGameReq) (*JoinGameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedGameServiceServer) MovePlayer(context.Context, *MovePlayerReq) (*MovePlayerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovePlayer not implemented")
}
func (UnimplementedGameServiceServer) QuitGame(context.Context, *QuitGameReq) (*QuitGameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitGame not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).JoinGame(ctx, req.(*JoinGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_MovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovePlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).MovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/MovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).MovePlayer(ctx, req.(*MovePlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_QuitGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).QuitGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/QuitGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).QuitGame(ctx, req.(*QuitGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_service.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGame",
			Handler:    _GameService_JoinGame_Handler,
		},
		{
			MethodName: "MovePlayer",
			Handler:    _GameService_MovePlayer_Handler,
		},
		{
			MethodName: "QuitGame",
			Handler:    _GameService_QuitGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_service.proto",
}
