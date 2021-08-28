// +build wireinject

package service

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/hello/data"
)

func InitService(ctx context.Context, db *sql.DB) service {
	panic(wire.Build(
		data.ProviderSet,
		wire.Struct(new(service), "*"),
	))
}
