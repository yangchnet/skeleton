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
var ProviderSet = wire.NewSet(
	NewEnt,
	NewCache,
	NewUserRepo,
	NewRoleRepo,
	NewBindRepo,
	NewTenantRepo,
	NewPolicyRepo,
)

// NewEnt create an ent client.
func NewEnt(ctx context.Context, c *conf.Bootstrap) (*ent.Client, error) {
	drv, err := entsql.Open(c.Data.Database.Driver, c.Data.Database.Dsn)
	if err != nil {
		logger.Error("Error creating database connection: %v", err)

		return nil, err
	}

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(
		ctx,
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

	if c.Data.Seed {
		_ = seed(ctx, client)
	}

	return client, nil
}

// NewCache creates a new cache.
func NewCache(ctx context.Context, c *conf.Bootstrap) cache.CacheInterface {
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

// var _ biz.UserRepo = (*Data)(nil)

// var _ biz.BindRepo = (*Data)(nil)

// var _ biz.RoleRepo = (*Data)(nil)

// var _ biz.TenantRepo = (*Data)(nil)

// var _ biz.PolicyRepo = (*Data)(nil)

// NewUserRepo creates a new data which is a biz.UserRepo.
func NewUserRepo(ctx context.Context, ec *ent.Client, cache cache.CacheInterface) biz.UserRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}

// NewRoleRepo creates a new data which is a biz.RoleRepo.
func NewRoleRepo(ctx context.Context, ec *ent.Client, cache cache.CacheInterface) biz.RoleRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}

// NewBindRepo creates a new data which is a biz.BindRepo.
func NewBindRepo(ctx context.Context, ec *ent.Client, cache cache.CacheInterface) biz.BindRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}

// NewTenantRepo creates a new data which is a biz.TenantRepo.
func NewTenantRepo(ctx context.Context, ec *ent.Client, cache cache.CacheInterface) biz.TenantRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}

// NewPolicyRepo creates a new data which is a biz.PolicyRepo.
func NewPolicyRepo(ctx context.Context, ec *ent.Client, cache cache.CacheInterface) biz.PolicyRepo {
	return &Data{
		db:    ec,
		cache: cache,
	}
}
