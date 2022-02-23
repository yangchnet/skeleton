package data

import (
	"github.com/google/wire"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/data/ent"
)

var ProviderSet = wire.NewSet(NewData, NewEnt)

type Data struct {
	db *ent.Client
	// TODO: redis
}

func NewData(c *conf.Data, ec *ent.Client) (*Data, error) {
	return nil, nil // TODO: implementation
}

func NewEnt(c *conf.Data) *ent.Client {
	return nil // TODO: implementation
}
