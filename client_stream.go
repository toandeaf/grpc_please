package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc_please/generated/grpc_please/hello/v1"
)

func main() {
	conn, err := grpc.NewClient("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)

	// Call SayHello
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.SayHelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	// Call Chat
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("could not establish chat: %v", err)
	}

	// Send messages to the server
	go func() {
		for _, msg := range []string{"Hello", "How are you?", "Goodbye"} {
			if err := stream.Send(&pb.ChatRequest{User: "client", Message: msg}); err != nil {
				log.Fatalf("could not send message: %v", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// Receive messages from the server
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Printf("Received message from %s: %s", resp.User, resp.Message)
	}
}
