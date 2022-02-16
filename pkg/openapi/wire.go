// +build wireinject

package apigateway

import (
	"context"

	"github.com/google/wire"
)

func InitService(context.Context) (Server, error) {
	panic(wire.Build(
		wire.Struct(new(Server), "*"),
	))
}
