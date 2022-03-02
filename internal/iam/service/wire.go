//go:build wireinject
// +build wireinject

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	manager "github.com/yangchnet/skeleton/internal/iam/biz/ladon_manager"
	"github.com/yangchnet/skeleton/internal/iam/conf"
	"github.com/yangchnet/skeleton/internal/iam/data"
)

func InitService(ctx context.Context, conf *conf.Bootstrap) (*service, error) {
	panic(wire.Build(
		manager.ProviderSet,
		biz.ProviderSet,
		data.ProviderSet,
		NewIamService,
	))
}
