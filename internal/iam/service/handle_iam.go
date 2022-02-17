package service

import (
	"context"

	"github.com/yangchnet/skeleton/pkg/logger"
	pb "github.com/yangchnet/skeleton/pkg/pb/iam/v1"
	"github.com/yangchnet/skeleton/pkg/utils/pbutil"
)

var _ pb.IamServiceServer = (*service)(nil)

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	username := req.GetUsername().GetValue()
	password := req.GetPassword().GetValue()
	email := req.GetEmail().GetValue()

	logger.Info("username: %s, password: %s, email: %s", username, password, email)
	logger.Warn("username: %s, password: %s, email: %s", username, password, email)
	logger.Debug("username: %s, password: %s, email: %s", username, password, email)
	logger.Error("username: %s, password: %s, email: %s", username, password, email)

	return &pb.CreateUserResponse{
		Username: pbutil.ToProtoString("lichang"),
		UserId:   pbutil.ToProtoString("111"),
	}, nil
}
