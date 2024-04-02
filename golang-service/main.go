package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/vinitparekh17/grpc-api/golang-service/protobufs"
	"google.golang.org/grpc"
)

// Server configuration
type Server struct {
	Port int
	pb.UnimplementedTodoListServer
}

func (s *Server) Start() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server started on %v", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterTodoListServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	server := Server{
		Port: 50051,
	}

	server.Start()
}
