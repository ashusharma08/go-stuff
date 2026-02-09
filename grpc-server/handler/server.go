package handler

import (
	"context"

	pb "github.com/esoptra/go-prac/grpc-server/proto"
)

type server struct {
	pb.UnimplementedAgentServer
}

func NewServer() *server {
	return &server{}
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello Mr. " + in.GetName(),
	}, nil
}
