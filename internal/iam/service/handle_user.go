package service

import (
	"context"
	"errors"
	"fmt"

	v1common "github.com/yangchnet/skeleton/api/common"
	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"github.com/yangchnet/skeleton/tools/pbtools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone().GetValue(),
		Email:    req.GetEmail().GetValue(),
	})
	if err != nil {
		if errors.Is(err, biz.ErrExistUsername) {
			return &v1.CreateUserResponse{
				Result:     v1common.OperationResult_FAIL,
				FailReason: pbtools.ToProtoString(err.Error()),
			}, status.Error(codes.AlreadyExists, err.Error())
		}

		return &v1.CreateUserResponse{
			Result:     v1common.OperationResult_FAIL,
			FailReason: pbtools.ToProtoString(err.Error()),
		}, err
	}

	return &v1.CreateUserResponse{
		UserId: pbtools.ToProtoString(fmt.Sprintf("%d", user.ID)),
		Result: v1common.OperationResult_SUCCESS,
	}, nil
}
