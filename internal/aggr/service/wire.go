// +build wireinject

package service

import (
	"context"

	"github.com/google/wire"

	helloClient "github.com/yangchnet/skeleton/pkg/client/hello"
	iamClient "github.com/yangchnet/skeleton/pkg/client/iam"
)

func InitService(ctx context.Context) (service, error) {
	panic(wire.Build(
		iamClient.ProviderSet,
		helloClient.ProviderSet,
		wire.Struct(new(service), "*"),
	))
}
