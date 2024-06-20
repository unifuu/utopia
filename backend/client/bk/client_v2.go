package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("Joined game with ID: %s\n", joinResp.GetGameId())

	reader := bufio.NewReader(os.Stdin)

	for {
		// Read game_id from command prompt
		fmt.Print("Enter game_id: ")
		gameID, _ := reader.ReadString('\n')
		gameID = gameID[:len(gameID)-1]

		// Read x position from command prompt
		fmt.Print("Enter x position: ")
		xPosStr, _ := reader.ReadString('\n')
		xPosStr = strings.TrimSpace(xPosStr)
		xPos, err := strconv.Atoi(xPosStr)
		if err != nil {
			fmt.Println("Invalid x position")
			continue
		}

		// Read y position from command prompt
		fmt.Print("Enter y position: ")
		yPosStr, _ := reader.ReadString('\n')
		yPosStr = strings.TrimSpace(yPosStr)
		yPos, err := strconv.Atoi(yPosStr)
		if err != nil {
			fmt.Println("Invalid y position")
			continue
		}

		// Move the player
		moveResp, err := client.MovePlayer(ctx, &pb.MovePlayerReq{
			GameId:   gameID,
			PlayerId: "player1",
			X:        int32(xPos),
			Y:        int32(yPos),
		})
		if err != nil {
			log.Fatalf("could not move player: %v", err)
		}
		fmt.Printf("Move response: %s\n", moveResp.GetMessage())
	}
}
