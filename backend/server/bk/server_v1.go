package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "game_service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGameServiceServer
	games map[string][]string
}

func NewServer() *server {
	return &server{
		games: make(map[string][]string),
	}
}

func (s *server) JoinGame(ctx context.Context, req *pb.JoinGameReq) (*pb.JoinGameResp, error) {
	gameID := fmt.Sprintf("game_%d", len(s.games)+1)
	s.games[gameID] = append(s.games[gameID], req.GetPlayerId())
	return &pb.JoinGameResp{GameId: gameID}, nil
}

func (s *server) MovePlayer(ctx context.Context, req *pb.MovePlayerReq) (*pb.MovePlayerResp, error) {
	game, ok := s.games[req.GetGameId()]
	player := req.GetPlayerId()
	fmt.Println("MovePlayer: game:", game, "player:", player)

	if !ok {
		return &pb.MovePlayerResp{Success: false, Message: "Game not found"}, nil
	}

	playerFound := false
	for _, p := range game {
		if p == req.GetPlayerId() {
			playerFound = true
			break
		}
	}
	if !playerFound {
		return &pb.MovePlayerResp{Success: false, Message: "Player not in game"}, nil
	}

	// Simulate updating player position (simplified example)
	playerPosition := fmt.Sprintf("(%d, %d)", req.GetX(), req.GetY())
	fmt.Printf("Player %s moved to position %s\n", req.GetPlayerId(), playerPosition)

	return &pb.MovePlayerResp{Success: true, Message: "Move successful"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, NewServer())

	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
