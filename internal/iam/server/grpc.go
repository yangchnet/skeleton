package server

import (
	"context"

	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/iam/conf"
	"github.com/yangchnet/skeleton/internal/iam/service"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	"github.com/yangchnet/skeleton/pkg/logger"
	"google.golang.org/grpc"
)

// GrpcServestart echo service.
func GrpcServe(ctx context.Context, c *conf.Bootstrap) {
	service, err := service.InitService(ctx, c, c.Token.Secret)
	if err != nil {
		logger.Panicf("error initializing service: %v", err)
	}

	localGrpc.NewGrpcServer(
		c.Server.Grpc.ServiceName,
		int(c.Server.Grpc.Port),
		grpc.ConnectionTimeout(c.Server.Grpc.Timeout.AsDuration()),
	).Serve(ctx, func(server *grpc.Server) {
		v1.RegisterIamServiceServer(server, service)
	})
}
