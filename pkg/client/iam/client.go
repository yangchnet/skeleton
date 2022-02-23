package iam

// import (
// 	"context"

// 	"github.com/google/wire"
// 	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
// 	pb "github.com/yangchnet/skeleton/pkg/pb/iam/v1"
// )

// var ProviderSet = wire.NewSet(NewIamServiceClient)

// func NewIamServiceClient(ctx context.Context) (pb.IamServiceClient, error) {
// 	client, err := localGrpc.NewClient("skeleton-iam-manage", 10001) // TODO: constants
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pb.NewIamServiceClient(client), nil
// }
