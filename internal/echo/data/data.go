package data

import (
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/data/ent"
	"github.com/yangchnet/skeleton/pkg/logger"
)

// ProviderSet provided NewData and NewEnt.
var ProviderSet = wire.NewSet(NewEnt, NewData)

// NewEnt create an ent client.
func NewEnt(c *conf.Bootstrap) (*ent.Client, error) {
	drv, err := entsql.Open(c.Data.Database.Driver, c.Data.Database.Dsn)
	if err != nil {
		logger.Error("Error creating database connection: %v", err)

		return nil, err
	}

	return ent.NewClient(ent.Driver(drv)), nil
}

// Data handle db connection and implemented biz.EchoRepo.
type Data struct {
	db *ent.Client
	// TODO: redis
}

// NewData creates a new data which is a biz.EchoRepo.
func NewData(ec *ent.Client) biz.EchoRepo {
	return &Data{db: ec}
}
