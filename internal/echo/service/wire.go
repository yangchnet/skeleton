//go:build wireinject
// +build wireinject

package service

import (
	"context"

	"github.com/google/wire"

	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/data"
)

func InitService(ctx context.Context, conf *conf.Bootstrap) (*EchoService, error) {
	panic(wire.Build(
		biz.ProviderSet,
		data.ProviderSet,
		NewEchoService,
	))
}
