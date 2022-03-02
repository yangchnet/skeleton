package iam

import (
	"context"

	"github.com/google/wire"
	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
)

var ProviderSet = wire.NewSet(NewIamServiceClient)

func NewIamServiceClient(ctx context.Context) (v1.IamServiceClient, error) {
	client, err := localGrpc.NewClient("skeleton-iam-manage", 10003) // TODO: constants
	if err != nil {
		return nil, err
	}
	return v1.NewIamServiceClient(client), nil
}
