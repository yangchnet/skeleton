package data

import (
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/data/ent"
	"github.com/yangchnet/skeleton/pkg/logger"
)

// ProviderSet provided NewData and NewEnt.
var ProviderSet = wire.NewSet(NewData, NewEnt)

// Data is a dao and implements biz.EchoRepo.
type Data struct {
	db *ent.Client
	// TODO: redis
}

// NewData creates a new data object.
func NewData(c *conf.Data, ec *ent.Client) (*Data, error) {
	return &Data{db: ec}, nil
}

// NewEnt create an ent client.
func NewEnt(c *conf.Data) (*ent.Client, error) {
	drv, err := entsql.Open(c.Database.Driver, c.Database.Dsn)
	if err != nil {
		logger.Error("Error creating database connection: %v", err)

		return nil, err
	}

	return ent.NewClient(ent.Driver(drv)), nil
}
