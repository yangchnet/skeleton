package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	ServiceName string
	Port        int
}

type RegisterCallBack func(*grpc.Server)

func NewGrpcServer(serviceName string, port int) *GrpcServer {
	return &GrpcServer{
		ServiceName: serviceName,
		Port:        port,
	}
}

func (g *GrpcServer) Serve(ctx context.Context, callback RegisterCallBack) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		fmt.Println("Error: %w", err)
	}

	grpcServer := grpc.NewServer()
	callback(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error: %w", err)
	}
}
