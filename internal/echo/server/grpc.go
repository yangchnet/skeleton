package server

import (
	"context"

	v1 "github.com/yangchnet/skeleton/api/echo/v1"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/service"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	"github.com/yangchnet/skeleton/pkg/logger"
	"google.golang.org/grpc"
)

// Serve start echo service.
func GrpcServe(ctx context.Context, c *conf.Bootstrap) {
	service, err := service.InitService(ctx, c)
	if err != nil {
		logger.Panicf("error initializing service: %v", err)
	}

	localGrpc.NewGrpcServer(
		c.Server.Grpc.ServiceName,
		int(c.Server.Grpc.Port),
		grpc.ConnectionTimeout(c.Server.Grpc.Timeout.AsDuration()),
	).Serve(ctx, func(server *grpc.Server) {
		v1.RegisterEchoServiceServer(server, service)
	})
}

// TODO: middleware recover and etc.
