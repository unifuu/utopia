package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "game_service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGameServiceServer
	mu    sync.Mutex
	games map[string]map[string]bool // map[game_id]map[player_id]bool
}

func NewServer() *server {
	return &server{
		games: make(map[string]map[string]bool),
	}
}

func (s *server) JoinGame(ctx context.Context, req *pb.JoinGameReq) (*pb.JoinGameResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	players, gameExists := s.games[req.GetGameId()]
	if !gameExists {
		s.games[req.GetGameId()] = make(map[string]bool)
		players = s.games[req.GetGameId()]
	}

	if _, playerExists := players[req.GetPlayerId()]; !playerExists {
		players[req.GetPlayerId()] = true
		fmt.Println(req.GetPlayerId(), "has joined the game:", req.GetGameId())
		return &pb.JoinGameResp{GameId: req.GetGameId(), Message: "Player added to the game"}, nil
	}

	return &pb.JoinGameResp{GameId: req.GetGameId(), Message: "Player already in the game"}, nil
}

func (s *server) MovePlayer(ctx context.Context, req *pb.MovePlayerReq) (*pb.MovePlayerResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	players, gameExists := s.games[req.GetGameId()]
	if !gameExists {
		return &pb.MovePlayerResp{Success: false, Message: "Game not found"}, nil
	}

	if _, playerExists := players[req.GetPlayerId()]; !playerExists {
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
