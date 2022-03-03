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
	"github.com/yangchnet/skeleton/pkg/token"
)

func InitService(ctx context.Context, conf *conf.Bootstrap, secretKey string) (*service, error) {
	panic(wire.Build(
		manager.ProviderSet,
		biz.ProviderSet,
		data.ProviderSet,
		token.ProviderSet,
		NewIamService,
	))
}
