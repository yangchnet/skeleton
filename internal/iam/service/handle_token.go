package service

import (
	"context"

	v1 "github.com/yangchnet/skeleton/api/iam/v1"
)

// Token.
func (s *service) Login(_ context.Context, _ *v1.LoginRequest) (*v1.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}
