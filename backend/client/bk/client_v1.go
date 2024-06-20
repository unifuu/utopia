package main

import (
	"context"
	"log"
	"time"

	pb "game_service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Join a game
	joinResp, err := client.JoinGame(ctx, &pb.JoinGameReq{PlayerId: "player1"})
	if err != nil {
		log.Fatalf("could not join game: %v", err)
	}
	log.Printf("Joined game with ID: %s", joinResp.GetGameId())

	// Move the player
	moveResp, err := client.MovePlayer(ctx, &pb.MovePlayerReq{
		GameId:   joinResp.GetGameId(),
		PlayerId: "player1",
		X:        10,
		Y:        20,
	})
	if err != nil {
		log.Fatalf("could not move player: %v", err)
	}
	log.Printf("Move response: %s", moveResp.GetMessage())
}
