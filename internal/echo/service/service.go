package service

import (
	"github.com/google/wire"
	v1 "github.com/yangchnet/skeleton/api/echo/v1"
	"github.com/yangchnet/skeleton/internal/echo/biz"
)

var ProviderSet = wire.NewSet(NewEchoService)

type EchoService struct {
	v1.UnimplementedEchoServiceServer

	uc *biz.EchoCase
}

func NewEchoService(uc *biz.EchoCase) *EchoService {
	panic("not implemented") // TODO not implemented
}
