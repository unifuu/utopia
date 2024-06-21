package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	pb "game_service/proto"

	"google.golang.org/grpc"
)

func main() {
	var playerID string
	var gameID string

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)

	// Set up a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	// Notify the channel on interrupt signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to signal when the program is exiting
	done := make(chan bool, 1)

	// Run a goroutine to handle cleanup on signal reception
	go func() {
		<-sigs
		fmt.Println("\nReceived an interrupt, cleaning up...")

		// Quit the game
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		quitResp, err := client.QuitGame(ctx, &pb.QuitGameReq{
			PlayerId: playerID,
			GameId:   gameID,
		})
		if err != nil {
			log.Printf("could not quit the game: %v", err)
		}

		log.Println("quitResp:", quitResp.GetMessage())

		done <- true
	}()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Read player ID from command prompt
		fmt.Print("Enter player ID: ")
		playerID, _ = reader.ReadString('\n')
		playerID = strings.TrimSpace(playerID)

		// Read game ID from command prompt
		fmt.Print("Enter game ID: ")
		gameID, _ = reader.ReadString('\n')
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

		if joinResp.Success {
			break
		} else {
			continue
		}
	}

	for {
		select {
		case <-done:
			fmt.Println("Exiting the program...")
			return

		default:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

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
}
