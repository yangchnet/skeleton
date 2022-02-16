package service

import (
	"github.com/yangchnet/skeleton/internal/iam/data"
	pb "github.com/yangchnet/skeleton/pkg/pb/iam/v1"
)

type service struct {
	iamData data.IamInterface
	pb.UnimplementedIamServiceServer
}
