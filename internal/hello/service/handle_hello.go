package service

import (
	"context"

	pb "github.com/yangchnet/skeleton/pkg/pb/hello/v1"
)

var _ pb.HelloServiceServer = (*service)(nil)

func (s *service) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return nil, nil
}
