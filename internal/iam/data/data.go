package data

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"github.com/yangchnet/skeleton/internal/iam/conf"
	"github.com/yangchnet/skeleton/internal/iam/data/ent"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/migrate"
	"github.com/yangchnet/skeleton/pkg/cache"
	"github.com/yangchnet/skeleton/pkg/logger"
)

// ProviderSet provided NewData and NewEnt.
var ProviderSet = wire.NewSet(NewEnt, NewCache, NewData)

// NewEnt create an ent client.
func NewEnt(c *conf.Bootstrap) (*ent.Client, error) {
	drv, err := entsql.Open(c.Data.Database.Driver, c.Data.Database.Dsn)
	if err != nil {
		logger.Error("Error creating database connection: %v", err)

		return nil, err
	}

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				// Run custom code here.
				return next.Create(ctx, tables...)
			})
		}),
	); err != nil {
		logger.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}

// NewCache creates a new cache.
func NewCache(c *conf.Bootstrap) cache.CacheInterface {
	return cache.NewRedisStore(
		fmt.Sprintf("%s:%d", c.Data.Redis.Host, c.Data.Redis.Port),
		c.Data.Redis.Password,
		int(c.Data.Redis.Db))
}

// Data handle db connection and implemented biz.EchoRepo.
type Data struct {
	db    *ent.Client
	cache cache.CacheInterface
}

var _ biz.IamRepo = (*Data)(nil)

// NewData creates a new data which is a biz.EchoRepo.
func NewData(ec *ent.Client, cache cache.CacheInterface) biz.IamRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}
