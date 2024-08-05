package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc_please/generated/grpc_please/hello/v1"
)

type serverStream struct {
	pb.UnimplementedExampleServiceServer
}

func (s *serverStream) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + req.Name}, nil
}

func (s *serverStream) Chat(stream pb.ExampleService_ChatServer) error {
	for {
		// Receive a message from the client
		req, err := stream.Recv()
		if err == io.EOF {
			// Client has finished sending messages
			return nil
		}
		if err != nil {
			return err
		}

		// Log the received message
		log.Printf("Received message from %s: %s", req.User, req.Message)

		// Send a response back to the client
		resp := &pb.ChatResponse{
			User:    req.User,
			Message: "Received: " + req.Message,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &serverStream{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
