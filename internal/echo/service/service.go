package service

import (
	"context"

	"github.com/google/wire"
	v1 "github.com/yangchnet/skeleton/api/echo/v1"
	"github.com/yangchnet/skeleton/internal/echo/biz"
)

// ProviderSet provided NewEchoService.
var ProviderSet = wire.NewSet(NewEchoService)

// EchoService implements grpc/http server.
type EchoService struct {
	v1.UnimplementedEchoServiceServer

	uc *biz.EchoCase
}

// NewEchoService creates a new EchoService.
func NewEchoService(ctx context.Context, uc *biz.EchoCase) *EchoService {
	return &EchoService{
		uc: uc,
	}
}
