package service

import (
	"context"
	"errors"

	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Token.
func (s *service) Token(ctx context.Context, req *v1.TokenRequest) (*v1.TokenResponse, error) {
	token, err := s.uc.CreateToken(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, biz.ErrNotFoundUser) {
			return nil, status.Errorf(codes.NotFound, "username not exist")
		}
		return nil, err
	}

	return &v1.TokenResponse{
		Token: token.Jwt,
	}, nil
}
