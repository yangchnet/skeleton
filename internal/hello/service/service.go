package service

import (
	"github.com/yangchnet/skeleton/internal/hello/data"
	pb "github.com/yangchnet/skeleton/pkg/pb/hello/v1"
)

type service struct {
	hello data.HelloInterface
	pb.UnimplementedHelloServiceServer
}
