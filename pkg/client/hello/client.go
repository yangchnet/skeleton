package hello

import (
	"context"

	"github.com/google/wire"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	pb "github.com/yangchnet/skeleton/pkg/pb/hello/v1"
)

var ProviderSet = wire.NewSet(NewHelloServiceClient)

func NewHelloServiceClient(ctx context.Context) (pb.HelloServiceClient, error) {
	client, err := localGrpc.NewClient("skeleton-hello-manage", 10001) // TODO: constants
	if err != nil {
		return nil, err
	}

	return pb.NewHelloServiceClient(client), nil
}
