package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/yangchnet/skeleton/pkg/logger"
	"google.golang.org/grpc"
)

// GrpcServer is a grpc server.
type GrpcServer struct {
	ServiceName string
	Port        int
	Opts        []grpc.ServerOption
}

// RegisterCallBack registers service for grpc server.
type RegisterCallBack func(*grpc.Server)

// NewGrpcServer creates a new grpc server with serviceName, port and opts.
func NewGrpcServer(serviceName string, port int, opts ...grpc.ServerOption) *GrpcServer {
	return &GrpcServer{
		ServiceName: serviceName,
		Port:        port,
		Opts:        opts,
	}
}

// Serve start a grpc server and panic when failed.
func (g *GrpcServer) Serve(ctx context.Context, callback RegisterCallBack) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		logger.Panicf("Error listening on endpoint %v", err)
	}

	grpcServer := grpc.NewServer(g.Opts...)
	callback(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Panicf("Error serving: %v", err)
	}
}

// TODO: recover middleware
