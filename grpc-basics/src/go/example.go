package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

import pb "./protocol"

type server struct{}

func (s *server) Ok(ctx context.Context, req *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	return &pb.HealthcheckResponse{Status: 2}, nil // positive response, no errors
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterHealthcheckServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
