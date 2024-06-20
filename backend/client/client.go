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

	reader := bufio.NewReader(os.Stdin)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Read player ID from command prompt
		fmt.Print("Enter player ID: ")
		playerID, _ := reader.ReadString('\n')
		playerID = strings.TrimSpace(playerID)

		// Read game_id from command prompt
		fmt.Print("Enter game ID: ")
		gameID, _ := reader.ReadString('\n')
		gameID = strings.TrimSpace(gameID)

		// Join the game
		joinResp, err := client.JoinGame(ctx, &pb.JoinGameReq{
			PlayerId: playerID,
			GameId:   gameID,
		})
		if err != nil {
			log.Printf("could not join game: %v", err)
			continue
		}
		fmt.Printf("Join response: %s\n", joinResp.GetMessage())

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
			PlayerId: playerID,
			X:        int32(xPos),
			Y:        int32(yPos),
		})
		if err != nil {
			log.Printf("could not move player: %v", err)
			continue
		}
		fmt.Printf("Move response: %s\n", moveResp.GetMessage())
	}
}
