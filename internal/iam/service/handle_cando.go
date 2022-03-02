package service

import (
	"context"

	v1 "github.com/yangchnet/skeleton/api/iam/v1"
)

func (s *service) CanDo(ctx context.Context, req *v1.CanDoRequest) (*v1.CanDoResponse, error) {
	if s.uc.CanDo(ctx, "root", req.GetResource(), req.GetAction()) {
		return &v1.CanDoResponse{
			Allowed: true,
		}, nil
	}

	return &v1.CanDoResponse{
		Allowed: false,
	}, nil
}
