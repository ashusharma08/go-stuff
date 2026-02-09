package main

import (
	"context"
	"log/slog"
	"net"

	"github.com/esoptra/go-prac/grpc-server/handler"
	pb "github.com/esoptra/go-prac/grpc-server/proto"
	"google.golang.org/grpc"
)

func main() {
	l, _ := net.Listen("tcp", ":8080")

	s := handler.NewServer()
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, s)
	if err := server.Serve(l); err != nil {
		slog.Log(context.Background(), slog.LevelError, "error fatal")
	}
}
