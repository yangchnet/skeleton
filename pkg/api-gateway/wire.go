//go:build wireinject
// +build wireinject

package apigateway

import (
	"context"

	"github.com/google/wire"
	iamClient "github.com/yangchnet/skeleton/pkg/client/iam"
	"github.com/yangchnet/skeleton/pkg/token"
)

func InitService(ctx context.Context, secretKey string) (*Server, error) {
	panic(wire.Build(
		iamClient.ProviderSet,
		token.ProviderSet,
		ProviderSet,
	))
}
