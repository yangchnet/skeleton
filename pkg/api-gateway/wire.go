//go:build wireinject
// +build wireinject

package apigateway

import (
	"context"

	"github.com/google/wire"
	"github.com/gin-gonic/gin"
)

func InitService(context.Context) (Server, error) {
	panic(wire.Build(
		gin.New(),
		wire.Struct(new(Server), "*"),
	))
}
